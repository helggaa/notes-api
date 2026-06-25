package repository

import (
	"notes_api/internal/model"

	"gorm.io/gorm"
)

type NoteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) *NoteRepository {
	return &NoteRepository{
		db: db,
	}
}

func (r *NoteRepository) CreateNote(note *model.Note) error {
	return r.db.Create(note).Error
}

func (r *NoteRepository) GetNoteByID(id uint, userID uint) (*model.Note, error) {

	var note model.Note

	err := r.db.
		Where("id = ? AND user_id = ?", id, userID).
		First(&note).Error

	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *NoteRepository) GetAllNotes(
	userID uint,
	search string,
	status string,
	page int,
	limit int,
) ([]model.Note, error) {

	var notes []model.Note

	query := r.db.Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	offset := (page - 1) * limit

	err := query.
		Offset(offset).
		Limit(limit).
		Order("created_at DESC").
		Find(&notes).Error

	if err != nil {
		return nil, err
	}

	return notes, nil
}

func (r *NoteRepository) CountNotes(
	userID uint,
	search string,
	status string,
) (int64, error) {

	var total int64

	query := r.db.Model(&model.Note{}).
		Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Count(&total).Error

	return total, err
}

func (r *NoteRepository) UpdateNote(note *model.Note) error {
	return r.db.Save(note).Error
}

func (r *NoteRepository) DeleteNote(note *model.Note) error {
	return r.db.Delete(note).Error
}
