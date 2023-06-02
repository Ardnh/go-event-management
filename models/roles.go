package models

import (
	"time"
)

type Roles struct {
	ID        uint      `gorm:"primaryKey; autoIncrement;"`
	Name      string    `gorm:"type:varchar(255); not null;" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime;" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;" json:"updated_at"`
}
