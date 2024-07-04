package handler

import (
	"api-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *http.Handler) CreateBook(c *gin.Context) {
	var req models.RegisterReq

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Invalid request payload": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tokens)
}
