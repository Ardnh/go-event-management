package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        int       `gorm:"primaryKey; autoIncrement;"`
	FirstName string    `gorm:"type:varchar(255); not null;" json:"firstname"`
	LastName  string    `gorm:"type:varchar(255); not null;" json:"lastname"`
	Username  string    `gorm:"type:varchar(255); not null;" json:"username"`
	Email     string    `gorm:"type:varchar(255); not null;" json:"email"`
	Password  string    `gorm:"type:varchar(255); not null;" json:"password"`
	Active    bool      `gorm:"not null; default:false;" json:"active"`
	RoleID    int       `gorm:"not null;" json:"role_id"`
	Roles     Roles     `gorm:"foreignKey:RoleID;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
