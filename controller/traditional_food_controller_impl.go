package controller

import (
	"net/http"
	"strconv"

	auth "technical-test-backend/auth"
	"technical-test-backend/helper"
	webResponse "technical-test-backend/model/web"
	"technical-test-backend/service"

	"github.com/gin-gonic/gin"
)

type TraditionalFoodControllerImpl struct {
	TraditionalFoodService service.TraditionalFoodService
}

func NewTraditionalFoodController(traditionalFoodService service.TraditionalFoodService) TraditionalFoodController {
	return &TraditionalFoodControllerImpl{
		TraditionalFoodService: traditionalFoodService,
	}
}

func (controller *TraditionalFoodControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.like")

	search := c.Query("search")

	traditionalFoodResponses := controller.TraditionalFoodService.FindAll(auth, &filters, c, search)
	webResponse := webResponse.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(traditionalFoodResponses),
		Data:    traditionalFoodResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TraditionalFoodControllerImpl) Create(c *gin.Context, auth *auth.AccessDetails) {
	request := webResponse.TraditionalFoodCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	traditionalFoodResponse := controller.TraditionalFoodService.Create(auth, &request, c)
	webResponse := webResponse.WebResponse{
		Success: true,
		Message: "Traditional Food created successfully",
		Data:    traditionalFoodResponse,
	}

	c.JSON(http.StatusCreated, webResponse)
}

func (controller *TraditionalFoodControllerImpl) Update(c *gin.Context, auth *auth.AccessDetails) {
	id := c.Param("id")

	traditionalFoodID, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	request := webResponse.TraditionalFoodUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	traditionalFoodResponse := controller.TraditionalFoodService.Update(auth, &traditionalFoodID, &request, c)
	webResponse := webResponse.WebResponse{
		Success: true,
		Message: "Traditional Food updated successfully",
		Data:    traditionalFoodResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *TraditionalFoodControllerImpl) Delete(c *gin.Context, auth *auth.AccessDetails) {
	id := c.Param("id")

	traditionalFoodID, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	controller.TraditionalFoodService.Delete(auth, &traditionalFoodID, c)
	webResponse := webResponse.WebResponse{
		Success: true,
		Message: "Traditional Food deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}
