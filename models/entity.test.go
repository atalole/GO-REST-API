package model

import (
	"time"
)

type Tests struct {
	ID        string `gorm:"primaryKey";serializer;unique; not null`
	Name      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
