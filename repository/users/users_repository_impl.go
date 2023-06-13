package repository

import (
	domain "go/ems/domain/users"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Register(ctx *gin.Context, tx *gorm.DB, req *domain.UserRegister) error {

	if err := tx.WithContext(ctx).Table("users").Create(&req).Error; err != nil {
		return err
	}
	return nil

}

func (repository *UserRepositoryImpl) FindByEmail(ctx *gin.Context, tx *gorm.DB, email string) (*domain.UserQueryResponse, error) {
	var result domain.UserQueryResponse

	// SELECT first_name, last_name, username, email, active, role_id, name as role FROM users JOIN roles ON users.role_id = roles.id WHERE users.id = 14;
	err := tx.
		WithContext(ctx).
		Table("users").
		Select("users.id, users.first_name, users.last_name, users.username, users.email, users.password, users.active, users.role_id, name as role").
		Joins("JOIN roles ON users.role_id = roles.id").
		Where("users.email = ?", strings.ToLower(email)).
		First(&result).
		Error

	if err != nil {
		return &result, err
	} else {
		return &result, nil
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

func (repository *UserRepositoryImpl) UpdateActiveStatusTrue(ctx *gin.Context, tx *gorm.DB, id *int) error {
	var active bool

	err := tx.
		WithContext(ctx).
		Table("users").
		Select("active").
		Where("events.id = ?", &id).
		First(&active).
		Error

	if err != nil {
		return err
	} else {
		active = true
	}

	errUpdate := tx.
		WithContext(ctx).
		Table("events").
		Update("views = ?", &active).
		Where("events.id = ?", &id).
		Error

	if errUpdate != nil {
		return errUpdate
	} else {
		return nil
	}
}

func (repository *UserRepositoryImpl) UpdateActiveStatusFalse(ctx *gin.Context, tx *gorm.DB, id int) error {
	var active bool

	err := tx.
		WithContext(ctx).
		Table("users").
		Select("active").
		Where("events.id = ?", &id).
		First(&active).
		Error

	if err != nil {
		return err
	} else {
		active = false
	}

	errUpdate := tx.
		WithContext(ctx).
		Table("events").
		Update("views = ?", &active).
		Where("events.id = ?", &id).
		Error

	if errUpdate != nil {
		return errUpdate
	} else {
		return nil
	}
}
