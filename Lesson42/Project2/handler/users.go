package handler

import (
	"fmt"
	"module/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// user yaratish
func (h *Handler) CreateUser(c *gin.Context) {
	var user model.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.User.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// id boyicha userni topish
func (h *Handler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.User.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", userID)})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers handles retrieving all users
func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.User.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetAllUsersFiltered handles retrieving users based on query parameters
func (h *Handler) GetAllUsersFiltered(c *gin.Context) {
	var filter model.UserGetAll

	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid query parameters: %v", err)})
		return
	}
	fmt.Println(filter)

	users, err := h.User.GetAllUsersFiltered(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, users)
}

// UpdateUser handles updating an existing user
func (h *Handler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	var user model.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.User.Update(user, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// DeleteUser handles deleting a user by their ID
func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := h.User.Delete(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func (h *Handler) SearchUsers(c *gin.Context) {
	var user model.UsersGet
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.User.SearchUsers(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User is not avialble: %v", err)})
		return
	}
	c.JSON(http.StatusOK, resp)
}
