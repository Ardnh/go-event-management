package app

import (
	"fmt"
	"go/ems/helper"
	"go/ems/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {

	// db_username := helper.LoadEnvFile("DB_USERNAME")
	// db_password := helper.LoadEnvFile("DB_PASSWORD")
	// db_name := helper.LoadEnvFile("DB_NAME")

	db_username := "root"
	db_password := "root"
	db_name := "go_ems"

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3307)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	helper.PanicIfError(err)

	db.AutoMigrate(&models.Roles{})
	db.AutoMigrate(&models.Advertisements{})
	db.AutoMigrate(&models.Categories{})
	db.AutoMigrate(&models.Events{})
	db.AutoMigrate(&models.Users{})

	return db
}
