package web

import (
	"time"
)

type TraditionalFoodResponse struct {
	// Required Fields
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedByID uint      `json:"created_by_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedByID uint      `json:"updated_by_id"`

	// Fields
	Name           string `json:"name"`
	RegionalOrigin string `json:"regional_origin"`
	MainIngredient string `json:"main_ingredient"`
	Type           string `json:"type"`
	Description    string `json:"description"`
}
