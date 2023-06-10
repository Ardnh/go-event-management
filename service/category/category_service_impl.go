package service

import (
	"errors"
	domain "go/ems/domain/category"
	"go/ems/helper"
	repository "go/ems/repository/category"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	Repository repository.CategoryRepository
	DB         *gorm.DB
	Validate   *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, db *gorm.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		Repository: repository,
		DB:         db,
		Validate:   validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx *gin.Context, request *domain.CategoryCreateRequest) (*domain.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	isAvailable := service.Repository.FindByName(ctx, tx, request.Name)

	if isAvailable {
		result, _ := service.Repository.Create(ctx, tx, &category)
		return domain.ToCategoryResponse(result), nil
	} else {
		return domain.ToCategoryResponse(&category), errors.New("roles already exist")
	}
}

func (service *CategoryServiceImpl) Update(ctx *gin.Context, request *domain.CategoryUpdateRequest) (*domain.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	role, err := service.Repository.FindById(ctx, tx, request.Id)

	if err != nil {
		return domain.ToCategoryResponse(role), errors.New(err.Error())
	}

	role.Name = request.Name
	role.UpdatedAt = time.Now()

	updatedRole, err := service.Repository.Update(ctx, tx, role)

	if err != nil {
		return domain.ToCategoryResponse(updatedRole), errors.New(err.Error())
	}

	return domain.ToCategoryResponse(updatedRole), nil
}

func (service *CategoryServiceImpl) Delete(ctx *gin.Context, request int) error {

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

func (service *CategoryServiceImpl) FindAll(ctx *gin.Context, request *domain.CategoryFindAllRequest) ([]*domain.CategoryResponse, error) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	roles, err := service.Repository.FindAll(ctx, tx, request.Page, request.Limit, request.Sort)

	if err != nil {
		return domain.ToCategoryResponses(roles), err
	}

	return domain.ToCategoryResponses(roles), nil
}

func (service *CategoryServiceImpl) FindById(ctx *gin.Context, request int) (*domain.CategoryResponse, error) {
	tx := service.DB.Begin()
	defer helper.CommitOrRollback(tx)

	roles, err := service.Repository.FindById(ctx, tx, request)

	if err != nil {
		return domain.ToCategoryResponse(roles), err
	}

	return domain.ToCategoryResponse(roles), nil
}
