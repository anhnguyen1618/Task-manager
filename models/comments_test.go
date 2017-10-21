package models

import (
	"reflect"
	"testing"

	"../interfaces"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllCommentsByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	commentModel := &Comments{db}

	rows := sqlmock.NewRows([]string{"ID", "content", "author", "date"}).
		AddRow(1, "comment 1", "hello", "20-10-2013").
		AddRow(2, "comment 2", "world", "20-10-2013")

	mock.ExpectQuery("^SELECT (.+) FROM comments").
		WillReturnRows(rows)

	data := []interfaces.Comment{
		interfaces.Comment{1, "comment 1", "hello", "20-10-2013"},
		interfaces.Comment{2, "comment 2", "world", "20-10-2013"},
	}

	comments := commentModel.GetByTaskID(1)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	if !reflect.DeepEqual(comments, data) {
		t.Error(`Comments data are not expected`)
	}
}

func TestGetCommentByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	commentModel := &Comments{db}

	rows := sqlmock.NewRows([]string{"content", "author", "date"}).
		AddRow("comment 1", "hello", "20-10-2013")

	mock.ExpectQuery("^SELECT (.+) FROM comments").WillReturnRows(rows)

	expectedData := &interfaces.Comment{1, "comment 1", "hello", "20-10-2013"}

	comment := commentModel.GetByID(1)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expections: %s", err)
	}

	if !reflect.DeepEqual(comment, expectedData) {
		t.Error(`Comment data are not expected`)
	}
}

func TestAddComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^INSERT INTO comments(.+)").
		WillReturnResult(result)

	commentModel := &Comments{db}

	addedComment := &interfaces.Comment{1, "comment 1", "hello", "20-10-2013"}

	addedID, err := commentModel.Add(addedComment, 1)

	if err != nil {
		t.Fatalf("Error occur when inserting comment")
	}

	if addedID != 1 {
		t.Fatalf("Wrong inserted ID")
	}

}

func TestUpdateComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^UPDATE comments").
		WillReturnResult(result)

	commentModel := &Comments{db}

	addedComment := &interfaces.Comment{1, "comment 1", "hello", "20-10-2013"}

	err = commentModel.Update(addedComment)

	if err != nil {
		t.Fatalf("Error occur when updating comment")
	}

}

func TestDeleteComment(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^DELETE FROM comments WHERE").
		WillReturnResult(result)

	commentModel := &Comments{db}

	err = commentModel.Delete(1)

	if err != nil {
		t.Fatalf("Error occur when delete comment")
	}

}
