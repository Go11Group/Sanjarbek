package handler

import (
	"fmt"
	"module/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateEnrollment(c *gin.Context) {
	var enrollment model.Enrollments
	
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Enrollment.CreateEnrollment(enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) GetEnrollmentId(c *gin.Context) {
	enrollmentID := c.Param("id")

	enrollment, err := h.Enrollment.GetEnrollmentByID(enrollmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", enrollmentID)})
		return
	}

	c.JSON(http.StatusOK, enrollment)
}

func (h *Handler) GetAllEnrollments(c *gin.Context) {
	enrollments, err := h.Enrollment.GetAllEnrollments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

func (h *Handler) UpdateEnrollment(c *gin.Context) {
	var enrollment model.Enrollments
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Enrollment.UpdateEnrollment(enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteEnrollment(c *gin.Context) {
	enrollmentID := c.Param("id")

	err := h.Enrollment.DeleteEnrollment(enrollmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete user: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Massage":"Succesfully deleted"})

	c.Status(http.StatusOK)
}
