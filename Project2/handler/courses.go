package handler

import (
	"fmt"
	"module/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreateCourse(c *gin.Context) {
	var Course model.Courses
	
	if err := c.ShouldBindJSON(&Course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Course.CreateCourse(Course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, resp)
}


func (h *Handler) GetCourse(c *gin.Context) {
	courseID := c.Query("id")

	course, err := h.Course.GetCourseByID(courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", courseID)})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *Handler) GetCourseId(c *gin.Context) {
	courseID := c.Param("id")

	course, err := h.Course.GetCourseByID(courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", courseID)})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *Handler) GetAllCourses(c *gin.Context) {
	courses, err := h.Course.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) UpdateCourse(c *gin.Context) {
	var course model.Courses
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Course.UpdateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteCourse(c *gin.Context) {
	courseID := c.Param("id")

	err := h.Course.DeleteCourse(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete user: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Massage":"Succesfully deleted"})

	c.Status(http.StatusOK)
}
