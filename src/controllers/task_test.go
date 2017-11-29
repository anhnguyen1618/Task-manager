package controllers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/anhnguyen300795/Task-manager/src/interfaces"
	"github.com/gorilla/mux"
)

func AssertJSON(actual []byte, data interface{}, t *testing.T) {
	expected, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when marshaling expected json data", err)
	}

	if bytes.Compare(expected, actual) != 0 {
		t.Errorf("the expected json: %s is different from actual %s", expected, actual)
	}
}

func AssertString(actual []byte, expected string, t *testing.T) {
	actualString := string(actual[:])

	if actualString != expected {
		t.Errorf("the expected string: %s is different from actual %s", expected, actual)
	}
}

func TestCommentControllerGET(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	env := &interfaces.Env{db, nil}
	Controllers := &Controllers{env}

	r := httptest.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	rows := sqlmock.NewRows([]string{"ID", "title", "status", "assignee", "assignor", "start_time", "end_time", "description"}).
		AddRow(1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription")

	mock.ExpectQuery("^SELECT (.+) FROM tasks").
		WillReturnRows(rows)

	commentRows := sqlmock.NewRows([]string{"ID", "content", "author", "date"}).
		AddRow(1, "comment 1", "hello", "20-10-2013").
		AddRow(2, "comment 2", "world", "20-10-2013")

	mock.ExpectQuery("^SELECT (.+) FROM comments").
		WithArgs(1).
		WillReturnRows(commentRows)

	expectedComments := []interfaces.Comment{
		interfaces.Comment{1, "comment 1", "hello", "20-10-2013"},
		interfaces.Comment{2, "comment 2", "world", "20-10-2013"},
	}

	expectedTasks := []interfaces.TaskQuery{
		interfaces.TaskQuery{1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription", expectedComments},
	}

	Controllers.AllTaskController(w, r)

	AssertJSON(w.Body.Bytes(), expectedTasks, t)
}

func TestCommentControllerPOST(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	env := &interfaces.Env{db, nil}
	Controllers := &Controllers{env}

	r := httptest.NewRequest("POST", "/api/tasks", nil)
	w := httptest.NewRecorder()

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^INSERT INTO tasks(.+)").
		WillReturnResult(result)

	rows := sqlmock.NewRows([]string{"title", "status", "assignee", "assignor", "start_time", "end_time", "description"}).
		AddRow("title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription")

	mock.ExpectQuery("^SELECT (.+) FROM tasks (.+)").
		WithArgs(1).
		WillReturnRows(rows)

	commentRows := sqlmock.NewRows([]string{"ID", "content", "author", "date"}).
		AddRow(1, "comment 1", "hello", "20-10-2013")

	mock.ExpectQuery("^SELECT (.+) FROM comments").
		WithArgs(1).
		WillReturnRows(commentRows)

	expectedComments := []interfaces.Comment{
		interfaces.Comment{1, "comment 1", "hello", "20-10-2013"},
	}

	expectedTask := interfaces.TaskQuery{1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription", expectedComments}

	Controllers.AllTaskController(w, r)

	AssertJSON(w.Body.Bytes(), expectedTask, t)
}

func TestUpdateTaskControllerPUT(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	env := &interfaces.Env{db, nil}
	Controllers := &Controllers{env}
	// NOTE THAT MUX ROUTER SHOULD BE CREATED TO TEST ROUTE URL PARAMS
	router := mux.NewRouter()
	router.HandleFunc("/api/tasks/{id}", Controllers.UpdateTaskController)

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^UPDATE tasks SET").
		WillReturnResult(result)

	reqBody := interfaces.Task{1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription"}
	jsonValue, _ := json.Marshal(reqBody)

	r := httptest.NewRequest("PUT", "/api/tasks/1", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, r)

	AssertJSON(w.Body.Bytes(), reqBody, t)
}
