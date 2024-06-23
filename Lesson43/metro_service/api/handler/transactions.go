package handler

import (
	"fmt"
	"metro_service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateTransaction(ctx *gin.Context) {
	transaction := models.Transaction{}

	err := ctx.ShouldBindJSON(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	err = h.Transactions.Create(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		fmt.Println("error:", err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, "OKAY")
}

func (h *handler) GetTransactionByID(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	station, err := h.Transactions.GetById(transactionId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Transaction not found: %s", transactionId)})
		return
	}

	ctx.JSON(http.StatusOK, station)
}

func (h *handler) GetAllTransactions(c *gin.Context) {
	stations, err := h.Transactions.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, stations)
}

func (h *handler) UpdateTransaction(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	if transactionId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	UpdateTransaction, err := h.Transactions.Update(&transaction)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, UpdateTransaction)
}

func (h *handler) DeleteTransaction(ctx *gin.Context) {
	transactionId := ctx.Param("id")

	err := h.Transactions.Delete(transactionId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete transaction: %v", err)})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"massage": "transaction deleted succesfully"})
}
