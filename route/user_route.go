package route

import (
	"technical-test-backend/controller"
	"technical-test-backend/repository"
	"technical-test-backend/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UserRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	userService := service.NewUserService(
		repository.NewUserRepository(),
		db,
		validate,
	)
	userController := controller.NewUserController(userService)

	router.POST("/users", userController.Create)
	router.POST("/users/login", userController.Login)
}
