package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name          string
	Price         float64
	ProductNumber string `gorm:"unique"`
	Stock         int
	TaxID         int
}
