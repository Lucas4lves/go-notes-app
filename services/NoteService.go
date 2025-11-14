package services

import (
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

func (ns *NoteService) GetById(id int64) ([]*models.Note, error) {

	data, err := ns.Repo.SelectById(id)

	if err != nil {
		log.Println("Error: ", err.Error())
		return nil, err
	}

	return data, nil
}
