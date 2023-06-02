package models

import (
	"time"

	"gorm.io/gorm"
)

type Events struct {
	gorm.Model
	ID              uint       `gorm:"primaryKey; autoIncrement;"`
	Name            string     `gorm:"type:varchar(255); not null;" json:"name"`
	Description     string     `gorm:"type:varchar(255); not null;" json:"description"`
	StartDate       time.Time  `gorm:"not null;" json:"start_date"`
	EndDate         time.Time  `gorm:"not null;" json:"end_date"`
	RegistrationURL string     `gorm:"type:varchar(255); not null;" json:"registration_url"`
	Banner          string     `gorm:"type:varchar(255); not null;" json:"banner"`
	Address         string     `gorm:"type:varchar(255); not null;" json:"address"`
	Views           int        `gorm:"not null; default:0;" json:"views"`
	UserID          int        `gorm:"not null;" json:"user_id"`
	Users           Users      `gorm:"foreignKey:UserID;"`
	CategoryID      int        `gorm:"not null;" json:"category_id"`
	Categories      Categories `gorm:"foreignKey:CategoryID;"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
}
