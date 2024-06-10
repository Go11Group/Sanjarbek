package handler

import (
	"fmt"
	"module/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProblem(c *gin.Context) {
	var problem model.Problems

	if err := c.ShouldBindJSON(&problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : fmt.Sprint("Invalid input %v", err)})
	}

	resp, err := h.Problem.Create(problem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : fmt.Sprint("Could not crate problem %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProblem(c *gin.Context) {
	problemId := c.Param("id")

	ID, err := strconv.Atoi(problemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : fmt.Sprintf("Invalid ProblemId %s", problemId)})
	}

	problem, err := h.Problem.GetById(ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : fmt.Sprintf("Problem not found %s", problemId)})
	}

	c.JSON(http.StatusOK, problem)

}

func (h *Handler) GetAllProblems(c *gin.Context) {
	resp, err := h.Problem.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateProblem(c *gin.Context) {
	var problem model.Problems
	if err := c.ShouldBindJSON(&problem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : fmt.Sprintf("Invalid input %s", err)})
	}

	resp, err := h.Problem.Update(problem)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : fmt.Sprintf("Could not create user %s", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteProblem(c *gin.Context) {
	problemId := c.Param("id")

	Id, err := strconv.Atoi(problemId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid Problem Id: %s", problemId)})
		return
	}

	err = h.Problem.Delete(Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Problem did not delete: %v", err)})
		return
	}

	c.Status(http.StatusOK)
}

