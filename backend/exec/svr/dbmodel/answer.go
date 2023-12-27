package dbmodel

import (
	"gorm.io/gorm"
	"time"
)

type Answer struct {
	ID          string `gorm:"primaryKey"`
	Question    string `gorm:"type:text"`
	FirstAnswer string `gorm:"type:longtext"`
	Evidence    string `gorm:"type:longtext"`
	IsOk        bool   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
