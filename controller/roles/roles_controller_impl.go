package controller

import (
	"go/ems/domain"
	service "go/ems/service/roles"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RolesControllerImpl struct {
	Service service.RolesService
}

func NewRolesController(service service.RolesService) RolesController {
	return &RolesControllerImpl{
		Service: service,
	}
}

func (controller *RolesControllerImpl) Create(c *gin.Context) {
	var request domain.RolesCreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, &webResponse)
		return
	}

	role, err := controller.Service.Create(c, request)

	if err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}

		c.JSON(http.StatusBadRequest, &webResponse)
		return
	}

	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "Successfully create new role",
		Data:   role,
	}

	c.JSON(http.StatusOK, &webResponse)
}

func (controller *RolesControllerImpl) Update(c *gin.Context) {
	var request domain.RolesUpdateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, &webResponse)
		return
	}

	role, err := controller.Service.Update(c, request)

	if err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   err.Error(),
		}

		c.JSON(http.StatusBadRequest, &webResponse)
		return
	}

	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "Successfully Updated Roles",
		Data:   role,
	}

	c.JSON(http.StatusOK, &webResponse)
}

func (controller *RolesControllerImpl) Delete(c *gin.Context) {
	var request domain.RolesDeleteRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	if err := controller.Service.Delete(c, request.Id); err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "Successfully Delete Roles",
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *RolesControllerImpl) FindById(c *gin.Context) {
	var id int
	if err := c.ShouldBindUri(&id); err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	roles, err := controller.Service.FindById(c, id)
	if err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "Successfully Get Role",
		Data:   roles,
	}

	c.JSON(http.StatusOK, webResponse)

}

func (controller *RolesControllerImpl) FindAll(c *gin.Context) {
	pageString := c.DefaultQuery("page", "1")
	limitString := c.DefaultQuery("limit", "100")
	sortString := c.DefaultQuery("sort", "asc")

	page, _ := strconv.Atoi(pageString)
	limit, _ := strconv.Atoi(limitString)

	data := domain.RolesFindAllRequest{
		Page:  page,
		Limit: limit,
		Sort:  sortString,
	}

	roles, err := controller.Service.FindAll(c, data)

	if err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	webResponse := domain.WebResponseWithPagination{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Page:   page,
		Limit:  limit,
		Sort:   sortString,
		Data:   roles,
	}

	c.JSON(http.StatusOK, webResponse)
}
