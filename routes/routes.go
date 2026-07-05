package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/handlers"
	"github.com/mahdee-123/bazarly-backend/middleware"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api")

	// Public routes — no auth
	sellers := api.Group("/sellers")
	{
		sellers.POST("/signup", handlers.SellerSignup)
		sellers.POST("/login", handlers.SellerLogin)
	}

	// Protected routes — JWT required
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/sellers/me", handlers.GetSellerProfile)
	}
}