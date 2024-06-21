package handler

import "net/http"

type Handler struct {
	Client *http.Client
}

func NewHandler() *Handler {
	return &Handler{&http.Client{}}
}
