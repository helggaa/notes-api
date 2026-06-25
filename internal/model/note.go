package model

import (
	"time"

	"gorm.io/gorm"
)

type Note struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint

	User      User   `gorm:"foreignKey:UserID"`
	Title     string `gorm:"size:255;not null"`
	Content   string `gorm:"type:text;not null"`
	Status    string `gorm:"type:enum('active','archived');default:'active'"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
