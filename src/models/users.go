package models

import (
	"database/sql"
	"fmt"

	"github.com/anhnguyen300795/Task-manager/src/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	DB *sql.DB
}

func (model *Users) GetAll() []interfaces.UserInfo {
	db := model.DB
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	allUsers := []interfaces.UserInfo{}

	for rows.Next() {
		var username, email, password, role string

		err = rows.Scan(&username, &email, &password, &role)

		if err != nil {
			panic(err.Error())
		}

		allUsers = append(allUsers, interfaces.UserInfo{username, "", email, role})
	}
	return allUsers
}

func (model *Users) AddOne(userInfo *(interfaces.UserInfo)) *interfaces.UserInfo {
	db := model.DB

	password := userInfo.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}

	_, err = db.Exec("INSERT INTO users(username, email, password, role) VALUES(?, ?, ?, ?)", userInfo.UserName, userInfo.Email, hashedPassword, userInfo.Role)

	if err != nil {
		return nil
	}

	userInfo.Password = ""

	return userInfo
}

func (model *Users) GetOne(userName string) *interfaces.UserInfo {
	db := model.DB
	var email, role string

	err := db.QueryRow("SELECT email, role FROM users WHERE username=?", userName).
		Scan(&email, &role)

	fmt.Println(err)
	if err != nil {
		return nil
	}

	return &interfaces.UserInfo{userName, "", email, role}
}

func (model *Users) UpdateOne(userInfo *(interfaces.UserInfo)) *interfaces.UserInfo {
	db := model.DB

	_, err := db.Exec("UPDATE users set email=?, role=? WHERE username=?", userInfo.Email, userInfo.Role, userInfo.UserName)

	if err != nil {
		return nil
	}

	return userInfo
}

func (model *Users) DeleteOne(userName string) error {
	db := model.DB

	_, err := db.Exec("DELETE FROM users WHERE username=?", userName)

	return err
}

func (model *Users) CheckCredential(userInfo *(interfaces.UserInfo)) *(interfaces.UserInfo) {
	db := model.DB
	rawPassword := userInfo.Password
	var hashPassword string

	var userName, email, role string
	err := db.QueryRow("SELECT password, username, email, role FROM users WHERE username=?", userInfo.UserName).
		Scan(&hashPassword, &userName, &email, &role)

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))

	if err != nil {
		return nil
	}

	return &interfaces.UserInfo{userName, "", email, role}

}
