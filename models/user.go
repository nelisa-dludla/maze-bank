package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `validate:"required"`
	LastName string `validate:"required"`
	Password string `validate:"required,min=8"`
	Email string `gorm:"unique" validate:"required,email"`
	Balance float64 `gorm:"default:0.00"`
	Transactions []Transaction
}
