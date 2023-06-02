package domain

import "time"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	RoleID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRegisterRequest struct {
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
}

type UserLoginRequest struct {
	Username string
	Password string
}

type UserUpdateRequest struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Email     string
}

type UserResponseWithToken struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Token     string
	CreatedAt string
	UpdatedAt string
}

type UserResponse struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Email     string
	CreatedAt string
	UpdatedAt string
}
