package controller

import (
	"net/http"
	"technical-test-backend/helper"
	"technical-test-backend/model/web"
	webResponse "technical-test-backend/model/web"
	"technical-test-backend/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}
func (controller *UserControllerImpl) Create(c *gin.Context) {
	request := web.UserCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	userResponse := controller.UserService.Create(&request, c)
	webResponse := webResponse.WebResponse{
		Success: true,
		Message: "User created successfully",
		Data:    userResponse,
	}

	c.JSON(http.StatusCreated, webResponse)
}
