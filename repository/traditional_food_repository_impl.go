package repository

import (
	helper "technical-test-backend/helper"
	"technical-test-backend/model/domain"

	"gorm.io/gorm"
)

type TraditionalFoodRepositoryImpl struct {
}

func NewTraditionalFoodRepository() TraditionalFoodRepository {
	return &TraditionalFoodRepositoryImpl{}
}

var searchColumns = []string{"name", "regional_origin", "main_ingredient", "type", "description"}

func (repository *TraditionalFoodRepositoryImpl) FindAll(db *gorm.DB, filters *map[string]string, search string) domain.TraditionalFoods {
	traditionalFood := domain.TraditionalFoods{}
	tx := db.Model(&domain.TraditionalFood{})

	err := helper.ApplyFilter(tx, filters)
	helper.PanicIfError(err)

	tx.Where(helper.SearchCondition(search, searchColumns))

	err = tx.Find(&traditionalFood).Error
	helper.PanicIfError(err)

	return traditionalFood
}

func (repository *TraditionalFoodRepositoryImpl) Create(db *gorm.DB, traditionalFood *domain.TraditionalFood) *domain.TraditionalFood {
	err := db.Create(&traditionalFood).Error
	helper.PanicIfError(err)

	err = helper.CreateHistory(db, traditionalFood, helper.HistoryCreate, traditionalFood.CreatedByID)
	helper.PanicIfError(err)
	return traditionalFood
}

func (repository *TraditionalFoodRepositoryImpl) Delete(db *gorm.DB, id *int, deletedByID *uint) {
	transactionType := &domain.TraditionalFood{}
	tx := db.First(transactionType, id).Updates(&domain.TraditionalFood{
		Model:       gorm.Model{ID: uint(*id)},
		DeletedByID: deletedByID,
	})

	// Creating a helper of the deleted TraditionalFood.
	err := helper.CreateHistory(db, transactionType, helper.HistoryDelete, *deletedByID)
	helper.PanicIfError(err)

	// Deleting the TraditionalFood from the database.
	err = tx.Unscoped().Delete(transactionType, id).Error
	helper.PanicIfError(err)
}

func (repository *TraditionalFoodRepositoryImpl) Update(db *gorm.DB, traditionalFood *domain.TraditionalFood, id *int) *domain.TraditionalFood {
	err := db.Model(&domain.TraditionalFood{}).
		Where(&domain.TraditionalFood{
			Model: gorm.Model{ID: uint(*id)},
		}).
		Updates(&traditionalFood).First(traditionalFood).Error
	helper.PanicIfError(err)

	err = helper.CreateHistory(db, traditionalFood, helper.HistoryUpdate, traditionalFood.UpdatedByID)
	helper.PanicIfError(err)

	return traditionalFood
}
