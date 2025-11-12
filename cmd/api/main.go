package main

import (
	"github.com/Lucas4lves/go-notes-app/database"
)

func main() {
	db := database.InitDatabase("notes.db")

	defer db.Close()

}
