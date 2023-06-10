package app

import (
	rolesController "go/ems/controller/roles"
	"go/ems/middleware"
	rolesRepository "go/ems/repository/roles"
	rolesService "go/ems/service/roles"

	categoryController "go/ems/controller/category"
	categoryRepository "go/ems/repository/category"
	categoryService "go/ems/service/category"

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

	categoryRepository := categoryRepository.NewRolesRepository()
	categoryService := categoryService.NewCategoryService(categoryRepository, db, validate)
	categoryController := categoryController.NewCategoryController(categoryService)

	userRepostory := userRepository.NewUserRepository()
	userService := userService.NewUserService(userRepostory, db, validate)
	userController := userController.NewUserController(userService)

	user := router.Group("/api/v1/user")
	admin := router.Group("/api/v1/admin")
	public := router.Group("/api/v1")

	// ADMIN
	{

		admin.Use(middleware.AuthAdminCheck).POST("/roles", rolesController.Create)
		admin.Use(middleware.AuthAdminCheck).PUT("/roles", rolesController.Update)
		admin.Use(middleware.AuthAdminCheck).DELETE("/roles", rolesController.Delete)
		admin.Use(middleware.AuthAdminCheck).GET("/roles", rolesController.FindAll)
		admin.Use(middleware.AuthAdminCheck).GET("/role", rolesController.FindById)

		admin.Use(middleware.AuthAdminCheck).POST("/category", categoryController.Create)
		admin.Use(middleware.AuthAdminCheck).PUT("/category", categoryController.Update)
		admin.Use(middleware.AuthAdminCheck).DELETE("/category", categoryController.Delete)
		admin.Use(middleware.AuthAdminCheck).GET("/category", categoryController.FindById)
	}

	// PUBLIC
	{
		public.GET("/category", categoryController.FindAll)

	}

	// USER
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)

	}

	return router
}
