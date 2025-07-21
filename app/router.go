package app

import (
	handler "technical-test-backend/handler"
	"technical-test-backend/helper"
	"technical-test-backend/route"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.New()
	router.Use(handler.ErrorHandler(helper.DatabaseError))
	route.UserRoute(router, db, validate)
	route.TraditionalFoodRoute(router, db, validate)
	return router
}
