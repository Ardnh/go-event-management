package repository

import (
	"go/ems/domain"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RolesRepository interface {
	FindByName(ctx *gin.Context, tx *gorm.DB, name string) bool
	FindById(ctx *gin.Context, tx *gorm.DB, id int) (domain.Roles, error)
	FindAll(ctx *gin.Context, tx *gorm.DB, pageNumber int, pageSize int, orderBy string) ([]domain.Roles, error)
	Create(ctx *gin.Context, tx *gorm.DB, data domain.Roles) (domain.Roles, error)
	Update(ctx *gin.Context, tx *gorm.DB, data domain.Roles) (domain.Roles, error)
	Delete(ctx *gin.Context, tx *gorm.DB, id int) error
}
