package handler

import (
	"encoding/json"
	"fmt"
	"gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetBalance(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	url := fmt.Sprintf("http://localhost:8081/balance?user_id=%s", userId)

	req, err := http.NewRequest(ctx.Request.Method, url, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var balance models.Balance
	if err = json.NewDecoder(resp.Body).Decode(&balance); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(resp.StatusCode, balance)
}

func (h *Handler) CheckBalance(ctx *gin.Context) {
	userId := ctx.Query("user_id")
	url := fmt.Sprintf("http://localhost:8081/check_balance?user_id=%s", userId)

	req, err := http.NewRequest(ctx.Request.Method, url, nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := h.Client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result models.BalanceResp
	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(resp.StatusCode, result)
}