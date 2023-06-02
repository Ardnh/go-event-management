package repository

import (
	"fmt"
	"go/ems/domain"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RolesRepositoryImpl struct {
}

func NewRolesRepository() RolesRepository {
	return &RolesRepositoryImpl{}
}

func (repository *RolesRepositoryImpl) FindByName(ctx *gin.Context, tx *gorm.DB, name string) bool {

	var role domain.Roles
	if err := tx.WithContext(ctx).Where("name = ?", strings.ToLower(name)).First(&role).Error; err != nil {
		return true
	} else {
		return false
	}

}

func (repository *RolesRepositoryImpl) FindById(ctx *gin.Context, tx *gorm.DB, id int) (domain.Roles, error) {

	var role domain.Roles
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&role).Error; err != nil {
		return role, err
	}

	return role, nil
}

func (repository *RolesRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB, pageNumber int, pageSize int, orderBy string) ([]domain.Roles, error) {

	var roles []domain.Roles
	sort := fmt.Sprintf("id %s", strings.ToUpper(orderBy))
	if err := tx.WithContext(ctx).Limit(pageSize).Offset((pageNumber - 1) * pageSize).Order(sort).Find(&roles).Error; err != nil {
		return roles, err
	}

	return roles, nil
}

func (repository *RolesRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, data domain.Roles) (domain.Roles, error) {

	if err := tx.WithContext(ctx).Create(&data).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (repository *RolesRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, data domain.Roles) (domain.Roles, error) {

	if err := tx.WithContext(ctx).Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (repository *RolesRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, id int) error {

	if err := tx.WithContext(ctx).Delete(&domain.Roles{}, id).Error; err != nil {
		return err
	}

	return nil
}
