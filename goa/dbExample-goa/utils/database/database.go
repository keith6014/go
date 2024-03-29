package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func SqlLiteConnect(filename string) (*sql.DB, error) {
	return sql.Open("sqlite3", filename)
}
