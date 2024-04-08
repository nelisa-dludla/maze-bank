package functions

import (
	"fmt"
	"maze/database"
	"maze/models"
	"os"
	"time"

	"github.com/fatih/color"
)

func PrintStatement(user models.User) {
	color.Green("Processing...")
	// Find users transactions
	var currentUser models.User
	result := database.DB.Preload("Transactions").Find(&currentUser, user.ID)
	if result.Error != nil {
		color.Red("Error finding user: ", result.Error)
	} else {
		// Create file/create statement
		now := time.Now()
		formattedDate := now.Format("2006_01_02_15_04_05")
		fileName := fmt.Sprintf("maze_bank_statement_%s", formattedDate)
		file, err := os.Create(fileName)

		if err != nil {
			color.Red("Error creating statement:", err)
		}
		defer file.Close()

		var transactions string
		// Loop through transactions to form a string
		for _, transaction := range currentUser.Transactions {
			transactionString := fmt.Sprintf("%s - %s - %s - %.2f\n", transaction.ID, transaction.CreatedAt.Format("2006-01-02 15:04:05"), transaction.Description, transaction.Amount)
			transactions += transactionString
		}

		data := fmt.Sprintf(`===== Maze Bank - Statement =====

Printed: %s
Current Balance: %.2f

%s`, now.Format("2006-01-02 15:02"), currentUser.Balance, transactions)
		// Write to file
		_, fileErr := file.Write([]byte(data))
		if fileErr != nil {
			color.Red("Error writing to file: ", fileErr)
		} else {
			os.Chmod(fileName, 0444)
			color.Green("Statement Printed.")
		}
	}
}
