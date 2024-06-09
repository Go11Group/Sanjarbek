package handler

import (
	"module/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)



type Handler struct {
	User *postgres.UserRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := mux.NewRouter()

	mux.HandleFunc("/createUser", handler.CreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", handler.GetUserId).Methods("GET")
	mux.HandleFunc("/users", handler.GetAllUser).Methods("GET")
	mux.HandleFunc("/updateUser/{id}", handler.UpdateUser).Methods("PUT")
	mux.HandleFunc("/deleteUser/{id}", handler.DeleteUser).Methods("DELETE")

	return &http.Server{Addr: ":8080", Handler: mux}
}