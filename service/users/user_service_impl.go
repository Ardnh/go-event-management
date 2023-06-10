package service

import (
	"errors"
	"fmt"
	domain "go/ems/domain/users"
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

func (service *UserServiceImpl) Register(ctx *gin.Context, request *domain.UserRegisterRequest) error {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	_, errFindByEmail := service.Repository.FindByEmail(ctx, tx, request.Email)

	if errors.Is(errFindByEmail, gorm.ErrRecordNotFound) {
		hashPassword, errGenerateHash := helper.GenerateHashPassword(request.Password)
		if errGenerateHash != nil {
			fmt.Println("errGenerateHash")
			return errGenerateHash
		}

		fmt.Println(&request)

		userData := &domain.UserRegister{
			FirstName: request.FirstName,
			LastName:  request.LastName,
			Username:  request.UserName,
			Email:     request.Email,
			Password:  hashPassword,
			RoleID:    20,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		errRegister := service.Repository.Register(ctx, tx, userData)

		if errRegister != nil {
			return errRegister
		}
		return nil
	} else {
		errText := fmt.Sprintf("users with %s already exist", request.Email)
		return errors.New(errText)
	}
}

func (service *UserServiceImpl) Login(ctx *gin.Context, request *domain.UserLoginRequest) (domain.UserResponse, error) {
	errValidate := service.Validate.Struct(request)
	helper.PanicIfError(errValidate)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	user, err := service.Repository.FindByEmail(ctx, tx, request.Email)

	if err != nil {
		return domain.ToUserResponse(user), err
	}

	return domain.ToUserResponse(user), nil
}
