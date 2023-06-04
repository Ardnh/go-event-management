package repository

import (
	domain "go/ems/domain/category"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindByName(ctx *gin.Context, tx *gorm.DB, name string) bool
	FindById(ctx *gin.Context, tx *gorm.DB, id int) (*domain.Category, error)
	FindAll(ctx *gin.Context, tx *gorm.DB, pageNumber int, pageSize int, orderBy string) ([]*domain.Category, error)
	Create(ctx *gin.Context, tx *gorm.DB, data *domain.Category) (*domain.Category, error)
	Update(ctx *gin.Context, tx *gorm.DB, data *domain.Category) (*domain.Category, error)
	Delete(ctx *gin.Context, tx *gorm.DB, id int) error
}
