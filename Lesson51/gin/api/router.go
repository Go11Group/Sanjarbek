package api

import (
	"gin/api/handler"
	"gin/api/middleware"
	"gin/storage/postgres"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	userRepo := postgres.UserRepo{}
	handler := handler.NewHandler(&userRepo)

	router := gin.Default()

	user := router.Group("/user")
	auth := router.Group("/auth")

	user.Use(middleware.Middleware)

	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)

	user.POST("", handler.CreateUser)
	user.GET("", handler.GetUser)

	return router
}
