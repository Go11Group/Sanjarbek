package handler

import (
	"context"
	"net/http"
	"strconv"

	"casbin/models"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type RedisHandler struct {
	RedisClient *redis.Client
	Enforcer    *casbin.Enforcer
}

// CreateBook creates a new book
func (h *RedisHandler) CreateBook(c *gin.Context) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No user ID header provided"})
		return
	}

	// Check authorization with Casbin
	ok, _ := h.Enforcer.Enforce(userID, "books", "create")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// Set book in Redis
	bookKey := "book:" + book.ID
	bookValue := map[string]interface{}{
		"title":    book.Title,
		"author":   book.Author,
		"quantity": book.Quantity,
		"price":    book.Price,
	}

	_, err := h.RedisClient.HSet(context.Background(), bookKey, bookValue).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book created successfully"})
}

// GetBook retrieves a book by ID
func (h *RedisHandler) GetBook(c *gin.Context) {
	bookID := c.Param("id")

	// Get book from Redis
	bookKey := "book:" + bookID
	bookData, err := h.RedisClient.HGetAll(context.Background(), bookKey).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	if len(bookData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	book := models.Book{
		ID:       bookID,
		Title:    bookData["title"],
		Author:   bookData["author"],
		Quantity: parseInt(bookData["quantity"]),
		Price:    parseFloat(bookData["price"]),
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book
func (h *RedisHandler) UpdateBook(c *gin.Context) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No user ID header provided"})
		return
	}

	// Check authorization with Casbin
	ok, _ := h.Enforcer.Enforce(userID, "books", "update")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	// Update book in Redis
	bookKey := "book:" + book.ID
	bookValue := map[string]interface{}{
		"title":    book.Title,
		"author":   book.Author,
		"quantity": book.Quantity,
		"price":    book.Price,
	}

	_, err := h.RedisClient.HSet(context.Background(), bookKey, bookValue).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully"})
}

// DeleteBook deletes a book by ID
func (h *RedisHandler) DeleteBook(c *gin.Context) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No user ID header provided"})
		return
	}

	// Check authorization with Casbin
	ok, _ := h.Enforcer.Enforce(userID, "books", "delete")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return
	}

	bookID := c.Param("id")
	bookKey := "book:" + bookID

	_, err := h.RedisClient.Del(context.Background(), bookKey).Result()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

// BuyBook buys a book by reducing its quantity
func (h *RedisHandler) BuyBook(c *gin.Context) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No user ID header provided"})
		return
	}

	// Check authorization with Casbin
	ok, _ := h.Enforcer.Enforce(userID, "books", "buy")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return
	}

	var req models.BuyBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	bookKey := "book:" + req.ID
	bookData, err := h.RedisClient.HGetAll(context.Background(), bookKey).Result()
	if err != nil || len(bookData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	currentQuantity := parseInt(bookData["quantity"])
	if req.Quantity > currentQuantity {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Not enough quantity available"})
		return
	}

	newQuantity := currentQuantity - req.Quantity
	_, err = h.RedisClient.HSet(context.Background(), bookKey, "quantity", newQuantity).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating book quantity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book bought successfully"})
}

// SellBook sells a book by increasing its quantity
func (h *RedisHandler) SellBook(c *gin.Context) {
	userID := c.GetHeader("user-id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No user ID header provided"})
		return
	}

	// Check authorization with Casbin
	ok, _ := h.Enforcer.Enforce(userID, "books", "sell")
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
		return
	}

	var req models.SellBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	bookKey := "book:" + req.ID
	bookData, err := h.RedisClient.HGetAll(context.Background(), bookKey).Result()
	if err != nil || len(bookData) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	currentQuantity := parseInt(bookData["quantity"])
	newQuantity := currentQuantity + req.Quantity
	_, err = h.RedisClient.HSet(context.Background(), bookKey, "quantity", newQuantity).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error updating book quantity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book sold successfully"})
}

// Utility functions to parse data
func parseInt(value string) int32 {
	i, _ := strconv.Atoi(value)
	return int32(i)
}

func parseFloat(value string) float64 {
	f, _ := strconv.ParseFloat(value, 64)
	return f
}
