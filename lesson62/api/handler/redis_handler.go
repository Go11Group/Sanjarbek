package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type RedisHandler struct {
	RedisClient *redis.Client
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with title and author
// @Tags Book
// @Accept json
// @Produce json
// @Param book body Book true "Book object"
// @Success 201 {object} Book
// @Failure 400 {object} gin.H
// @Router /api/v1/books [post]
func (h *RedisHandler) CreateBook(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	book.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	err := h.RedisClient.HSet(ctx, "books", book.ID, book).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create book"})
		return
	}

	c.JSON(http.StatusCreated, book)
}

// GetBook godoc
// @Summary Get a book by ID
// @Description Retrieve a book by its ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} Book
// @Failure 404 {object} gin.H
// @Router /api/v1/books/{id} [get]
func (h *RedisHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	bookData, err := h.RedisClient.HGet(ctx, "books", id).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get book"})
		return
	}

	var book Book
	if err := json.Unmarshal([]byte(bookData), &book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to unmarshal book data"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update a book by ID
// @Description Update a book's title or author by ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Param book body Book true "Book object"
// @Success 200 {object} Book
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/books/{id} [put]
func (h *RedisHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	book.ID = id
	err := h.RedisClient.HSet(ctx, "books", id, book).Err()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update book"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete a book by ID
// @Description Delete a book by its ID
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "Book ID"
// @Success 200 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /api/v1/books/{id} [delete]
func (h *RedisHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	err := h.RedisClient.HDel(ctx, "books", id).Err()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
