package main

import (
	"database/sql"
)

func NewDB(dbDriver string, dbSource string) *sql.DB {
	db, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		panic(err)
	}

	return db
}
