package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"module/model"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		fmt.Println("OKOKOK")
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	fmt.Println(user)
	resp, err := h.User.Create(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["id"]

	user, err := h.User.GetById(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {

	resp, err := h.User.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	resp, err := h.User.Update(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
	}

	err = json.NewEncoder(w).Encode(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path,"/Users/")

	err := h.User.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		}
	w.WriteHeader(http.StatusOK)
}