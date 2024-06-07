package main

import (
    "module/handler"
    "module/storage/postgres"
    "net/http"
)

func main() {
    db, err := postgres.ConnectDB()
    if err != nil {
        panic(err)
    }
    usersRepo := postgres.NewUserRepo(db)

    usersHandler := &handler.Users{
        User: usersRepo,
    }

    mux := http.NewServeMux()
    mux.HandleFunc("/users/create", usersHandler.UserCreate)
    mux.HandleFunc("/users/read", usersHandler.UserRead)
    mux.HandleFunc("/users/update", usersHandler.UserUpdate)
    mux.HandleFunc("/users/delete", usersHandler.UserDelete)

    http.ListenAndServe(":8080", mux)
}
