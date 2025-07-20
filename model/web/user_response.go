package web

import (
	"time"
)

type UserResponse struct {
	// Required Fields
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedByID uint      `json:"created_by_id"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedByID uint      `json:"updated_by_id"`

	// Fields
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Password string `json:"password"`
}
