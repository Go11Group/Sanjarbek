package handler

import (
	"encoding/json"
	"fmt"
	"module/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateProblem(w http.ResponseWriter, r *http.Request) {
	var problem model.Problems
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid input: %v"}`, err), http.StatusBadRequest)
		return
	}

	resp, err := h.Problem.Create(problem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not create problem: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	problemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	problem, err := h.Problem.GetById(problemID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Problem not found: %s"}`, params["id"]), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(problem)
}

func (h *Handler) GetAllProblems(w http.ResponseWriter, r *http.Request) {
	resp, err := h.Problem.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Internal server error: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateProblem(w http.ResponseWriter, r *http.Request) {
	var problem model.Problems
	err := json.NewDecoder(r.Body).Decode(&problem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid input: %v"}`, err), http.StatusBadRequest)
		return
	}

	resp, err := h.Problem.Update(problem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not update problem: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DeleteProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	problemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	err = h.Problem.Delete(problemID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not delete problem: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
