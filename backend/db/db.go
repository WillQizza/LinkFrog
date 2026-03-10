package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() error {
	dsn := os.Getenv("DB_DSN")

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := DB.Ping(); err != nil {
		return err
	}
	return initializeTables()
}

func initializeTables() error {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id    INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE
	)`)

	if err != nil {
		return err
	}

	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS links (
		id    INT AUTO_INCREMENT PRIMARY KEY,
		owner INT NOT NULL,
		path  VARCHAR(255) NOT NULL UNIQUE,
		url   TEXT NOT NULL,
		FOREIGN KEY (owner) REFERENCES users(id)
	)`)

	return err
}
