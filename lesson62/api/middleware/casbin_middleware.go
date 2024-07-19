package middleware

import (
	"context"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CasbinMiddleware(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		sub := c.GetHeader("X-User-ID")
		if sub == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No user ID header provided"})
			c.Abort()
			return
		}

		obj := c.Request.URL.Path
		act := c.Request.Method

		allowed, err := enforcer.Enforce(context.Background(), sub, obj, act)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error in Casbin enforcement"})
			c.Abort()
			return
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{"message": "Forbidden"})
			c.Abort()
			return
		}

		c.Next()
	}
}
