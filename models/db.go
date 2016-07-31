package models

import (
	"database/sql"

	//SQLite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// NewDB creates a new database if one does not exist.
func NewDB(dataSource string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	sqlStmt := "create table if not exists employees(id integer not null primary key, name text, title text);"
	if _, err = db.Exec(sqlStmt); err != nil {
		panic(err)
	}

	return db, nil
}
