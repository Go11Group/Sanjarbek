package handler

import (
	"fmt"
	"module/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Assuming Handler struct and NewHandler function are defined elsewhere

func (h *Handler) CreateSolvedProblem(c *gin.Context) {
	var solvedProblem model.SolvedProblems
	if err := c.ShouldBindJSON(&solvedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.SolvedProblem.Create(solvedProblem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create solved problem: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetSolvedProblem(c *gin.Context) {
	solvedProblemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid solved problem ID: %s", c.Param("id"))})
		return
	}

	solvedProblem, err := h.SolvedProblem.GetById(solvedProblemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Solved problem not found: %s", c.Param("id"))})
		return
	}

	c.JSON(http.StatusOK, solvedProblem)
}

func (h *Handler) GetAllSolvedProblems(c *gin.Context) {
	resp, err := h.SolvedProblem.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateSolvedProblem(c *gin.Context) {
	var solvedProblem model.SolvedProblems
	if err := c.ShouldBindJSON(&solvedProblem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.SolvedProblem.Update(solvedProblem)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update solved problem: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteSolvedProblem(c *gin.Context) {
	solvedProblemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid solved problem ID: %s", c.Param("id"))})
		return
	}

	err = h.SolvedProblem.Delete(solvedProblemID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete solved problem: %v", err)})
		return
	}

	c.Status(http.StatusOK)
}
