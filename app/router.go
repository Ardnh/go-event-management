package app

import (
	rolesController "go/ems/controller/roles"
	"go/ems/middleware"
	rolesRepository "go/ems/repository/roles"
	rolesService "go/ems/service/roles"

	userController "go/ems/controller/users"
	userRepository "go/ems/repository/users"
	userService "go/ems/service/users"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.Default()

	rolesRepository := rolesRepository.NewRolesRepository()
	rolesService := rolesService.NewRolesService(rolesRepository, db, validate)
	rolesController := rolesController.NewRolesController(rolesService)

	userRepostory := userRepository.NewUserRepository()
	userService := userService.NewUserService(userRepostory, db, validate)
	userController := userController.NewUserController(userService)

	user := router.Group("/api/v1/user")
	admin := router.Group("/api/v1/admin")

	// ADMIN
	{
		admin.Use(middleware.AuthCheck).POST("/roles", rolesController.Create)
		admin.Use(middleware.AuthCheck).PUT("/roles", rolesController.Update)
		admin.Use(middleware.AuthCheck).DELETE("/roles", rolesController.Delete)
		admin.Use(middleware.AuthCheck).GET("/roles", rolesController.FindAll)
		admin.Use(middleware.AuthCheck).GET("/role", rolesController.FindById)
	}

	// USER
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)

	}

	return router
}
