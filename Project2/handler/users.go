package handler

import (
	"fmt"
	"module/model"
	"module/storage/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	User postgres.UserRepository
}

func NewHandler(userRepo postgres.UserRepository) *Handler {
	return &Handler{User: userRepo}
}

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


func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Query("id")

	user, err := h.User.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", userID)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUserId(c *gin.Context) {
	userID := c.Param("id")

	user, err := h.User.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", userID)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.User.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user model.Users
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.User.Update(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	err := h.User.Delete(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete user: %v", err)})
		return
	}

	c.Status(http.StatusOK)
}
