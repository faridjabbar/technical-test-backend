package service

import (
	"technical-test-backend/model/web"

	auth "technical-test-backend/auth"

	"github.com/gin-gonic/gin"
)

type TraditionalFoodService interface {
	Create(auth *auth.AccessDetails, request *web.TraditionalFoodCreateRequest, c *gin.Context) web.TraditionalFoodResponse
	Delete(auth *auth.AccessDetails, id *int, c *gin.Context)
	Update(auth *auth.AccessDetails, id *int, request *web.TraditionalFoodUpdateRequest, c *gin.Context) web.TraditionalFoodResponse
	FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context) []web.TraditionalFoodResponse
}
