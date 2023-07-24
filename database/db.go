package database

import (
	"github/ericoliveiras/basic-crud-go/config"
	"github/ericoliveiras/basic-crud-go/models"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file")
	}

	dsn := config.DbURL()
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{}, &models.Task{})
}
