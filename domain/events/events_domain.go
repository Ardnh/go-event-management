package domain

import "time"

type Events struct {
	Id              int
	Name            string
	Description     string
	StartDate       string
	EndDate         string
	RegistrationUrl string
	Banner          string
	Address         string
	Views           int
	UserId          int
	CategoryId      int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type EventsCreateRequest struct {
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"required"`
	StartDate       string `json:"start_date" validate:"required"`
	EndDate         string `json:"end_date" validate:"required"`
	RegistrationUrl string `json:"registration_url" validate:"required"`
	Banner          string `json:"banner" validate:"required"`
	Address         string `json:"address" validate:"required"`
	UserId          int    `json:"user_id" validate:"required"`
	CategoryId      int    `json:"category_id" validate:"required"`
}

type EventsUpdateRequest struct {
	Id              int    `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:"required"`
	StartDate       string `json:"start_date" validate:"required"`
	EndDate         string `json:"end_date" validate:"required"`
	RegistrationUrl string `json:"registration_url" validate:"required"`
	Banner          string `json:"banner" validate:"required"`
	Address         string `json:"address" validate:"required"`
	UserId          int    `json:"user_id" validate:"required"`
	CategoryId      int    `json:"category_id" validate:"required"`
}

type EventsFindAllRequest struct {
	Page  int    `json:"page" validate:"required"`
	Limit int    `json:"limit" validate:"required"`
	Sort  string `json:"sort" validate:"required"`
}

type EventsDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

// Response ===============

type EventsQueryResponse struct {
	Id              int
	Name            string
	Description     string
	StartDate       string
	EndDate         string
	RegistrationUrl string
	Banner          string
	Address         string
	Views           int
	UserId          int
	Username        string
	CategoryId      int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}
