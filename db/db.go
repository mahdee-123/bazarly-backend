package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mahdee-123/bazarly-backend/config"
	_ "github.com/lib/pq"
)

var DB *sql.DB


func Connect() {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.App.DBHost,
		config.App.DBPort,
		config.App.DBUser,
		config.App.DBPassword,
		config.App.DBName,
	)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB ping failed:", err)
	}

	fmt.Println("✅ Database connected!")
}