package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/create", h.CreateUser)
		users.GET("/get", h.GetUser)
		users.GET("/get/:id", h.GetUserId)
		users.GET("/getall", h.GetAllUsers)
		users.PUT("/update/:id", h.UpdateUser)
		users.DELETE("/delete/:id", h.DeleteUser)
	}
}
