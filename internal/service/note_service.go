package service

import (
	"math"

	"notes_api/internal/dto"
	"notes_api/internal/model"
	"notes_api/internal/repository"
)

type NoteService struct {
	noteRepo *repository.NoteRepository
}

func NewNoteService(noteRepo *repository.NoteRepository) *NoteService {
	return &NoteService{
		noteRepo: noteRepo,
	}
}

func (s *NoteService) CreateNote(
	userID uint,
	req dto.CreateNoteRequest,
) (*model.Note, error) {
	status := req.Status

	if status == "" {
		status = "active"
	}

	note := &model.Note{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
		Status:  status,
	}

	err := s.noteRepo.CreateNote(note)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (s *NoteService) GetNoteByID(
	id uint,
	userID uint,
) (*model.Note, error) {
	return s.noteRepo.GetNoteByID(id, userID)
}

func (s *NoteService) GetAllNotes(
	userID uint,
	search string,
	status string,
	page int,
	limit int,
) (*dto.NoteListResponse, error) {
	notes, err := s.noteRepo.GetAllNotes(
		userID,
		search,
		status,
		page,
		limit,
	)

	if err != nil {
		return nil, err
	}

	total, err := s.noteRepo.CountNotes(
		userID,
		search,
		status,
	)

	if err != nil {
		return nil, err
	}

	var response []dto.NoteResponse

	for _, note := range notes {
		response = append(response, dto.NoteResponse{
			ID:        note.ID,
			Title:     note.Title,
			Content:   note.Content,
			Status:    note.Status,
			CreatedAt: note.CreatedAt,
			UpdatedAt: note.UpdatedAt,
		})

	}

	if limit <= 0 {
		limit = 10
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &dto.NoteListResponse{
		Data:       response,
		Page:       page,
		Limit:      limit,
		TotalData:  total,
		TotalPages: totalPages,
	}, nil
}

func (s *NoteService) UpdateNote(
	id uint,
	userID uint,
	req dto.UpdateNoteRequest,
) error {
	note, err := s.noteRepo.GetNoteByID(id, userID)

	if err != nil {
		return err
	}

	note.Title = req.Title
	note.Content = req.Content
	note.Status = req.Status

	return s.noteRepo.UpdateNote(note)

}

func (s *NoteService) DeleteNote(
	id uint,
	userID uint,
) error {
	note, err := s.noteRepo.GetNoteByID(id, userID)

	if err != nil {
		return err
	}

	return s.noteRepo.DeleteNote(note)
}
