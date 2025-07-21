package service

import (
	"technical-test-backend/model/web"

	"github.com/gin-gonic/gin"
)

type UserService interface {
	Create(request *web.UserCreateRequest, c *gin.Context) web.UserResponse
	Login(request *web.UserLoginRequest, c *gin.Context)
}
