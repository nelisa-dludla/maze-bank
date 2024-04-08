package functions

import (
	"errors"
	"fmt"
	"maze/database"
	"maze/models"
	"os"

	"github.com/fatih/color"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
	"gorm.io/gorm"
)

type LoginDetails struct {
	email string `validate:"required"`
	password string `validate:"required"`
}

func Login() *models.User {
	var details LoginDetails

	color.Cyan("===== Maze Bank - Login =====")
	fmt.Print("\nEnter email: ")
	fmt.Scanf("%s", &details.email)
	fmt.Print("Enter password: ")
	password, readPasswordErr := term.ReadPassword(int(os.Stdin.Fd()))
	if readPasswordErr != nil {
		color.Red("Error reading password.")
	}
	details.password = string(password)

	validate := validator.New()
	err := validate.Struct(details)
	if err != nil {
		errors := err.(validator.ValidationErrors)
		fmt.Println(errors)
		color.Red("\nUser login cancelled.")
		return &models.User{}
	}

	var user *models.User
	result := database.DB.Where("email = ?", details.email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		color.Yellow("\nIncorrect email or password.")
		return &models.User{}
	}
	
	if !CheckPasswordHash(details.password, user.Password) {
		color.Yellow("\nIncorrect email or password.")
		return &models.User{}
	}

	color.Green("\nLogged in succesfully.")
	return user
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
