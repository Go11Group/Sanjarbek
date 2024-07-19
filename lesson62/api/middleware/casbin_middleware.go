package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthMiddleware checks if User-ID is provided
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("User-ID")
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No user ID header provided"})
			c.Abort()
			return
		}

		// Set the user ID in the context to use it later
		c.Set("userID", userID)
		c.Next()
	}
}
