package controller

import (
	domain "go/ems/domain/category"
	"go/ems/exception"
	service "go/ems/service/categories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryControllerImpl struct {
	Service service.CategoryService
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Service: service,
	}
}

func (controller *CategoryControllerImpl) Create(c *gin.Context) {
	var request domain.CategoryCreateRequest

	// Internal Server Error
	if err := c.ShouldBindJSON(&request); err != nil {
		exception.Response(c, http.StatusInternalServerError, nil, err)
		return
	}

	// Bad Request
	data, err := controller.Service.Create(c, &request)
	if err != nil {
		exception.Response(c, http.StatusBadRequest, nil, err)
		return
	}

	// Success
	exception.Response(c, http.StatusOK, data, nil)
}
