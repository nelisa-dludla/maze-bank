package models

import "time"

type Transaction struct {
	ID string
	UserID uint
	CreatedAt time.Time
	Description string
	Amount float64
}
