package service

import (
	"go/ems/domain"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Register(ctx *gin.Context, request *domain.UserRegisterRequest) (domain.UserResponse, error)
	Login(ctx *gin.Context, request *domain.UserLoginRequest) (domain.UserResponse, error)
}
