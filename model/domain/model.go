package domain

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        string `gorm:"primarykey; size:20"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
