package handler

import (
	"module/storage/postgres"
	"net/http"
)


type Handler struct {
	Book *postgres.BookRepo
}


func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/book/", handler.book)

	return &http.Server{Handler: mux}
}

type Book struct {
	Name, Author, Publisher string
}
