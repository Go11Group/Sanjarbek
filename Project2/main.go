package main

import (
	"log"
	"module/handler"
	"module/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)

	h := handler.NewHandler(*userRepo)

	r := gin.Default()
	h.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
