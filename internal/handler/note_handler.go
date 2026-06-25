package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"notes_api/internal/dto"
	"notes_api/internal/service"
	"notes_api/internal/utils"
)

type NoteHandler struct {
	noteService *service.NoteService
}

func NewNoteHandler(noteService *service.NoteService) *NoteHandler {
	return &NoteHandler{
		noteService: noteService,
	}
}

func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req dto.CreateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	userID := c.MustGet("user_id").(uint)

	note, err := h.noteService.CreateNote(userID, req)

	if err != nil {
		utils.InternalServerError(c)
		return
	}

	utils.Success(
		c,
		http.StatusCreated,
		"Note created successfully",
		note,
	)
}

func (h *NoteHandler) GetAllNotes(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	search := c.DefaultQuery("search", "")
	status := c.DefaultQuery("status", "")

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	result, err := h.noteService.GetAllNotes(
		userID,
		search,
		status,
		page,
		limit,
	)

	if err != nil {
		utils.InternalServerError(c)
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Notes retrieved successfully",
		result,
	)
}

func (h *NoteHandler) GetNoteByID(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.BadRequest(c, "Invalid note id")
		return
	}

	note, err := h.noteService.GetNoteByID(
		uint(id),
		userID,
	)

	if err != nil {
		utils.NotFound(c, "Note not found")
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Note retrieved successfully",
		note,
	)
}

func (h *NoteHandler) UpdateNote(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.BadRequest(c, "Invalid note id")
		return
	}

	var req dto.UpdateNoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.BadRequest(c, err.Error())
		return
	}

	err = h.noteService.UpdateNote(
		uint(id),
		userID,
		req,
	)

	if err != nil {
		utils.NotFound(c, "Note not found")
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Note updated successfully",
		nil,
	)
}

func (h *NoteHandler) DeleteNote(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		utils.BadRequest(c, "Invalid note id")
		return
	}

	err = h.noteService.DeleteNote(
		uint(id),
		userID,
	)

	if err != nil {
		utils.NotFound(c, "Note not found")
		return
	}

	utils.Success(
		c,
		http.StatusOK,
		"Note deleted successfully",
		nil,
	)
}
