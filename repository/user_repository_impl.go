package repository

import (
	"technical-test-backend/helper"
	"technical-test-backend/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Create(db *gorm.DB, user *domain.User) *domain.User {
	err := db.Create(&user).Error
	helper.PanicIfError(err)

	err = helper.CreateHistory(db, user, helper.HistoryCreate, user.CreatedByID)
	helper.PanicIfError(err)
	return user
}

func (repository *UserRepositoryImpl) FindByEmail(db *gorm.DB, email *string) domain.User {
	var user domain.User
	err := db.Where(&domain.User{
		Email: *email,
	}).First(&user).Error
	helper.PanicIfError(err)
	return user
}
