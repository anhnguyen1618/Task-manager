package users

import (
	"../../database"
	"../../interfaces"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func readAll() {
	db := database.DBCon
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id int
		var username string
		var password string
		var email string
		err = rows.Scan(&id, &username, &password, &email)

		if err != nil {
			panic(err.Error())
		}

		fmt.Println(id)
		fmt.Println(username)
		fmt.Println(password)
		fmt.Println(email)
	}
}

func AddOne(userInfo *(interfaces.UserInfo)) (string, error) {
	db := database.DBCon

	var username string
	err := db.QueryRow("SELECT username FROM users WHERE username=? OR email=?", userInfo.UserName, userInfo.Email).Scan(&username)

	if err == sql.ErrNoRows {
		password := userInfo.Password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		result, err := db.Exec("INSERT INTO users(username, email, password) VALUES(?, ?, ?)", userInfo.UserName, userInfo.Email, hashedPassword)

		fmt.Println(result.LastInsertId())

		if err != nil {
			return "Create failed", err
		}
		return "Create successfully", nil
	}

	return "Create failed", errors.New("this username/email has been used")
}

func CheckCredential(userInfo *(interfaces.UserInfo)) *(interfaces.UserInfo) {
	db := database.DBCon
	rawPassword := userInfo.Password
	var hashPassword string

	var id int
	var userName, email string
	err := db.QueryRow("SELECT password, id, username, email FROM users WHERE username=?", userInfo.UserName).Scan(&hashPassword, &id, &userName, &email)

	err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(rawPassword))

	if err != nil {
		return nil
	}

	return &interfaces.UserInfo{id, userName, "", email}

}
