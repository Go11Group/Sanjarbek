package main

import (
	"gateway/api"
	"gateway/api/handler"
	"log"
)

func main() {
	router := api.Router(handler.Handler{})

	log.Fatal(router.Run(":8080"))
}