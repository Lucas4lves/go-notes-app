package models

type Note struct {
	ID        *int
	Title     string
	Content   string
	CreatedAt *string
	UpdatedAt *string
}

func NewNote(title string, content string) *Note {
	return &Note{
		ID:        nil,
		Title:     title,
		Content:   content,
		CreatedAt: nil,
		UpdatedAt: nil,
	}
}

func (n *Note) UpdateTitle(newTitle string) {
	n.Title = newTitle
}

func (n *Note) UpdateContent(newContent string) {
	n.Content = newContent
}
