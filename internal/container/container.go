package container

import (
	"database/sql"

	"github.com/Lucas4lves/go-notes-app/controllers"
	"github.com/Lucas4lves/go-notes-app/repositories"
	"github.com/Lucas4lves/go-notes-app/services"
)

type DependencyContainer struct {
	Driver          *sql.DB
	NotesRepo       *repositories.NoteRepository
	NotesService    *services.NoteService
	NotesController *controllers.NotesController
}

func NewDependencyContainer(driver *sql.DB) *DependencyContainer {

	notesRepo := repositories.NewNoteRepository(driver)
	notesService := services.NewNoteService(notesRepo)
	notesController := controllers.NewNotesController(notesService)
	return &DependencyContainer{
		Driver:          driver,
		NotesRepo:       notesRepo,
		NotesService:    notesService,
		NotesController: notesController,
	}
}
