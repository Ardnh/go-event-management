package repository

import (
	"fmt"
	domain "go/ems/domain/category"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewRolesRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindByName(ctx *gin.Context, tx *gorm.DB, name string) bool {

	var category *domain.Category
	if err := tx.WithContext(ctx).Where("name = ?", strings.ToLower(name)).First(&category).Error; err != nil {
		return true
	} else {
		return false
	}

}

func (repository *CategoryRepositoryImpl) FindById(ctx *gin.Context, tx *gorm.DB, id int) (*domain.Category, error) {

	var category *domain.Category
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB, pageNumber int, pageSize int, orderBy string) ([]*domain.Category, error) {

	var category []*domain.Category
	sort := fmt.Sprintf("id %s", strings.ToUpper(orderBy))
	if err := tx.WithContext(ctx).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Order(sort).Find(&category).Error; err != nil {
		return category, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, data *domain.Category) (*domain.Category, error) {

	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (repository *CategoryRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, data *domain.Category) (*domain.Category, error) {

	if err := tx.WithContext(ctx).Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (repository *CategoryRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, id int) error {

	if err := tx.WithContext(ctx).Delete(&domain.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
