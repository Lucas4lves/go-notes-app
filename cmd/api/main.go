package main

import (
	"github.com/Lucas4lves/go-notes-app/database"
	"github.com/Lucas4lves/go-notes-app/internal/container"
)

func main() {
	db := database.InitDatabase("notes.db")

	depsContainer := container.NewDependencyContainer(db)

	defer db.Close()

}
