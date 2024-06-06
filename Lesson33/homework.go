package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strconv"

// 	"github.com/gorilla/mux"
// )

// type Car struct {
// 	ID    int    `json:"id"`
// 	Make  string `json:"make"`
// 	Model string `json:"model"`
// 	Price int    `json:"price"`
// }

// var cars = []Car{
// 	{ID: 1, Make: "Toyota", Model: "Corolla", Price: 20000},
// 	{ID: 2, Make: "Honda", Model: "Civic", Price: 22000},
// }

// func main() {
// 	r := mux.NewRouter()

// 	r.HandleFunc("/cars", getCars).Methods("GET")
// 	r.HandleFunc("/cars/{id}", getCar).Methods("GET")
// 	r.HandleFunc("/cars", addCar).Methods("POST")
// 	r.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")
// 	r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")

// 	http.ListenAndServe(":8080", r)
// }

// func getCars(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(cars)
// }

// func getCar(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid car ID", http.StatusBadRequest)
// 		return
// 	}

// 	for _, car := range cars {
// 		if car.ID == id {
// 			json.NewEncoder(w).Encode(car)
// 			return
// 		}
// 	}
// 	http.Error(w, "Car not found", http.StatusNotFound)
// }

// func addCar(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var car Car
// 	err := json.NewDecoder(r.Body).Decode(&car)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	cars = append(cars, car)
// 	json.NewEncoder(w).Encode(car)
// }

// func deleteCar(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid car ID", http.StatusBadRequest)
// 		return
// 	}

// 	for index, car := range cars {
// 		if car.ID == id {
// 			cars = append(cars[:index], cars[index+1:]...)
// 			json.NewEncoder(w).Encode(cars)
// 			return
// 		}
// 	}
// 	http.Error(w, "Car not found", http.StatusNotFound)
// }

// func updateCar(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid car ID", http.StatusBadRequest)
// 		return
// 	}

// 	var updatedCar Car
// 	err = json.NewDecoder(r.Body).Decode(&updatedCar)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	for index, car := range cars {
// 		if car.ID == id {
// 			cars[index] = updatedCar
// 			json.NewEncoder(w).Encode(updatedCar)
// 			return
// 		}
// 	}
// 	http.Error(w, "Car not found", http.StatusNotFound)
// }
