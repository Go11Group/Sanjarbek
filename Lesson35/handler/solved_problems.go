package handler

import (
	"encoding/json"
	"fmt"
	"module/model"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	var solvedProblem model.SolvedProblems
	err := json.NewDecoder(r.Body).Decode(&solvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid input: %v"}`, err), http.StatusBadRequest)
		return
	}

	resp, err := h.SolvedProblem.Create(solvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not create solved problem: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) GetSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	solvedProblemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid solved problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	solvedProblem, err := h.SolvedProblem.GetById(solvedProblemID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Solved problem not found: %s"}`, params["id"]), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(solvedProblem)
}

func (h *Handler) GetAllSolvedProblems(w http.ResponseWriter, r *http.Request) {
	resp, err := h.SolvedProblem.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Internal server error: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) UpdateSolvedProblem(w http.ResponseWriter, r *http.Request) {
	var solvedProblem model.SolvedProblems
	err := json.NewDecoder(r.Body).Decode(&solvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid input: %v"}`, err), http.StatusBadRequest)
		return
	}

	resp, err := h.SolvedProblem.Update(solvedProblem)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not update solved problem: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) DeleteSolvedProblem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	solvedProblemID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Invalid solved problem ID: %s"}`, params["id"]), http.StatusBadRequest)
		return
	}

	err = h.SolvedProblem.Delete(solvedProblemID)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "Could not delete solved problem: %v"}`, err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
