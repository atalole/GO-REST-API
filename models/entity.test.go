package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tests struct {
	ID        string `gorm:"primaryKey;unique;not null"`
	Name      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (entity *Tests) BeforeCreateTest(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Tests) BeforeUpdateTest(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
