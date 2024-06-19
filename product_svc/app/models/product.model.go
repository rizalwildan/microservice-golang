package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"type:varchar(200);not null"`
	Description string
	Quantity    float32
	Price       float32 `gorm:"type:decimal(10,2)"`
	MerchantID  uint    `gorm:"not null"`
}
