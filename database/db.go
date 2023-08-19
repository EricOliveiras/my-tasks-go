package database

import (
	"github/ericoliveiras/basic-crud-go/config"
	"github/ericoliveiras/basic-crud-go/models"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file")
	}

	dsn := config.DbURL()

	maxRetries := 5
	retryInterval := time.Second * 5

	var dbErr error

	for retries := 0; retries < maxRetries; retries++ {
		DB, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if dbErr == nil {
			break
		}

		log.Printf("Attempt %d of %d: Error connecting to the database: %s\n", retries+1, maxRetries, dbErr)
		if retries < maxRetries-1 {
			log.Printf("Retrying in %s...\n", retryInterval)
			time.Sleep(retryInterval)
		}
	}

	if dbErr != nil {
		log.Panic("Failed to connect to the database after multiple attempts.")
	}

	DB.AutoMigrate(&models.User{}, &models.Task{})
}
