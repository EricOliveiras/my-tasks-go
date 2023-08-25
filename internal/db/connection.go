package db

import (
	"fmt"
	"github/ericoliveiras/basic-crud-go/internal/config"
	"github/ericoliveiras/basic-crud-go/internal/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		cfg.DB.Host, cfg.DB.User, cfg.DB.Password, cfg.DB.Name, cfg.DB.Port)

	maxRetries := 5
	retryInterval := time.Second * 5

	var db *gorm.DB
	var dbErr error

	for retries := 0; retries < maxRetries; retries++ {
		db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})
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

	db.AutoMigrate(&models.User{}, &models.Task{})

	return db
}
