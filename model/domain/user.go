package domain

import (
	"technical-test-backend/model/web"

	"gorm.io/gorm"
)

type Users []User
type User struct {
	// Required Fields
	gorm.Model
	CreatedByID uint  `gorm:""`
	UpdatedByID uint  `gorm:""`
	DeletedByID *uint `gorm:""`

	// Fields
	Name     string `gorm:"size:50;uniqueIndex:idx_user"`
	Email    string `gorm:"size:50;uniqueIndex:idx_user"`
	Role     string `gorm:"size:15"`
	Password string `gorm:"size:255"`
}

func (user *User) ToUserResponse() web.UserResponse {
	return web.UserResponse{
		// Required Fields
		CreatedAt:   user.CreatedAt,
		CreatedByID: user.CreatedByID,
		UpdatedAt:   user.UpdatedAt,
		UpdatedByID: user.UpdatedByID,

		// Fields
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

func (users Users) ToUserResponses() []web.UserResponse {
	userResponses := []web.UserResponse{}
	for _, user := range users {
		userResponses = append(userResponses, user.ToUserResponse())
	}
	return userResponses
}
