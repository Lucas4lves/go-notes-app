package repositories

import (
	"database/sql"
	"log"
	"strconv"
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

func (nr *NoteRepository) SelectById(id int) (*models.Note, error) {

	sql := "select * from notes where id = " + strconv.Itoa(id)

	row := nr.Driver.QueryRow(sql)

	var data *models.Note = &models.Note{}

	err := row.Scan(&data.ID, &data.Title, &data.Content, &data.CreatedAt, &data.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (nr *NoteRepository) SelectAll() ([]*models.Note, error) {
	rowsData := make([]*models.Note, 0)
	sql := "select * from notes"
	rows, err := nr.Driver.Query(sql)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var rowData *models.Note = &models.Note{}

		err := rows.Scan(&rowData.ID, &rowData.Title, &rowData.Content, &rowData.CreatedAt, &rowData.UpdatedAt)

		if err != nil {
			return nil, err
		}

		rowsData = append(rowsData, rowData)
	}

	return rowsData, nil
}
