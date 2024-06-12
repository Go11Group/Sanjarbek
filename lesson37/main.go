package main

import (
	"module/handler"
	"module/postgres"
)

func main() {
	db, err := postgres.ConnectDb()
	if err != nil {
		panic(err)
	}
	user := postgres.CreateUser(db)
	start := handler.ConnectGin(handler.Handler{User: user})
	start.Run()
}
