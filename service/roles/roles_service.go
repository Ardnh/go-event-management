package service

import (
	"go/ems/domain"

	"github.com/gin-gonic/gin"
)

type RolesService interface {
	Create(ctx *gin.Context, request domain.RolesCreateRequest) (domain.RolesResponse, error)
	Update(ctx *gin.Context, request domain.RolesUpdateRequest) (domain.RolesResponse, error)
	Delete(ctx *gin.Context, request int) error
	FindById(ctx *gin.Context, request int) (domain.RolesResponse, error)
	FindAll(ctx *gin.Context, request domain.RolesFindAllRequest) ([]domain.RolesResponse, error)
}
