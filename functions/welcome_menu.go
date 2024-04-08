package functions

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func WelcomeMenu() (int, error) {
	var userResponse int
	// Welcome Menu
	color.Cyan("===== Welcome to Maze Bank =====")
	fmt.Println()
	fmt.Print(`
(1) Log In
(2) Register

(9) Quit

> `)
	// Get user input
	_, err := fmt.Scanf("%d", &userResponse)
	if err != nil {
		fmt.Scanln()
		return -1, errors.New("Invalid input. Please enter number 1 for Log In or 2 for Register.")
	}
	return userResponse, nil
}
