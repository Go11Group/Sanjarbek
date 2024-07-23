package main

import (
	"gin/api"
	"log"
)

func main() {

	// run

	router := api.Router()

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}

}
