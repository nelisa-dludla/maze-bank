package main

import (
	"maze/database"
	"maze/functions"
	"os"

	"github.com/fatih/color"
)

func init() {
	database.ConnectToDb()
}

func main() {

	for {
		response, err := functions.WelcomeMenu()
		if err != nil {
			color.Yellow(err.Error())
			continue
		}

		if response == 9 {
			color.Green("\nClosing Maze Bank...")
			os.Exit(0)
		}

		functions.Access(response)
	}
}
