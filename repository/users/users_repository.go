package repository

import (
	domain "go/ems/domain/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx *gin.Context, tx *gorm.DB, req *domain.UserRegister) error
	FindByEmail(ctx *gin.Context, tx *gorm.DB, email string) (*domain.UserQueryResponse, error)
	FindById(ctx *gin.Context, tx *gorm.DB, id int) (*domain.User, error)
}
