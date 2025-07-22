package repository

import (
	"technical-test-backend/model/domain"

	"gorm.io/gorm"
)

type TraditionalFoodRepository interface {
	Create(db *gorm.DB, traditionalFood *domain.TraditionalFood) *domain.TraditionalFood
	Delete(db *gorm.DB, id *int, deletedByID *uint)
	FindAll(db *gorm.DB, filters *map[string]string, search string) domain.TraditionalFoods
	Update(db *gorm.DB, traditionalFood *domain.TraditionalFood, id *int) *domain.TraditionalFood
}
