package main

import (
	"module/handler"
	"module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	book := postgres.NewBookRepo(db)

	server := handler.NewHandler(handler.Handler{Book: book})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
