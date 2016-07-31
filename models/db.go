package models

import (
	"database/sql"

	"gopkg.in/gorp.v1"

	//SQLite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// NewDB creates a new database if one does not exist.
func NewDB(dataSource string) (*gorp.DbMap, error) {
	db, err := sql.Open("sqlite3", dataSource)
	if err != nil {
		return nil, err
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// Employees
	dbmap.AddTableWithName(Employee{}, "employees").SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbmap, nil
}
