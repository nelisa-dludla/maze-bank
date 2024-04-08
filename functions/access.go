package functions

import (
	"github.com/fatih/color"
)

func Access(response int) {
		switch response {
		case 1:
			user := Login()
			if user.ID != 0 {
				MainMenu(user)
			}
		case 2:
			Register()
		default:
			color.Yellow("Invalid option. Please try again")
		}
}
