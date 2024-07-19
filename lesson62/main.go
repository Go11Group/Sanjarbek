package main

import (
	"log"
	"casbin/api"
	"casbin/api/handler"
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Initialize Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// Initialize Casbin enforcer
	enforcer, err := casbin.NewEnforcer("config/casbin_model.conf", "config/casbin_policy.csv")
	if err != nil {
		log.Fatalf("Failed to create Casbin enforcer: %v", err)
	}

	redisHandler := &handler.RedisHandler{
		RedisClient: redisClient,
		Enforcer:    enforcer,
	}

	r := api.Router(redisHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
