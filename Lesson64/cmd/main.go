package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	companies := []string{"Amazon", "BYD", "Google", "Microsoft", "Apple"}

	for {
		for _, company := range companies {
			price := rand.Intn(1000)
			err := rdb.Publish(ctx, company, price).Err()
			if err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
