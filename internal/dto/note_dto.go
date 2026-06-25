package dto

import "time"

type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status" binding:"omitempty,oneof=active archived"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  string `json:"status" binding:"required,oneof=active archived"`
}

type NoteResponse struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NoteListResponse struct {
	Data       []NoteResponse `json:"data"`
	Page       int            `json:"page"`
	Limit      int            `json:"limit"`
	TotalData  int64          `json:"total_data"`
	TotalPages int            `json:"total_pages"`
}
