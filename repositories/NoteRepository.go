package repositories

import (
	"database/sql"
	"log"
	"strings"
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

func (ns *NoteRepository) Insert(note *models.Note) (int64, error) {

	tx, err := ns.Driver.Begin()

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return 0, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(database.InsertNoteStmt)

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return 0, err
	}

	defer stmt.Close()

	note.CreatedAt = time.Now().Format(time.RFC3339)
	note.UpdatedAt = time.Now().Format(time.RFC3339)

	res, err := stmt.Exec(note.Title, note.Content, note.CreatedAt, note.UpdatedAt)

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal("Error: ", err.Error())
		return 0, err
	}

	return res.LastInsertId()
}

func (nr *NoteRepository) Update(id int64, payload *models.NoteRequest) error {
	queryComponents := []string{}
	arguments := []interface{}{}

	if payload.Title != nil {
		queryComponents = append(queryComponents, "title = ?")
		arguments = append(arguments, *payload.Title)
	}

	if payload.Content != nil {
		queryComponents = append(queryComponents, "content = ?")
		arguments = append(arguments, *payload.Content)
	}

	if len(queryComponents) == 0 {
		return nil
	}

	query := "UPDATE notes SET " + strings.Join(queryComponents, ", ") + ", updated_at = ? WHERE id = ?"
	arguments = append(arguments, time.Now().Format(time.RFC3339), id)

	_, err := nr.Driver.Exec(query, arguments...)

	return err
}

func (nr *NoteRepository) SelectById(id int64) (*models.Note, error) {
	var data *models.Note = &models.Note{}

	tx, err := nr.Driver.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("select * from notes")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&data.ID, &data.Title, &data.Content, &data.CreatedAt, &data.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return data, nil

}
