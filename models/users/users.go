package users

import (
	"../../database"
	"../../interfaces"
	"golang.org/x/crypto/bcrypt"
)

func readAll() {
	db := database.DBCon
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var username, password, email, role string

		err = rows.Scan(&username, &password, &email, &role)

		if err != nil {
			panic(err.Error())
		}
	}
}

func AddOne(userInfo *(interfaces.UserInfo)) (string, error) {
	db := database.DBCon

	password := userInfo.Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "Create failed", err
	}

	_, err = db.Exec("INSERT INTO users(username, email, password, role) VALUES(?, ?, ?, ?)", userInfo.UserName, userInfo.Email, hashedPassword, "USER")

	if err != nil {
		return "Create failed", err
	}
	return "Create successfully", nil
}

func CheckCredential(userInfo *(interfaces.UserInfo)) *(interfaces.UserInfo) {
	db := database.DBCon
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
