package handler

import (
	"fmt"
	"metro_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateTerminal(ctx *gin.Context) {
	terminal := models.Terminal{}

	err := ctx.ShouldBindJSON(&terminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Terminal.Create(&terminal)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) GetTerminalByID(ctx *gin.Context) {
	TerminalId := ctx.Param("id")

	terminal, err := h.Terminal.GetById(TerminalId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Terminal not found: %s", TerminalId)})
		return
	}

	ctx.JSON(http.StatusOK, terminal)
}

func (h *handler) GetAllTerminals(c *gin.Context) {
	terminals, err := h.Terminal.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, terminals)
}

func (h *handler) UpdateTerminal(ctx *gin.Context) {
	TerminalId := ctx.Param("id")

	if TerminalId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Terminal ID"})
		return
	}

	var terminal models.Terminal
	if err := ctx.ShouldBindJSON(&terminal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTerminal, err := h.Terminal.Update(terminal)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateTerminal)
}

func (h *handler) DeleteTerminal(ctx *gin.Context) {
	TerminalId := ctx.Param("id")

	err := h.Terminal.Delete(TerminalId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete station: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "Terminal deleted succesfully"})
}
