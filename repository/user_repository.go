package repository

import (
	"technical-test-backend/model/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(db *gorm.DB, user *domain.User) *domain.User
}
