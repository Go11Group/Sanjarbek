package api

import (
	"github.com/gin-gonic/gin"
	"casbin/api/handler"
)

func Router(redisHandler *handler.RedisHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/api/book", redisHandler.CreateBook)
	r.GET("/api/book/:id", redisHandler.GetBook)
	r.PUT("/api/book", redisHandler.UpdateBook)
	r.DELETE("/api/book/:id", redisHandler.DeleteBook)
	r.POST("/api/book/buy", redisHandler.BuyBook)
	r.POST("/api/book/sell", redisHandler.SellBook)

	return r
}
