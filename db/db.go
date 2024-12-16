package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbInstance *sql.DB

func InitDB() {
	DBInstance, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		log.Fatal(err)
	}
	defer DBInstance.Close()
}
