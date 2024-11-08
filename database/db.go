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
			password TEXT NOT NULL,
			FOREIGN KEY (referrerId) REFERENCES users(id),
		)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err)
	}
	createCodesTable := `
		CREATE TABLE IF NOT EXISTS codes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			code TEXT NOT NULL,
			TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (userId) REFERENCES users(id)
		) ON SCHEDULE EVERY 1 HOUR 
		DO
			DELETE FROM codes WHERE timestamp < NOW() - INTERVAL 1 DAY
	`
	_, err = DB.Exec(createCodesTable)
	if err != nil {
		panic(err)
	}
}
