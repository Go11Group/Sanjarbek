package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	s := http.NewServeMux()

	s.HandleFunc("POST /user", CreateUser)
	s.HandleFunc("GET /user", GetUser)

	log.Println("servet is running on :8080 ...")
	if err := http.ListenAndServe(":8080", s); err != nil {
		log.Fatal(err)
	}
}

type UserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := UserReq{}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.Write([]byte("error while reading data from "))
		return
	}

	fmt.Println("User dsata -> ", req)

	w.WriteHeader(201)
	w.Write([]byte("user crated !"))

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	res := UserReq{
		Name: "Test",
		Age: 17,
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		w.Write([]byte("error while reading"))
	}
}
