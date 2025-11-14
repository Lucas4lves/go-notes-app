package main

import (
	"github.com/Lucas4lves/go-notes-app/database"
	"github.com/Lucas4lves/go-notes-app/internal/container"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	db := database.InitDatabase("notes.db")
	depsContainer := container.NewDependencyContainer(db)

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	server.POST("/notes", depsContainer.NotesController.CreateNote)
	server.PUT("/notes/:id", depsContainer.NotesController.UpdateNote)
	server.GET("/notes/:id", depsContainer.NotesController.GetNoteById)
	server.GET("/notes", depsContainer.NotesController.GetAllNotes)

	defer db.Close()

	server.Run()
}
