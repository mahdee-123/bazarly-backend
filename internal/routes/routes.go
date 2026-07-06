package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/internal/middleware"
	"github.com/mahdee-123/bazarly-backend/internal/product"
	"github.com/mahdee-123/bazarly-backend/internal/sellers"
	"github.com/mahdee-123/bazarly-backend/internal/store"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api")

	// Public routes — no auth
	sellerRoutes := api.Group("/sellers")
	{
		sellerRoutes.POST("/signup", sellers.SellerSignupHandler)
		sellerRoutes.POST("/login", sellers.SellerLoginHandler)
		sellerRoutes.GET("/verify-email", sellers.VerifyEmailHandler)
	}

	// Protected routes — JWT required
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// Seller
		auth.GET("/sellers/me", sellers.GetSellerProfileHandler)

		// Store list + create
		auth.POST("/stores", store.CreateStoreHandler)
		auth.GET("/stores", store.GetMyStoresHandler)

		// Store detail + products — sob ekta group e
		storeGroup := auth.Group("/stores/:storeId")
		{
			storeGroup.GET("", store.GetStoreHandler)
			storeGroup.DELETE("", store.DeleteStoreHandler)

			// Products
			storeGroup.POST("/products", product.CreateHandler)
			storeGroup.GET("/products", product.ListHandler)
			storeGroup.GET("/products/:productId", product.GetHandler)
			storeGroup.PUT("/products/:productId", product.UpdateHandler)
			storeGroup.DELETE("/products/:productId", product.DeleteHandler)
		}
	}
}