package models

type NoteRequest struct {
	Title   *string
	Content *string
}

func NewNoteRequest(title *string, content *string) *NoteRequest {
	return &NoteRequest{
		Title:   title,
		Content: content,
	}
}
