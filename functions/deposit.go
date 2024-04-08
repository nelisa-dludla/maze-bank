package functions

import (
	"fmt"
	"maze/database"
	"maze/models"
	"maze/utils"

	"github.com/fatih/color"
)

func Deposit(user *models.User) {
	var depositAmount float64

	color.Cyan("===== Maze Bank - Deposit =====")
	fmt.Println()
	fmt.Print("Enter amount: ")

	_, err := fmt.Scanf("%f", &depositAmount)
	// Check if a float was entered and not a string
	if err != nil {
		fmt.Scanln()
		color.Yellow("Invalid input. Transaction cancelled.")
	} else {
		// Calculate new balance
		newBalance := user.Balance + depositAmount
		// Updated the users balance value on the database
		result := database.DB.Model(&models.User{}).Where("ID = ?", user.ID).Update("Balance", newBalance)
		// Check if the update was done successfully
		if result.Error != nil {
			color.Red("Application Error: Could not process deposit.")
		} else {
			logErr := utils.LogTransaction(user.ID, "Deposit", depositAmount)
			if logErr != nil {
				color.Red(logErr.Error())
			} else {
				user.Balance = newBalance
				color.Green("Deposit processed successfully.")
			}
		}
	}
}
