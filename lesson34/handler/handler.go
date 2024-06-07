package handler

import (
	"module/storage/postgres"
	"net/http"
)


type Handler struct {
	Book *postgres.BookRepo
}

type Users struct {
	User *postgres.UsersRepo
}

type Products struct {
	Product *postgres.ProductsRepo
}

func NewHandler(users *postgres.UsersRepo, products *postgres.ProductsRepo) *http.ServeMux {
    mux := http.NewServeMux()

    userHandler := Users{User: users}
    productHandler := Products{Product: products}

    mux.HandleFunc("/CreateUsers/", userHandler.UserCreate)
    mux.HandleFunc("/ReadUsers/", userHandler.UserRead)
    mux.HandleFunc("/UpdateUsers/", userHandler.UserUpdate)
    mux.HandleFunc("/DeleteUsers/", userHandler.UserDelete)

    mux.HandleFunc("/CreateProducts/", productHandler.ProductCreate)
    mux.HandleFunc("/ReadProducts/", productHandler.ProductRead)
    mux.HandleFunc("/UpdateProducts/", productHandler.ProductUpdate)
    mux.HandleFunc("/DeleteProducts/", productHandler.ProductDelete)

    return mux
}

