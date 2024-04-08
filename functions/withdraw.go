package functions

import (
	"fmt"
	"maze/database"
	"maze/models"
	"maze/utils"

	"github.com/fatih/color"
)

func Withdraw(user *models.User) {
	color.Cyan("===== Maze Bank - Withdraw =====")
	fmt.Println()

	var withdrawalAmount float64
	// Get withdrawal amount
	fmt.Print("Enter amount you would like to withdraw: ")
	_, err := fmt.Scanf("%f", &withdrawalAmount)
	// Check if anything but numbers were entered
	if err != nil {
		fmt.Scanln()
		color.Yellow("Invalid input. Transaction cancelled.")
	} else {
		// Check if there are enough funds to Withdraw
		fmt.Scanln()
		if withdrawalAmount > user.Balance {
			color.Yellow("Insufficient funds. Transaction cancelled.")
		} else {
			// Make the withdrawal
			newBalance := user.Balance - withdrawalAmount
			result := database.DB.Model(&models.User{}).Where("ID = ?", user.ID).Update("Balance", newBalance)
			// Check if transaction was successful
			if result.Error != nil {
				color.Red("Application Error: Could not process withdrawal")
			} else {
				// Log transaction
				logErr := utils.LogTransaction(user.ID, "Withdrawal", withdrawalAmount)
				if logErr != nil {
					fmt.Println(logErr)
				} else {
					user.Balance = newBalance
					color.Green("Withdrawal processed successfully.")
				}
			}
		}
	}
}
