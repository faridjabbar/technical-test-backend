package web

type BalanceGroupCreateRequest struct {
	// Fields
	ID       string `json:"id" validate:"required,min=1,max=20"`
	Name     string `json:"name"`
	Type     string `json:"type" validate:"type_debt_and_credit"`
	Category string `json:"category" validate:"category"`
}
