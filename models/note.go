package models

import "time"

type Note struct {
	ID        *int
	Title     string
	Content   string
	CreatedAt string
	UpdatedAt string
}

func NewNote(title string, content string) *Note {
	return &Note{
		ID:        nil,
		Title:     title,
		Content:   content,
		CreatedAt: "",
		UpdatedAt: "",
	}
}

func (n *Note) Update(newTitle, newContent *string) {
	if newTitle != nil {
		n.Title = *newTitle
	}

	if newContent != nil {
		n.Content = *newContent
	}

	n.UpdatedAt = time.Now().Format(time.RFC3339)
}
