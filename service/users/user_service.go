package service

import (
	domain "go/ems/domain/users"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(ctx *gin.Context, request *domain.UserRegisterRequest) error
	Login(ctx *gin.Context, request *domain.UserLoginRequest) (domain.UserResponse, error)
}
