package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/config"
	"github.com/mahdee-123/bazarly-backend/db"
	"github.com/mahdee-123/bazarly-backend/routes"
)

func main() {
	// Config load
	config.Load()

	// Database connect
	db.Connect()

	// Gin router
	r := gin.Default()


		// CORS setup
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Routes setup
	routes.Setup(r)

	// Server start
	fmt.Println("🚀 Bazarly server running on port " + config.App.AppPort)
	r.Run(":" + config.App.AppPort)
}