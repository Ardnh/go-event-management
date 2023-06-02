package service

import (
	"errors"
	"go/ems/domain"
	"go/ems/helper"
	repository "go/ems/repository/roles"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type RolesServiceImpl struct {
	Repository repository.RolesRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewRolesService(repository repository.RolesRepository, db *gorm.DB, validate *validator.Validate) RolesService {
	return &RolesServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *RolesServiceImpl) Create(ctx *gin.Context, request domain.RolesCreateRequest) (domain.RolesResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	role := domain.Roles{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	isAvailable := service.Repository.FindByName(ctx, tx, request.Name)

	if isAvailable {
		result, _ := service.Repository.Create(ctx, tx, role)
		return domain.ToRolesResponse(result), nil
	} else {
		return domain.ToRolesResponse(role), errors.New("roles already exist")
	}
}

func (service *RolesServiceImpl) Update(ctx *gin.Context, request domain.RolesUpdateRequest) (domain.RolesResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	role, err := service.Repository.FindById(ctx, tx, request.Id)

	if err != nil {
		return domain.ToRolesResponse(role), errors.New(err.Error())
	}

	role.Name = request.Name
	role.UpdatedAt = time.Now()

	updatedRole, err := service.Repository.Update(ctx, tx, role)

	if err != nil {
		return domain.ToRolesResponse(updatedRole), errors.New(err.Error())
	}

	return domain.ToRolesResponse(updatedRole), nil
}

func (service *RolesServiceImpl) Delete(ctx *gin.Context, request int) error {

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	role, err := service.Repository.FindById(ctx, tx, request)

	if err != nil {
		return errors.New(err.Error())
	}

	errDelete := service.Repository.Delete(ctx, tx, role.ID)

	if errDelete != nil {
		return errors.New(errDelete.Error())
	}

	return nil
}

func (service *RolesServiceImpl) FindAll(ctx *gin.Context, request domain.RolesFindAllRequest) ([]domain.RolesResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	roles, err := service.Repository.FindAll(ctx, tx, request.Page, request.Limit, request.Sort)

	if err != nil {
		return domain.ToRolesResponses(roles), err
	}

	return domain.ToRolesResponses(roles), nil
}

func (service *RolesServiceImpl) FindById(ctx *gin.Context, request int) (domain.RolesResponse, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	roles, err := service.Repository.FindById(ctx, tx, request)

	if err != nil {
		return domain.ToRolesResponse(roles), err
	}

	return domain.ToRolesResponse(roles), nil
}
