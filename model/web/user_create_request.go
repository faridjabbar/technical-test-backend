package web

type UserCreateRequest struct {
	// Fields
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required,role"`
	Password string `json:"password" validate:"required"`
}
