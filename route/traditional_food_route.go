package route

import (
	"technical-test-backend/controller"
	"technical-test-backend/repository"
	"technical-test-backend/service"

	auth "technical-test-backend/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func TraditionalFoodRoute(router *gin.Engine, db *gorm.DB, validate *validator.Validate) {

	traditionalFoodService := service.NewTraditionalFoodService(
		repository.NewTraditionalFoodRepository(),
		db,
		validate,
	)
	traditionalFoodController := controller.NewTraditionalFoodController(traditionalFoodService)

	router.DELETE("/traditional-food/:id", auth.Auth(traditionalFoodController.Delete, []string{auth.RoleAdministrator}))
	router.GET("/traditional-food", auth.Auth(traditionalFoodController.FindAll, []string{auth.RoleAdministrator}))
	router.POST("/traditional-food", auth.Auth(traditionalFoodController.Create, []string{auth.RoleAdministrator}))
	router.PUT("/traditional-food/:id", auth.Auth(traditionalFoodController.Update, []string{auth.RoleAdministrator}))
}
