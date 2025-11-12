package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDatabase(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)

	if err != nil {
		log.Fatal("ERROR: Unable to initialize database -> ", err)
	}

	_, err = db.Exec(CreateNotesTableQuery)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}

	return db
}
