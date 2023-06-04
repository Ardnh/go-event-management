package domain

import "time"

type Category struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type CategoryFindAllRequest struct {
	Page  int    `json:"page" validate:"required"`
	Limit int    `json:"limit" validate:"required"`
	Sort  string `json:"sort" validate:"required"`
}

type CategoryDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

// Service level
type CategoryResponse struct {
	Id        int
	Name      string
	CreatedAt string
	UpdatedAt string
}
