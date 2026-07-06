package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/mahdee-123/bazarly-backend/internal/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Header theke token nao
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Error(c, http.StatusUnauthorized, "Authorization header required")
			c.Abort()
			return
		}

		// "Bearer <token>" format check
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(c, http.StatusUnauthorized, "Invalid token format")
			c.Abort()
			return
		}

		// Token verify
		claims, err := utils.VerifyToken(parts[1])
		if err != nil {
			utils.Error(c, http.StatusUnauthorized, "Invalid or expired token")
			c.Abort()
			return
		}

		// seller_id context e set koro
		c.Set("seller_id", claims.SellerID)
		c.Next()
	}
}