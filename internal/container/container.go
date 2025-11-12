package container

import (
	"database/sql"

	"github.com/Lucas4lves/go-notes-app/repositories"
)

type DependencyContainer struct {
	Driver    *sql.DB
	NotesRepo *repositories.NoteRepository
}

func NewDependencyContainer(driver *sql.DB) *DependencyContainer {

	notesRepo := repositories.NewNoteRepository(driver)

	return &DependencyContainer{
		Driver:    driver,
		NotesRepo: notesRepo,
	}
}
