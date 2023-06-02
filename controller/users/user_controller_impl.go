package controller

import (
	"fmt"
	"go/ems/domain"
	"go/ems/helper"
	service "go/ems/service/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	Service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &UserControllerImpl{
		Service: service,
	}
}

func (controller *UserControllerImpl) Register(c *gin.Context) {
	var request domain.UserRegisterRequest

	fmt.Println("register")

	// Internal  Server Error
	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	_, err := controller.Service.Register(c, &request)
	if err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	// Success
	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "Created",
	}
	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Login(c *gin.Context) {
	var request domain.UserLoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	user, errLogin := controller.Service.Login(c, &request)

	if errLogin != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   errLogin.Error(),
		}
		c.JSON(http.StatusBadRequest, webResponse)
		return
	}

	id := strconv.Itoa(user.Id)
	token, errToken := helper.GenerateJWTKey(id)

	if errToken != nil {
		webResponse := domain.WebResponse{
			Code:   http.StatusInternalServerError,
			Status: "Internal Server Error",
			Data:   errToken,
		}
		c.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	response := domain.UserResponseWithToken{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	webResponse := domain.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   response,
	}
	c.JSON(http.StatusOK, webResponse)
}
