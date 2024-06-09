package handler

import (
	"module/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	User *postgres.UserRepo
	Problem *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := mux.NewRouter()

	mux.HandleFunc("/createUser", handler.CreateUser).Methods("POST")
	mux.HandleFunc("/users/{id}", handler.GetUser).Methods("GET")
	mux.HandleFunc("/users", handler.GetAllUser).Methods("GET")
	mux.HandleFunc("/updateUser/{id}", handler.UpdateUser).Methods("PUT")
	mux.HandleFunc("/deleteUser/{id}", handler.DeleteUser).Methods("DELETE")


	mux.HandleFunc("/createProblem", handler.CreateProblem).Methods("POST")
	mux.HandleFunc("/problem/{id}", handler.GetProblem).Methods("GET")
	mux.HandleFunc("/problems/", handler.GetAllProblems).Methods("GET")
	mux.HandleFunc("/updateProblem/{id}", handler.UpdateProblem).Methods("PUT")
	mux.HandleFunc("/deleteProblem/{id}", handler.DeleteProblem).Methods("DELETE")

	mux.HandleFunc("/createSolvedProblem", handler.CreateSolvedProblem).Methods("POST")
	mux.HandleFunc("/SolvedProblem/{id}", handler.GetSolvedProblem).Methods("GET")
	mux.HandleFunc("/SolvedProblems/", handler.GetAllSolvedProblems).Methods("GET")
	mux.HandleFunc("/updateSolvedProblem/{id}", handler.UpdateSolvedProblem).Methods("PUT")
	mux.HandleFunc("/deleteSolvedProblem/{id}", handler.DeleteSolvedProblem).Methods("DELETE")


	return &http.Server{Addr: ":8080", Handler: mux}
}
