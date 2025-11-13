package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Lucas4lves/go-notes-app/models"
	"github.com/Lucas4lves/go-notes-app/services"
	"github.com/gin-gonic/gin"
)

type NotesController struct {
	Service *services.NoteService
}

func NewNotesController(ns *services.NoteService) *NotesController {
	return &NotesController{
		Service: ns,
	}
}

func (nc *NotesController) CreateNote(ctx *gin.Context) {
	var newNote *models.Note
	if err := ctx.ShouldBindJSON(&newNote); err != nil {
		ctx.JSON(400, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	newNote.CreatedAt = time.Now().Format(time.RFC3339)
	newNote.UpdatedAt = time.Now().Format(time.RFC3339)

	id, err := nc.Service.Create(newNote)

	if err != nil {
		ctx.JSON(500, gin.H{
			"ERROR": err.Error(),
		})
		return
	}

	ctx.JSON(201, gin.H{
		"message": fmt.Sprintf("note with id %d created successfully!", id),
	})

}

func (nc *NotesController) UpdateNote(ctx *gin.Context) {
	id := ctx.Param("id")

	var updateReqContent *models.NoteRequest

	if err := ctx.ShouldBindJSON(&updateReqContent); err != nil {
		ctx.JSON(400, gin.H{
			"ERROR": err.Error(),
		})

		return
	}

	newId, _ := strconv.ParseInt(id, 10, 64)

	err := nc.Service.Update(newId, updateReqContent)

	if err != nil {
		ctx.JSON(500, gin.H{
			"ERROR": err.Error(),
		})

		return
	}

	ctx.JSON(200, gin.H{
		"message": "Note updated successfully",
	})

}
