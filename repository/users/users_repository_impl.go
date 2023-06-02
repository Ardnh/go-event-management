package repository

import (
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
