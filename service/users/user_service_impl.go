package service

import (
	"fmt"
	"go/ems/domain"
	"go/ems/helper"
	repository "go/ems/repository/users"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Repository repository.UserRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewUserService(repository repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *UserServiceImpl) Register(ctx *gin.Context, request *domain.UserRegisterRequest) (domain.UserResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, errFindByUsername := service.Repository.FindByUsername(ctx, tx, request.UserName)

	if errFindByUsername != nil {
		hashPassword, errGenerateHash := helper.GenerateHashPassword(request.Password)
		if errGenerateHash != nil {
			fmt.Println("errGenerateHash")
			return domain.ToUserResponse(user), errGenerateHash
		}

		userData := domain.User{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Username:  request.UserName,
			Email:     request.Email,
			Password:  hashPassword,
			RoleID:    20,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		errRegister := service.Repository.Register(ctx, tx, &userData)

		if errRegister != nil {
			fmt.Println("errRegister")
			return domain.ToUserResponse(&userData), errRegister
		}
		return domain.ToUserResponse(&userData), nil
	} else {
		return domain.ToUserResponse(user), err
	}
}

func (service *UserServiceImpl) Login(ctx *gin.Context, request *domain.UserLoginRequest) (domain.UserResponse, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindByUsername(ctx, tx, request.Username)

	if err != nil {
		return domain.ToUserResponse(user), err
	}

	return domain.ToUserResponse(user), nil
}
