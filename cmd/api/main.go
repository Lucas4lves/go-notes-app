package main

import (
	"github.com/Lucas4lves/go-notes-app/database"
	"github.com/Lucas4lves/go-notes-app/internal/container"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.InitDatabase("notes.db")
	depsContainer := container.NewDependencyContainer(db)

	server := gin.Default()

	server.POST("/notes", depsContainer.NotesController.CreateNote)
	server.PUT("/notes/:id", depsContainer.NotesController.UpdateNote)

	defer db.Close()

	server.Run()
}
