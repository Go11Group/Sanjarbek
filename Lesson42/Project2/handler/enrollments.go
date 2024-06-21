package handler

import (
	"fmt"
	"module/Project2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// enrollment yaratish
func (h *Handler) CreateEnrollment(c *gin.Context) {
	var enrollment model.Enrollments

	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Enrollment.CreateEnrollment(enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create enrollment: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// enrollment id boyicha olish
func (h *Handler) GetEnrollmentByID(c *gin.Context) {
	enrollmentID := c.Param("id")

	enrollment, err := h.Enrollment.GetEnrollmentByID(enrollmentID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Enrollment not found: %s", enrollmentID)})
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

// enrollment filterlab olish
func (h *Handler) GetAllEnrollmentsFiltered(c *gin.Context) {
	var filter model.EnrollmentGetAll

	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	enrollments, err := h.Enrollment.GetAllEnrollmentsFiltered(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, enrollments)
}

// enrollment yangilash
func (h *Handler) UpdateEnrollment(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid enrollment ID"})
		return
	}

	var enrollment model.EnrollmentsUpdate
	enrollment.EnrolmentId = id
	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	err := h.Enrollment.EmrolUpdate(enrollment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update enrollment: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"massage": "updated"})
}

// enrollmentni ochirish
func (h *Handler) DeleteEnrollment(c *gin.Context) {
	enrollmentID := c.Param("id")

	err := h.Enrollment.DeleteEnrollment(enrollmentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete enrollment: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment deleted successfully"})
}

func (h *Handler) GetEnrolledUsersByCourseID(c *gin.Context) {
	courseID := c.Param("course_id")
	if courseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	users, err := h.Enrollment.GetEnrolledUsersByCourseID(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not get enrolled users: %v", err)})
		return
	}

	c.JSON(http.StatusOK, users)
}
