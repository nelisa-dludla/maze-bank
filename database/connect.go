package database

import (
	"fmt"
	"maze/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDb() {
	db, err := gorm.Open(sqlite.Open("maze_bank.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("Failed to connect to database.")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Transaction{})

	DB = db
}
