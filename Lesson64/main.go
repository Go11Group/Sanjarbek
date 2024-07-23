package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	model "redis/models"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	companies := []string{"Amazon", "BYD", "Google", "Microsoft", "Apple"}
	pubsub := rdb.Subscribe(ctx, companies...)
	defer pubsub.Close()

	_, err := pubsub.Receive(ctx)
	if err != nil {
		log.Fatal(err)
	}

	ch := pubsub.Channel()
	for msg := range ch {
		price, err := strconv.Atoi(msg.Payload)
		if err != nil {
			log.Printf("Invalid price received: %s", msg.Payload)
			continue
		}

		trackPrice(ctx, rdb, msg.Channel, price)
		fmt.Printf("Received message from %s: Price: %d, Time: %s\n", msg.Channel, price, time.Now().Format("2006-01-02 15:04:05"))
	}
}

func trackPrice(ctx context.Context, rdb *redis.Client, company string, price int) {
	highKey := fmt.Sprintf("%s:high", company)
	lowKey := fmt.Sprintf("%s:low", company)

	timestamp := time.Now().Format("2006-01-02 15:04:05")

	highPriceData := model.PriceData{}
	highPriceJSON, _ := rdb.Get(ctx, highKey).Result()
	if highPriceJSON != "" {
		json.Unmarshal([]byte(highPriceJSON), &highPriceData)
	}

	if price > highPriceData.Price {
		newHighPriceData := model.PriceData{
			Price:     price,
			Timestamp: timestamp,
		}
		newHighPriceJSON, _ := json.Marshal(newHighPriceData)
		rdb.Set(ctx, highKey, newHighPriceJSON, 0)
	}

	lowPriceData := model.PriceData{}
	lowPriceJSON, _ := rdb.Get(ctx, lowKey).Result()
	if lowPriceJSON != "" {
		json.Unmarshal([]byte(lowPriceJSON), &lowPriceData)
	}

	if lowPriceData.Price == 0 || price < lowPriceData.Price {
		newLowPriceData := model.PriceData{
			Price:     price,
			Timestamp: timestamp,
		}
		newLowPriceJSON, _ := json.Marshal(newLowPriceData)
		rdb.Set(ctx, lowKey, newLowPriceJSON, 0)
	}
}
