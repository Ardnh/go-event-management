package models

import (
	"time"

	"gorm.io/gorm"
)

type Advertisements struct {
	gorm.Model
	ID          int       `gorm:"primaryKey; autoIncrement;" json:"id"`
	UserID      int       `gorm:"not null;" json:"user_id"`
	Users       Users     `gorm:"foreignKey:UserID;"`
	EventID     int       `gorm:"not null;" json:"eventID"`
	Events      Events    `gorm:"foreignKey:EventID;"`
	FacebookUrl string    `gorm:"type:varchar(255); null;" json:"facebook_url"`
	TwitterUrl  string    `gorm:"type:varchar(255); null;" json:"twitter_url"`
	BannerUrl   string    `gorm:"type:varchar(255); null;" json:"banner_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
