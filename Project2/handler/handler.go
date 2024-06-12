package handler

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the routes for handling user-related operations
func (h *Handler) RegisterRoutes(r *gin.Engine) {
	// Create a route group for user-related endpoints
	users := r.Group("/users")
	{
		// Define routes for creating, getting, updating, and deleting users
		users.POST("/create", h.CreateUser)     // Route for creating a new user
		users.GET("/get", h.GetUser)            // Route for getting user(s)
		users.GET("/get/:id", h.GetUserId)      // Route for getting a user by ID
		users.GET("/getall", h.GetAllUsers)     // Route for getting all users
		users.PUT("/update/:id", h.UpdateUser)  // Route for updating a user
		users.DELETE("/delete/:id", h.DeleteUser) // Route for deleting a user
	}
}
