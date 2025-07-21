package web

type UserLoginRequest struct {
	// Fields
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
