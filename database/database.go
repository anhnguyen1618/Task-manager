package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DBCon *sql.DB
var err error

func Initialize() {
	DBCon, err = sql.Open("mysql", "root@tcp(127.0.0.1:3306)/taskManager")

	if err != nil {
		panic(err.Error())
	}

	// Test the connection to the database
	err = DBCon.Ping()
	if err != nil {
		panic(err.Error())
	}
}

func Close() {
	// sql.DB should be long lived "defer" closes it once this function ends
	DBCon.Close()
}
