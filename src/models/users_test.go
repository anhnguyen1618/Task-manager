package models

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/anhnguyen300795/Task-manager/src/interfaces"
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

	createdUser := userModel.AddOne(newUser)

	expectedUser := &interfaces.UserInfo{"Test", "", "test@gmail.com", "USER"}

	if !reflect.DeepEqual(createdUser, expectedUser) {
		t.Error(`Create user failed`)
	}

}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userModel := &Users{db}

	rows := sqlmock.NewRows([]string{"username", "email", "password", "role"}).
		AddRow("testName", "testPassword", "test@gmail.com", "ADMIN").
		AddRow("testName1", "testPassword1", "test@gmail.com", "ADMIN")

	mock.ExpectQuery("^SELECT (.+) FROM users").
		WillReturnRows(rows)

	actualUsers := userModel.GetAll()

	expectedUser := []interfaces.UserInfo{
		interfaces.UserInfo{"testName", "", "testPassword", "ADMIN"},
		interfaces.UserInfo{"testName1", "", "testPassword1", "ADMIN"},
	}

	if !reflect.DeepEqual(actualUsers, expectedUser) {
		t.Error(`Get all users failed`)
	}
}

func TestGetOneUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userModel := &Users{db}

	rows := sqlmock.NewRows([]string{"email", "role"}).
		AddRow("test@gmail.com", "ADMIN")

	mock.ExpectQuery("^SELECT (.+) FROM users WHERE").
		WithArgs("testName1").
		WillReturnRows(rows)

	actualUser := userModel.GetOne("testName1")

	expectedUser := &interfaces.UserInfo{"testName1", "", "test@gmail.com", "ADMIN"}

	if !reflect.DeepEqual(actualUser, expectedUser) {
		t.Error(`Get one user failed`)
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userModel := &Users{db}

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^UPDATE users(.+)").
		WillReturnResult(result)

	updatedUser := &interfaces.UserInfo{"Test", "", "test@gmail.com", "new role"}

	returnedUser := userModel.UpdateOne(updatedUser)

	expectedUser := &interfaces.UserInfo{"Test", "", "test@gmail.com", "new role"}

	if !reflect.DeepEqual(returnedUser, expectedUser) {
		t.Error(`Update user failed`)
	}

}

func TestDeleteUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	userModel := &Users{db}

	result := sqlmock.NewResult(1, 1)
	mock.ExpectExec("^DELETE users(.+)").
		WillReturnResult(result)

	err = userModel.DeleteOne("testName")

	if err != nil {
		t.Error(`Delete user failed`)
	}

}
