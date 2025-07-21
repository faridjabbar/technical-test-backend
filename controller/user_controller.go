package controller

import (
	"github.com/gin-gonic/gin"
)

type UserController interface {
	Create(context *gin.Context)
	Login(context *gin.Context)
}
