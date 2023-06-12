package repository

import (
	"fmt"
	domain "go/ems/domain/events"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EventsRepositoryImpl struct {
}

func NewEventsRepository() EventsRepository {
	return &EventsRepositoryImpl{}
}

func (repository *EventsRepositoryImpl) Create(ctx *gin.Context, tx *gorm.DB, data *domain.Events) (*domain.Events, error) {

	err := tx.
		WithContext(ctx).
		Table("events").
		Create(&data).
		Error

	if err != nil {
		return data, err
	} else {
		return data, nil
	}
}

func (repository *EventsRepositoryImpl) Update(ctx *gin.Context, tx *gorm.DB, data *domain.Events) (*domain.Events, error) {

	err := tx.
		WithContext(ctx).
		Table("events").
		Where("id = ?", &data.Id).
		Updates(&data).
		Error

	if err != nil {
		return data, err
	} else {
		return data, nil
	}

}

func (repository *EventsRepositoryImpl) UpdateEventsViews(ctx *gin.Context, tx *gorm.DB, id int) error {
	var view int
	err := tx.
		WithContext(ctx).
		Table("events").
		Select("views").
		Where("events.id = ?", id).
		First(&view).
		Error

	if err != nil {
		return err
	}

	errUpdate := tx.
		WithContext(ctx).
		Table("events").
		Select("views").
		Update("views", view+1).
		Where("events.id = ?", id).
		Error

	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (repository *EventsRepositoryImpl) Delete(ctx *gin.Context, tx *gorm.DB, id int) error {

	if err := tx.WithContext(ctx).Delete(&domain.Events{}, id).Error; err != nil {
		return err
	} else {
		return nil
	}

}

func (repository *EventsRepositoryImpl) FindByEventsId(ctx *gin.Context, tx *gorm.DB, id int) (*domain.Events, error) {

	var result *domain.Events

	err := tx.
		WithContext(ctx).
		Table("events").
		Where("id = ?", &id).
		First(&result).
		Error

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func (repository *EventsRepositoryImpl) FindAll(ctx *gin.Context, tx *gorm.DB, pageNumber int, pageSize int, orderBy string) (*[]domain.Events, error) {
	var result *[]domain.Events
	sort := fmt.Sprintf("id %s", strings.ToUpper(orderBy))

	err := tx.
		WithContext(ctx).
		Table("events").
		Select("events.id, ").
		Limit(pageSize).
		Offset((pageNumber - 1) * pageSize).
		Order(sort).
		Find(&result).
		Error

	if err != nil {
		return result, err
	} else {
		return result, nil
	}

}

func (repository *EventsRepositoryImpl) FindByUserId(ctx *gin.Context, tx *gorm.DB, id int, pageNumber int, pageSize int, orderBy string) (*[]domain.Events, error) {
	var result *[]domain.Events
	sort := fmt.Sprintf("id %s", strings.ToUpper(orderBy))

	err := tx.
		WithContext(ctx).
		Table("events").
		Select("events.id, events.name, events.description, events.start_date, events.end_date, events.registration_url, events.banner, events.address, events.views, events.user_id, events.category_id, users.username").
		Joins("JOIN users ON events.user_id = users.id").
		Limit(pageSize).
		Offset((pageNumber - 1) * pageSize).
		Order(sort).
		Find(&result).
		Error

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func (repository *EventsRepositoryImpl) FindByCategoryId(ctx *gin.Context, tx *gorm.DB, id int, pageNumber int, pageSize int, orderBy string) (*[]domain.Events, error) {
	var result *[]domain.Events
	sort := fmt.Sprintf("id %s", strings.ToUpper(orderBy))

	err := tx.
		WithContext(ctx).
		Table("events").
		Select("events.id, events.name, events.description, events.start_date, events.end_date, events.registration_url, events.banner, events.address, events.views, events.user_id, events.category_id").
		Where("category_id = ?", id).
		Limit(pageSize).
		Offset((pageNumber - 1) * pageSize).
		Order(sort).
		Find(&result).
		Error

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func (repository *EventsRepositoryImpl) FindTodayEvents(ctx *gin.Context, tx *gorm.DB, date string, pageNumber int, pageSize int, orderBy string) (*[]domain.Events, error) {
	var result *[]domain.Events
	sort := fmt.Sprintf("id %s", strings.ToUpper(orderBy))

	err := tx.
		WithContext(ctx).
		Table("events").
		Select("events.id, events.name, events.description, events.start_date, events.end_date, events.registration_url, events.banner, events.address, events.views, events.user_id, events.category_id").
		Where("start_date = ?", date).
		Limit(pageSize).
		Offset((pageNumber - 1) * pageSize).
		Order(sort).
		Find(&result).
		Error

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}

func (repository *EventsRepositoryImpl) FindWeeklyEvents(ctx *gin.Context, tx *gorm.DB, date string) (*[]domain.Events, error) {

}

func (repository *EventsRepositoryImpl) FindMonthlyEvents(ctx *gin.Context, tx *gorm.DB, date string) (*[]domain.Events, error) {

}

func (repository *EventsRepositoryImpl) FindUpcomingEvents(ctx *gin.Context, tx *gorm.DB, date string) (*[]domain.Events, error) {

}

func (repository *EventsRepositoryImpl) FindMostViewedEvents(ctx *gin.Context, tx *gorm.DB, pageSize int, orderBy string) (*[]domain.Events, error) {
	var result *[]domain.Events
	sort := fmt.Sprintf("views %s", strings.ToUpper(orderBy))

	err := tx.
		WithContext(ctx).
		Table("events").
		Select("events.id, events.name, events.description, events.start_date, events.end_date, events.registration_url, events.banner, events.address, events.views, events.user_id, events.category_id").
		Limit(pageSize).
		Order(sort).
		Find(&result).
		Error

	if err != nil {
		return result, err
	} else {
		return result, nil
	}
}
