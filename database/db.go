package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Couldn't connect to DataBase")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}
	createRefersTable := `
		CREATE TABLE IF NOT EXISTS referals (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
		    FOREIGN KEY(refery_id) REFERENCES users(id),
		    FOREIGN KEY(referred_id) REFERENCES users(id)
		)
	`
	_, err = DB.Exec(createRefersTable)
	if err != nil {
		panic(err)
	}
}
