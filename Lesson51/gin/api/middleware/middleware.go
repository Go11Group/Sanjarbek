package middleware

import (
	"gin/api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	tokenStr := c.GetHeader("authorization")

	if tokenStr == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	err := token.ExtaracClaims(tokenStr)
	if err != nil {
		c.AbortWithStatusJSON(403, gin.H{
			"error": "invalid token",
		})
		return
	}

	c.Next()

}
