package models

import (
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey; autoIncrement;"`
	Name      string    `gorm:"type:varchar(255); not null;" json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
