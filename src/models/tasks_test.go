package models

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/anhnguyen300795/Task-manager/src/interfaces"
)

func TestGetAllTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	taskModel := &Tasks{db}

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

	tasks := taskModel.GetAll()

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Error(`Comments data are not expected`)
	}
}

func TestGetOneTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	taskModel := &Tasks{db}

	rows := sqlmock.NewRows([]string{"title", "status", "assignee", "assignor", "start_time", "end_time", "description"}).
		AddRow("title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription")

	mock.ExpectQuery("^SELECT (.+) FROM tasks").
		WithArgs(1).
		WillReturnRows(rows)

	commentRows := sqlmock.NewRows([]string{"ID", "content", "author", "date"}).
		AddRow(1, "comment 1", "hello", "20-10-2013").
		AddRow(2, "comment 2", "world", "20-10-2013")

	mock.ExpectQuery("^SELECT (.+) FROM comments").
		WillReturnRows(commentRows)

	expectedComments := []interfaces.Comment{
		interfaces.Comment{1, "comment 1", "hello", "20-10-2013"},
		interfaces.Comment{2, "comment 2", "world", "20-10-2013"},
	}

	expectedTasks := &interfaces.TaskQuery{1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription", expectedComments}

	tasks := taskModel.GetOne(1)
	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	if !reflect.DeepEqual(tasks, expectedTasks) {
		t.Error(`Comments data are not expected`)
	}
}

func TestAddTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	taskModel := &Tasks{db}

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^INSERT INTO tasks(.+)").
		WillReturnResult(result)

	rows := sqlmock.NewRows([]string{"title", "status", "assignee", "assignor", "start_time", "end_time", "description"}).
		AddRow("title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription")

	mock.ExpectQuery("^SELECT (.+) FROM tasks (.+)").
		WithArgs(1).
		WillReturnRows(rows)

	commentRows := sqlmock.NewRows([]string{"ID", "content", "author", "date"})

	mock.ExpectQuery("^SELECT (.+) FROM comments").
		WithArgs(1).
		WillReturnRows(commentRows)

	task, _ := taskModel.Add(&interfaces.Task{1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription"})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	expectedTask := &interfaces.TaskQuery{1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription", []interfaces.Comment{}}

	if !reflect.DeepEqual(task, expectedTask) {
		t.Error(`Comments data are not expected`)
	}
}

func TestUpdateTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	taskModel := &Tasks{db}

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^UPDATE tasks SET").
		WillReturnResult(result)

	updateErr := taskModel.Update(&interfaces.Task{1, "title1", "todo", "tester1", "assignor1", "20-10-2013", "20-10-2013", "testDescription"})

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	if updateErr != nil {
		t.Error(`Comments data are not updated`)
	}
}

func TestDeleteTask(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	taskModel := &Tasks{db}

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^DELETE FROM tasks").
		WithArgs(1).
		WillReturnResult(result)

	DeleteErr := taskModel.Delete(1)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	if DeleteErr != nil {
		t.Error(`Comments data are not updated`)
	}
}
