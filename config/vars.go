package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type CookieConfigModel struct {
	MaxAge   int
	Path     string
	Domain   string
	Secure   bool
	HttpOnly bool
}

var err = godotenv.Load()

func DbURL() string {
	if err != nil {
		log.Panic("Error loading .env file")
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
		log.Panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	return fmt.Sprintf("%s:%s", host, port)
}

func CookieConfig() CookieConfigModel {
	cookieConfig := CookieConfigModel{
		MaxAge:   604800,
		Path:     "/",
		Domain:   "*",
		Secure:   true,
		HttpOnly: true,
	}

	return cookieConfig
}
