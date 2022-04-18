package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/Tubes3_13520112")

	if err != nil {
		log.Fatal(err)
	}

	return db
}