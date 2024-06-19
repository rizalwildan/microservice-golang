package models

import "gorm.io/gorm"

type Merchant struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
	City        string `gorm:"not null"`
	UserID      uint   `gorm:"not null"`
}
