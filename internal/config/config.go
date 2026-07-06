package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	AppPort    string
	SMTPHost   string
	SMTPPort   string
	SMTPEmail  string
	SMTPPass   string
	AppURL     string
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
		SMTPHost:   os.Getenv("SMTP_HOST"),
		SMTPPort:   os.Getenv("SMTP_PORT"),
		SMTPEmail:  os.Getenv("SMTP_EMAIL"),
		SMTPPass:   os.Getenv("SMTP_PASSWORD"),
		AppURL:     os.Getenv("APP_URL"),
	}
}