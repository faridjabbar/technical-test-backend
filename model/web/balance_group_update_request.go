package web

type BalanceGroupUpdateRequest struct {
	// Fields
	Name     string `json:"name"`
	Type     string `json:"type"`
	Category string `json:"category"`
}
