package api

import (
	"database/sql"
	"user_service/api/handler"

	"github.com/gin-gonic/gin"
)

func Router(db *sql.DB) *gin.Engine {
	router := gin.Default()

	h := handler.NewHandler(db)

	user := router.Group("/user")
	user.POST("/create", h.CreateUser)
	user.GET("/getById/:id", h.GetUserById)
	user.GET("/GetAll", h.GetAllUser)
	user.PUT("/update/:id", h.UpdateUser)
	user.DELETE("/delete/:id", h.DeleteUser)

	return router
}
