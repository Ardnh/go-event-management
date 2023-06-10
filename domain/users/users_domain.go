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
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname" validate:"required"`
	UserName  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

type UserRegister struct {
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	RoleID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Email     string
}

type UserResponseWithToken struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Token     string `json:"token"`
	RoleId    int    `json:"role_id"`
	Role      string `json:"role"`
}

type UserResponse struct {
	Id        int
	FirstName string
	LastName  string
	UserName  string
	Email     string
	Password  string
	RoleId    int
	Role      string
}

type UserQueryResponse struct {
	Id        int
	FirstName string
	LastName  string
	Username  string
	Email     string
	Password  string
	Active    bool
	RoleId    int
	Role      string
}
