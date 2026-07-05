package main

import (
	"fmt"

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

	// Routes setup
	routes.Setup(r)

	// Server start
	fmt.Println("🚀 Bazarly server running on port " + config.App.AppPort)
	r.Run(":" + config.App.AppPort)
}