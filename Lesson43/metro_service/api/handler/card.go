package handler

import (
	"fmt"
	"metro_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateCard(ctx *gin.Context) {
	stn := models.CreateCard{}

	err := ctx.ShouldBindJSON(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Card.Create(&stn)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) GetCardByID(ctx *gin.Context) {
	CardId := ctx.Param("id")

	card, err := h.Card.GetById(CardId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Station not found: %s", CardId)})
		return
	}

	ctx.JSON(http.StatusOK, card)
}

func (h *handler) GetAllCards(ctx *gin.Context) {
	cards, err := h.Card.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, cards)
}

func (h *handler) UpdateCard(ctx *gin.Context) {
	cardID := ctx.Param("id")

	if cardID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid card ID"})
		return
	}

	var card models.Card
	if err := ctx.ShouldBindJSON(&card); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateCard, err := h.Card.UpdateCard(card)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updateCard)
}

func (h *handler) DeleteCard(ctx *gin.Context) {
	cardID := ctx.Param("id")

	err := h.Card.DeleteCard(cardID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete station: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "Card deleted succesfully"})
}
