package database

import (
	"strings"

	"github.com/Lucas4lves/go-notes-app/models"
)

func InsertQueryBuilder(payload *models.Note) string {
	var b strings.Builder
	var components []string

	if payload.Title != "" {
		components = append(components, "title")
	}

	if payload.Content != "" {
		components = append(components, "content")
	}

	if payload.CreatedAt != "" {
		components = append(components, "created_at")
	}

	if payload.UpdatedAt != "" {
		components = append(components, "updated_at")
	}

	b.WriteString("insert into notes")
	b.WriteString("values(")

	for i := range components {
		b.WriteString(components[i])
		if i <= len(components)-1 {
			b.WriteString(",")
		}
	}

	b.WriteString(")")

	b.WriteString(")")
	b.WriteString(");")

	return b.String()
}
