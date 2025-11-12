package repositories

import (
	"database/sql"

	"github.com/Lucas4lves/go-notes-app/models"
)

type NoteRepository struct {
	Driver *sql.DB
}

func NewNoteRepository(driver *sql.DB) *NoteRepository {
	return &NoteRepository{
		Driver: driver,
	}
}

func (ns *NoteRepository) Insert(note *models.Note) {

}
