package handler

import (
	"module/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	User          *postgres.UserRepo
	Problem       *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := mux.NewRouter()

	user := mux.PathPrefix("/users").Subrouter()
	// TODO fdsfdsgfdhgcv

	user.HandleFunc("/create", handler.CreateUser).Methods("POST")
	user.HandleFunc("/get/", handler.GetUser).Methods("GET")
	user.HandleFunc("/get/{id}", handler.GetUserId).Methods("GET")
	user.HandleFunc("/getall", handler.GetAllUser).Methods("GET")
	user.HandleFunc("/update/{id}", handler.UpdateUser).Methods("PUT")
	user.HandleFunc("/delete/{id}", handler.DeleteUser).Methods("DELETE")

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
