package handler

import (
	"encoding/json"
	"fmt"
	"module/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid input: %v"}`, err), http.StatusBadRequest)
		return
	}

	resp, err := h.User.Create(user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not create user: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")

	ID, err := strconv.Atoi(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "User not found: %s"}`, userID), http.StatusNotFound)
		return
	}

	user, err := h.User.GetUserByID(ID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "User not found: %s"}`, userID), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUserId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	user, err := h.User.GetUserByID(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "User not found: %s"}`, params["id"]), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetAllUser(w http.ResponseWriter, r *http.Request) {
	resp, err := h.User.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Internal server error: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid input: %v"}`, err), http.StatusBadRequest)
		return
	}

	resp, err := h.User.Update(user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not update user: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid user ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	err = h.User.Delete(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not delete user: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
