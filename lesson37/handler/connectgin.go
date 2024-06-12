package handler

import (
	"module/postgres"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	User *postgres.UserRepo
}

func ConnectGin(handler Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/users", handler.GetAllUser)
	return r
}
