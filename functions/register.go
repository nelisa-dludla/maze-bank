package functions

import (
	"fmt"
	"maze/database"
	"maze/models"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Register() models.User {
	// Collect user data
	var user models.User
	color.Cyan("===== Maze Bank - Register =====\n\nPlease fill in your details below.")
	fmt.Println()
	fmt.Print("First Name (Required): ")
	fmt.Scanf("%s", &user.FirstName)
	fmt.Print("Last Name (Required): ")
	fmt.Scanf("%s", &user.LastName)
	fmt.Print("Email (Required): ")
	fmt.Scanf("%s", &user.Email)
	fmt.Print("Password (Minimum of 8 characters): ")
	fmt.Scanf("%s", &user.Password)

	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		color.Red(errors.Error())
		color.Red("User registration cancelled.")
		return models.User{}
	}

	fmt.Println("\nYour details have been recorded as follows:")
	fmt.Printf(`
First Name: %s
Last Name : %s
Email     : %s
Password  : %s
`, user.FirstName, user.LastName, user.Email,  user.Password)

	// Hash password
	hash, hashErr := HashPassword(user.Password)
	if hashErr != nil {
		color.Red("Error hashing password.")
	}
	user.Password = hash

	// Add user to the database
	result := database.DB.Create(&user)
	if result.Error != nil {
		color.Red("User registration failed. Try again.")
		return models.User{}
	}

	color.Green("User registered successfully.")
	return user
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

