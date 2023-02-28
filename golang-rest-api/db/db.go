package db

import (
	"database/sql"
)

func OpenConnection() *sql.DB {
	psqlInfo := "postgres://postgres:password@0.0.0.0:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
