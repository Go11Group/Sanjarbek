package main

import (
	"casbin/api"
	"casbin/api/handler"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"log"
)

func main() {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Initialize Casbin enforcer
	enforcer, err := casbin.NewEnforcer("config/casbin_model.conf", "config/casbin_policy.csv")
	if err != nil {
		log.Fatalf("Failed to create Casbin enforcer: %v", err)
	}

	// Initialize Redis handler
	redisHandler := &handler.RedisHandler{
		RedisClient: redisClient,
	}

	// Initialize the router with Casbin middleware
	r := api.Router(redisHandler, enforcer)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
