package services

import (
	"errors"
	"log"

	"github.com/Lucas4lves/go-notes-app/models"
	"github.com/Lucas4lves/go-notes-app/repositories"
)

type NoteService struct {
	Repo *repositories.NoteRepository
}

func NewNoteService(repo *repositories.NoteRepository) *NoteService {
	return &NoteService{
		Repo: repo,
	}
}

func (ns *NoteService) Create(n *models.Note) (int64, error) {

	if len(n.Title) == 0 {
		return -1, errors.New("ERROR: Titles must not be empty")
	}

	if len(n.Content) == 0 {
		return -1, errors.New("ERROR: Content must not be empty")
	}

	id, err := ns.Repo.Insert(n)

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return -1, err
	}

	return id, nil
}

func (ns *NoteService) Update(id int64, n *models.NoteRequest) error {
	err := ns.Repo.Update(id, n)

	if err != nil {
		log.Fatal("Error: ", err.Error())
		return err
	}

	return nil
}

func (ns *NoteService) GetById(id int) (*models.Note, error) {

	data, err := ns.Repo.SelectById(id)

	if err != nil {
		return nil, err
	}

	if data == nil {
		return nil, err
	}

	return data, nil
}

func (ns *NoteService) GetAll() ([]*models.Note, error) {
	data, err := ns.Repo.SelectAll()

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("empty data array")
	}

	return data, nil
}
