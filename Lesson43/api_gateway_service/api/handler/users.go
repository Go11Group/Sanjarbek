package handler

import (
	"encoding/json"
	"fmt"
	"gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	url := "http://localhost:8082/user/"
	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetUsers(ctx *gin.Context) {
	url := "http://localhost:8082/user/"

	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var user []models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response body"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/user/%s", id)

	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var user models.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing response body"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/user/%s", id)

	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User updated successfully"})
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	url := fmt.Sprintf("http://localhost:8082/user/%s", id)

	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Do(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		ctx.JSON(resp.StatusCode, gin.H{"error": "Unexpected status code"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "User deleted successfully"})
}
