package handler

import (
	"fmt"
	"module/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.User.Create(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Query("id")

	ID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", userID)})
		return
	}

	user, err := h.User.GetUserByID(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", userID)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetUserId(c *gin.Context) {
	userID := c.Param("id")

	ID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid user ID: %s", userID)})
		return
	}

	user, err := h.User.GetUserByID(ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", userID)})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) GetAllUser(c *gin.Context) {
	resp, err := h.User.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	var user model.User
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

	ID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid user ID: %s", userID)})
		return
	}

	err = h.User.Delete(ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete user: %v", err)})
		return
	}

	c.Status(http.StatusOK)
}
