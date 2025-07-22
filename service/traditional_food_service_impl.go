package service

import (
	auth "technical-test-backend/auth"
	"technical-test-backend/repository"

	helper "technical-test-backend/helper"
	"technical-test-backend/model/domain"
	"technical-test-backend/model/web"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type TraditionalFoodServiceImpl struct {
	TraditionalFoodRepository repository.TraditionalFoodRepository
	DB                        *gorm.DB
	Validate                  *validator.Validate
}

func NewTraditionalFoodService(
	traditionalFood repository.TraditionalFoodRepository,
	db *gorm.DB,
	validate *validator.Validate,
) TraditionalFoodService {
	return &TraditionalFoodServiceImpl{
		TraditionalFoodRepository: traditionalFood,
		DB:                        db,
		Validate:                  validate,
	}
}

func (service *TraditionalFoodServiceImpl) FindAll(auth *auth.AccessDetails, filters *map[string]string, c *gin.Context, search string) []web.TraditionalFoodResponse {
	db := service.DB.Begin()
	defer helper.CommitOrRollback(db)

	traditionalFoods := service.TraditionalFoodRepository.FindAll(db, filters, search)
	return traditionalFoods.ToTraditionalFoodResponses()
}

func (service *TraditionalFoodServiceImpl) Create(auth *auth.AccessDetails, request *web.TraditionalFoodCreateRequest, c *gin.Context) web.TraditionalFoodResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	db := service.DB.Begin()
	defer helper.CommitOrRollback(db)

	traditionalFood := &domain.TraditionalFood{
		// Required Fields
		CreatedByID: auth.ID,
		UpdatedByID: auth.ID,

		// Fields
		Name:           request.Name,
		RegionalOrigin: request.RegionalOrigin,
		MainIngredient: request.MainIngredient,
		Type:           request.Type,
		Description:    request.Description,
	}
	traditionalFood = service.TraditionalFoodRepository.Create(db, traditionalFood)

	return traditionalFood.ToTraditionalFoodResponse()
}

func (service *TraditionalFoodServiceImpl) Delete(auth *auth.AccessDetails, id *int, c *gin.Context) {
	db := service.DB.Begin()
	defer helper.CommitOrRollback(db)

	service.TraditionalFoodRepository.Delete(db, id, &auth.ID)
}

func (service *TraditionalFoodServiceImpl) Update(auth *auth.AccessDetails, id *int, request *web.TraditionalFoodUpdateRequest, c *gin.Context) web.TraditionalFoodResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	db := service.DB.Begin()
	defer helper.CommitOrRollback(db)

	helper.PanicIfError(err)
	traditionalFood := &domain.TraditionalFood{
		// Required Fields
		CreatedByID: auth.ID,
		UpdatedByID: auth.ID,

		// Fields
		Name:           request.Name,
		RegionalOrigin: request.RegionalOrigin,
		MainIngredient: request.MainIngredient,
		Type:           request.Type,
		Description:    request.Description,
	}
	traditionalFood = service.TraditionalFoodRepository.Update(db, traditionalFood, id)

	return traditionalFood.ToTraditionalFoodResponse()
}
