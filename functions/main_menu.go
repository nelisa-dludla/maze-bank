package functions

import (
	"fmt"
	"maze/models"
	"os"

	"github.com/fatih/color"
)

func MainMenu(user *models.User) {
	var mainMenuResponse int
	keepLooping := true

	for keepLooping {
		color.Cyan("===== Maze Bank - Account =====")
		fmt.Println()
		fmt.Printf(
`
Welcome Back, %s

Current Balance: %.2f

(1) Deposit
(2) Withdraw

(8) Print Statement
(9) Log Out

> `, user.FirstName, user.Balance)
	
		fmt.Scanf("%d", &mainMenuResponse)
	
		if mainMenuResponse == 9 {
			color.Green("\nLogged Out.")
			color.Green("Closing Maze Bank...")
			os.Exit(0)
		}
	
		switch mainMenuResponse {
		case 1:
			Deposit(user)
		case 2:
			Withdraw(user)
		case 8:
			PrintStatement(*user)
		default:
			color.Yellow("Invalid input. Please enter number 1 for Deposit, 2 for Withdraw, 8 for Print Statement or 9 for Log Out.")
		}
	}
}
