package api

import (
	"casbin/api/handler"
	"casbin/api/middleware"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(redisHandler *handler.RedisHandler, enforcer *casbin.Enforcer) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.CasbinMiddleware(enforcer))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	books := router.Group("/api/v1/books")
	{
		books.POST("", redisHandler.CreateBook)
		books.GET("/:id", redisHandler.GetBook)
		books.PUT("/:id", redisHandler.UpdateBook)
		books.DELETE("/:id", redisHandler.DeleteBook)
	}

	return router
}
