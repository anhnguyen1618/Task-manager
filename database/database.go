package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Initialize() *sql.DB {
	DBCon, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/taskManager")

	if err != nil {
		panic(err.Error())
	}

	// Test the connection to the database
	err = DBCon.Ping()
	if err != nil {
		panic(err.Error())
	}

	return DBCon
}
