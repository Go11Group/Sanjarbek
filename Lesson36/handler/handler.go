package handler

import (
	"module/storage/postgres"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User          *postgres.UserRepo
	Problem       *postgres.ProblemRepo
	SolvedProblem *postgres.SolvedProblemRepo
}

func NewHandler(handler Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/user/create", handler.CreateUser)
	r.GET("/user", handler.GetUser)
	r.GET("/user/:id", handler.GetUserId)
	r.GET("/users", handler.GetAllUser)
	r.PUT("/user", handler.UpdateUser)
	r.DELETE("/deleteUser/:id", handler.DeleteUser)

	r.POST("/problem/create", handler.CreateProblem)
	r.GET("/problems", handler.GetAllProblems)
	r.GET("/problem/:id", handler.GetProblem)
	r.PUT("/updateProblem/:id", handler.UpdateProblem)
	r.DELETE("/deleteProblem/:id", handler.DeleteProblem)

	r.POST("/createSolvedProblem", handler.CreateSolvedProblem)
	r.GET("/SolvedProblem/:id", handler.GetSolvedProblem)
	r.GET("/SolvedProblems", handler.GetAllSolvedProblems)
	r.PUT("/updateSolvedProblem/:id", handler.UpdateSolvedProblem)
	r.DELETE("/deleteSolvedProblem/:id", handler.DeleteSolvedProblem)

	return r
}
