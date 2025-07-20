package domain

import (
	"gorm.io/gorm"
)

type Histories []History

type History struct {
	gorm.Model
	CreatedByID uint `gorm:""`

	// Fields
	TableName string `gorm:"Index:idx_history;size:100"`
	TableID   string `gorm:"Index:idx_history;size:50"`
	Data      string `gorm:"type:text"`
	Type      string `gorm:"Index:idx_history;size:10"`
}
