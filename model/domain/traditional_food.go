package domain

import (
	"technical-test-backend/model/web"

	"gorm.io/gorm"
)

type TraditionalFoods []TraditionalFood
type TraditionalFood struct {
	// Required Fields
	gorm.Model
	CreatedByID uint  `gorm:""`
	UpdatedByID uint  `gorm:""`
	DeletedByID *uint `gorm:""`

	// Fields
	Name           string `gorm:"size:50;uniqueIndex:idx_traditiona_food"`
	RegionalOrigin string `gorm:"size:50;uniqueIndex:idx_traditiona_food"`
	MainIngredient string `gorm:""`
	Type           string `gorm:"size:20"`
	Description    string `gorm:""`
}

func (traditionalFood *TraditionalFood) ToTraditionalFoodResponse() web.TraditionalFoodResponse {
	return web.TraditionalFoodResponse{
		// Required Fields
		ID:          traditionalFood.ID,
		CreatedAt:   traditionalFood.CreatedAt,
		CreatedByID: traditionalFood.CreatedByID,
		UpdatedAt:   traditionalFood.UpdatedAt,
		UpdatedByID: traditionalFood.UpdatedByID,

		// Fields
		Name:           traditionalFood.Name,
		RegionalOrigin: traditionalFood.RegionalOrigin,
		MainIngredient: traditionalFood.MainIngredient,
		Type:           traditionalFood.Type,
		Description:    traditionalFood.Description,
	}
}

func (traditionalFoods TraditionalFoods) ToTraditionalFoodResponses() []web.TraditionalFoodResponse {
	traditionalFoodResponses := []web.TraditionalFoodResponse{}
	for _, traditionalFood := range traditionalFoods {
		traditionalFoodResponses = append(traditionalFoodResponses, traditionalFood.ToTraditionalFoodResponse())
	}
	return traditionalFoodResponses
}
