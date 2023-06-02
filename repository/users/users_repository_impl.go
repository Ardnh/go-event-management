package repository

import (
	"fmt"
	"go/ems/domain"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx *gin.Context, tx *gorm.DB, req *domain.User) error {

	if err := tx.WithContext(ctx).Table("users").Create(&req).Error; err != nil {
		return err
	}
	return nil

}

func (repository *UserRepositoryImpl) FindByUsername(ctx *gin.Context, tx *gorm.DB, username string) (*domain.User, error) {
	var user domain.User

	if err := tx.WithContext(ctx).Table("users").Where("username = ?", strings.ToLower(username)).First(&user).Error; err != nil {
		return &user, err
	} else {
		return &user, nil
	}
}

func (repository *UserRepositoryImpl) FindById(ctx *gin.Context, tx *gorm.DB, id int) (*domain.User, error) {
	var user domain.User

	if err := tx.WithContext(ctx).Table("users").Where("id = ?", id).First(&user).Error; err != nil {
		return &user, err
	} else {
		return &user, nil
	}
}

func (repository *UserRepositoryImpl) UpdateUserById(ctx *gin.Context, tx *gorm.DB, req *domain.UserUpdateRequest) (*domain.User, error) {
	var user domain.User

	updateData := map[string]interface{}{
		"Id":        req.Id,
		"FirstName": req.FirstName,
		"LastName":  req.LastName,
		"Username":  req.UserName,
		"Email":     req.Email,
	}

	if err := tx.WithContext(ctx).Model(&user).Updates(&updateData).Error; err != nil {
		fmt.Println("Update success")
		return &user, err
	} else {
		fmt.Println("Update Fail")
		return &user, nil
	}
}
