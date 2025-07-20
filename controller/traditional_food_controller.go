package controller

import (
	auth "technical-test-backend/auth"

	"github.com/gin-gonic/gin"
)

type TraditionalFoodController interface {
	Create(context *gin.Context, auth *auth.AccessDetails)
	FindAll(context *gin.Context, auth *auth.AccessDetails)
	Delete(context *gin.Context, auth *auth.AccessDetails)
	Update(context *gin.Context, auth *auth.AccessDetails)
}
