package repository

import (
	"go/ems/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepository interface {
	Register(ctx *gin.Context, tx *gorm.DB, req *domain.User) error
	FindByUsername(ctx *gin.Context, tx *gorm.DB, username string) (*domain.User, error)
	FindById(ctx *gin.Context, tx *gorm.DB, id int) (*domain.User, error)
}
