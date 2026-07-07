package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	AppPort      string
	ResendAPIKey string
	AppURL       string
}
var App Config

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	App = Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		AppPort:    os.Getenv("APP_PORT"),
		ResendAPIKey: os.Getenv("RESEND_API_KEY"),
		AppURL:       os.Getenv("APP_URL"),
	}
}