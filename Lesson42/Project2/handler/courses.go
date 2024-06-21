package handler

import (
	"fmt"
	"module/Project2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// kurs yaratish
func (h *Handler) CreateCourse(c *gin.Context) {
	var course model.Courses

	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Course.CreateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create course: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// kursni id boyicha topish
func (h *Handler) GetCourseByID(c *gin.Context) {
	courseID := c.Param("id")

	course, err := h.Course.GetCourseByID(courseID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Course not found: %s", courseID)})
		return
	}

	c.JSON(http.StatusOK, course)
}

// hamma kurslarni olish
func (h *Handler) GetAllCourses(c *gin.Context) {
	courses, err := h.Course.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// kusrlarni filtrlab olish
func (h *Handler) GetAllCoursesFiltered(c *gin.Context) {
	var filter model.CourseGetAll

	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	courses, err := h.Course.GetAllCoursesFiltered(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, courses)
}

// kurslarni yangilash
func (h *Handler) UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	var course model.Courses
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	course.CourseId = id
	updatedCourse, err := h.Course.UpdateCourse(course)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCourse)
}

// kursni ochirish
func (h *Handler) DeleteCourse(c *gin.Context) {
	courseID := c.Param("id")

	err := h.Course.DeleteCourse(courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete course: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// 2 vaqt oraligidagi eng mashhur kurslarni olish
func (h *Handler) GetMostPopularCourses(c *gin.Context) {
	var dateRange struct {
		StartDate string `json:"start_date" binding:"required"`
		EndDate   string `json:"end_date" binding:"required"`
	}

	if err := c.ShouldBindJSON(&dateRange); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	courses, err := h.Course.GetMostPopularCourses(dateRange.StartDate, dateRange.EndDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not retrieve popular courses: %v", err)})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) GetCoursesbyUser(c *gin.Context) {
	userID := c.Param("id")

	resp, err := h.Course.GetCoursesbyUser(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "UserId not found"})
	}

	c.JSON(http.StatusOK, resp)
}
