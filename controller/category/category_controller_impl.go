package controller

import (
	domain "go/ems/domain/category"
	"go/ems/exception"
	service "go/ems/service/categories"
	"net/http"
	"strconv"

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

func (controller *CategoryControllerImpl) Update(c *gin.Context) {
	var request domain.CategoryUpdateRequest

	// Internal Server Error
	if err := c.ShouldBindJSON(&request); err != nil {
		exception.Response(c, http.StatusInternalServerError, nil, err)
		return
	}

	// Bad Request
	data, err := controller.Service.Update(c, &request)
	if err != nil {
		exception.Response(c, http.StatusBadRequest, nil, err)
		return
	}

	// Success
	exception.Response(c, http.StatusOK, data, nil)
}

func (controller *CategoryControllerImpl) Delete(c *gin.Context) {
	var request domain.CategoryDeleteRequest

	// Internal Server Error
	if err := c.ShouldBindJSON(&request); err != nil {
		exception.Response(c, http.StatusInternalServerError, nil, err)
		return
	}

	// Bad Request
	err := controller.Service.Delete(c, request.Id)
	if err != nil {
		exception.Response(c, http.StatusBadRequest, nil, err)
		return
	}

	// Success
	exception.Response(c, http.StatusOK, nil, nil)
}

func (controller *CategoryControllerImpl) FindAll(c *gin.Context) {

	pageString := c.DefaultQuery("page", "1")
	limitString := c.DefaultQuery("limit", "100")
	sortString := c.DefaultQuery("sort", "asc")

	page, _ := strconv.Atoi(pageString)
	limit, _ := strconv.Atoi(limitString)

	data := domain.CategoryFindAllRequest{
		Page:  page,
		Limit: limit,
		Sort:  sortString,
	}

	category, err := controller.Service.FindAll(c, &data)

	if err != nil {
		exception.Response(c, http.StatusBadRequest, nil, err)
		return
	}

	webResponse := domain.WebResponseWithPagination{
		Code:   http.StatusOK,
		Status: "Successfully Get Roles",
		Page:   page,
		Limit:  limit,
		Data:   category,
	}

	c.JSON(http.StatusOK, &webResponse)

}

func (controller *CategoryControllerImpl) FindById(c *gin.Context) {

	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		exception.Response(c, http.StatusInternalServerError, nil, err)
		return
	}

	category, errFindById := controller.Service.FindById(c, id)
	if errFindById != nil {
		exception.Response(c, http.StatusBadRequest, nil, errFindById)
		return
	}

	exception.Response(c, http.StatusBadRequest, category, nil)
}
