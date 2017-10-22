package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/anhnguyen300795/Task-manager/interfaces"
)

func TestAddUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userModel := &Users{db}

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^INSERT INTO users(.+)").
		WillReturnResult(result)

	newUser := &interfaces.UserInfo{"Test", "testPassword", "test@gmail.com", "USER"}

	status, _ := userModel.AddOne(newUser)

	if status != "Create successfully" {
		t.Error(`Create failed`)
	}

}
