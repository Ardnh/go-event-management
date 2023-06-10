package controller

import (
	userDomain "go/ems/domain/users"
	"go/ems/exception"
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
	var request userDomain.UserRegisterRequest

	// Internal  Server Error
	if err := c.ShouldBindJSON(&request); err != nil {

		exception.Response(c, http.StatusBadRequest, nil, err)
		return
	}

	errRegister := controller.Service.Register(c, &request)
	if errRegister != nil {

		exception.Response(c, http.StatusBadRequest, nil, errRegister)
		return
	}

	// Success
	exception.Response(c, http.StatusOK, nil, nil)
}

func (controller *UserControllerImpl) Login(c *gin.Context) {
	var request userDomain.UserLoginRequest

	// Json parse
	if err := c.ShouldBindJSON(&request); err != nil {

		exception.Response(c, http.StatusInternalServerError, nil, err)
		return
	}

	//	Validate request

	// Send request to users service
	user, errLogin := controller.Service.Login(c, &request)

	if errLogin != nil {

		exception.Response(c, http.StatusBadRequest, nil, errLogin)
		return
	}

	// Check password
	errCheckPassword := helper.CheckPassword(request.Password, user.Password)

	if errCheckPassword != nil {

		exception.Response(c, http.StatusBadRequest, nil, errCheckPassword)
		return
	}

	// Generate token
	id := strconv.Itoa(user.Id)
	token, errToken := helper.GenerateJWTKey(id, user.Role)

	if errToken != nil {

		exception.Response(c, http.StatusInternalServerError, nil, errToken)
		return
	}

	response := userDomain.UserResponseWithToken{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		UserName:  user.UserName,
		Email:     user.Email,
		RoleId:    user.RoleId,
		Role:      user.Role,
		Token:     token,
	}

	// Send response
	exception.Response(c, http.StatusOK, &response, nil)
}
