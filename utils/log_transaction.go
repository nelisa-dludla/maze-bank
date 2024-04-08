package utils

import (
	"fmt"
	"maze/database"
	"maze/models"
	"time"

	"github.com/google/uuid"
)

func LogTransaction(userID uint, description string, amount float64) error {
	// Create transaction
	transaction := models.Transaction{
		ID: uuid.NewString(),
		UserID: userID,
		CreatedAt: time.Now(),
		Description: "Deposit",
		Amount: amount,
	}
	// Add transaction to database
	transactionResult := database.DB.Create(&transaction)
	if transactionResult.Error != nil {
		return fmt.Errorf("Error logging transaction: %s", transactionResult.Error)
	}
	return nil
}
