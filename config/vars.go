package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var err = godotenv.Load()

func DbURL() string {
	if err != nil {
		panic("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPass, dbName, dbPort)
}

func BaseURL() string {
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	return fmt.Sprintf("%s:%s", host, port)
}