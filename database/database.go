package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Initialize() *sql.DB {
	DBCon, err := sql.Open("mysql", "test:test@tcp(35.198.190.39)/taskmanager")

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
