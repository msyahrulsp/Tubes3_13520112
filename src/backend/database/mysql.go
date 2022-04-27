package database

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "AMOYSC0JU4:Lc98Bq6A1y@tcp(remotemysql.com:3306)/AMOYSC0JU4")

	if err != nil {
		log.Fatal(err)
	}

	return db
}