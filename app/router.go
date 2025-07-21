package app

import (
	"net/http"
	handler "technical-test-backend/handler"
	"technical-test-backend/helper"
	"technical-test-backend/route"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB, validate *validator.Validate) *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.Use(handler.ErrorHandler(helper.DatabaseError))
	route.UserRoute(router, db, validate)
	route.TraditionalFoodRoute(router, db, validate)
	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// A CORS middleware.
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
