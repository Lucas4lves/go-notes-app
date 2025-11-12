package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/Lucas4lves/go-notes-app/database"
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

func (ns *NoteRepository) Insert(note *models.Note) error {

	tx, err := ns.Driver.Begin()

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(database.InsertNoteStmt)

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return err
	}

	defer stmt.Close()

	note.CreatedAt = time.Now().Format(time.RFC3339)
	note.UpdatedAt = time.Now().Format(time.RFC3339)

	_, err = stmt.Exec(note.Title, note.Content, note.CreatedAt, note.UpdatedAt)

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error: ", err.Error())
		return err
	}

	return nil
}
