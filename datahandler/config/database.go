package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://maulanazn:t00r123@localhost/paybook?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(0)
	db.SetMaxOpenConns(50)

	return db
}
