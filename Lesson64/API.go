package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// PriceData represents the structure for price information
type PriceData struct {
	Price     int    `json:"price"`
	Timestamp string `json:"timestamp"`
}

var rdb *redis.Client
var ctx = context.Background()

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	router := gin.Default()
	router.GET("/price/:company", getPrice)
	log.Fatal(router.Run(":8080"))
}

func getPrice(c *gin.Context) {
	company := c.Param("company")
	highKey := fmt.Sprintf("%s:high", company)
	lowKey := fmt.Sprintf("%s:low", company)

	// Retrieve and unmarshal high price data
	highPriceJSON, err := rdb.Get(ctx, highKey).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No high price found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var highPriceData PriceData
	if err := json.Unmarshal([]byte(highPriceJSON), &highPriceData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal high price data"})
		return
	}

	// Retrieve and unmarshal low price data
	lowPriceJSON, err := rdb.Get(ctx, lowKey).Result()
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No low price found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var lowPriceData PriceData
	if err := json.Unmarshal([]byte(lowPriceJSON), &lowPriceData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unmarshal low price data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"company":       company,
		"highPrice":     highPriceData.Price,
		"highTimestamp": highPriceData.Timestamp,
		"lowPrice":      lowPriceData.Price,
		"lowTimestamp":  lowPriceData.Timestamp,
	})
}
