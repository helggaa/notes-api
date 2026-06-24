package model

import "time"

type Note struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Title     string `gorm:"size:255;not null"`
	Content   string `gorm:"type:text;not null"`
	Status    string `gorm:"type:enum('active','archived');default:'active'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
