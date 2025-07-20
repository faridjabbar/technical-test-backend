package web

type TraditionalFoodUpdateRequest struct {
	// Fields
	Name           string `json:"name" validate:"required"`
	RegionalOrigin string `json:"regional_origin" validate:"required"`
	MainIngredient string `json:"main_ingredient" validate:"required"`
	Type           string `json:"type" validate:"required,type_food"`
	Description    string `json:"description"`
}
