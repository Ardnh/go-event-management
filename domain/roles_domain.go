package domain

import (
	"time"
)

type Roles struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// =======================================
// REQUEST
// =======================================

// Service level
type RolesCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type RolesUpdateRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type RolesFindAllRequest struct {
	Page  int    `json:"page" validate:"required"`
	Limit int    `json:"limit" validate:"required"`
	Sort  string `json:"sort" validate:"required"`
}

type RolesDeleteRequest struct {
	Id int `json:"id" validate:"required"`
}

// =======================================
// RESPONSE
// =======================================

// Service level
type RolesResponse struct {
	Id        int
	Name      string
	CreatedAt string
	UpdatedAt string
}
