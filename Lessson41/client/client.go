package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	CreateUser()
}

type UserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateUser() {
	var name string
	var age int
	fmt.Print("Enter your name : ")
	fmt.Scan(&name)
	fmt.Print("Enter your age : ")
	fmt.Scan(&age)

	user := UserReq{
		Name: name,
		Age:  age,
	}

	jsonUser, err := json.Marshal(user)
	if err != nil {
		log.Fatal("Error encoding JSON : ", err)
	}

	res, err := http.Post("http://localhost:8080/user", "application/json", bytes.NewBuffer(jsonUser))
	if err != nil {
		log.Fatal("Error while making POST request : ", err)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		fmt.Println("User crated succesfully")
	} else {
		fmt.Println("Failed to crate user. Status code : ", res.StatusCode)
	}
}
