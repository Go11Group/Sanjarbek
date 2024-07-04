package middleware

import (
	"api-service/api/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		if auth == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid tokin",
			})
			ctx.Abort()
			return
		}

		valid, err := token.ValidateToken(auth)
		if err != nil || !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": err.Error(),
			})
			ctx.Abort()
			return
		}

		claims, err := token.ExtractClaim(auth)

		if err != nil || !valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"Error": err.Error(),
			})
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		ctx.Next()
	}
}
