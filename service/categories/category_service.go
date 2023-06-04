package service

import (
	domain "go/ems/domain/category"

	"github.com/gin-gonic/gin"
)

type CategoryService interface {
	Create(ctx *gin.Context, request *domain.CategoryCreateRequest) (*domain.CategoryResponse, error)
	Update(ctx *gin.Context, request *domain.CategoryUpdateRequest) (*domain.CategoryResponse, error)
	Delete(ctx *gin.Context, request int) error
	FindById(ctx *gin.Context, request int) (*domain.CategoryResponse, error)
	FindAll(ctx *gin.Context, request *domain.CategoryFindAllRequest) ([]*domain.CategoryResponse, error)
}
