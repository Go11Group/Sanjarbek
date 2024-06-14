package handler

import (
	"fmt"
	"module/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// lesson yaratish
func (h *Handler) CreateLesson(c *gin.Context) {
	var lesson model.Lessons

	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Lesson.CreateLesson(lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create lesson: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// id boyicha lessonni olish
func (h *Handler) GetLessonByID(c *gin.Context) {
	lessonID := c.Param("id")

	lesson, err := h.Lesson.GetLessonByID(lessonID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Lesson not found: %s", lessonID)})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

// GetAllLessons handles retrieving all lessons with optional filters
func (h *Handler) GetAllLessons(c *gin.Context) {
	title := c.Query("title")
	content := c.Query("content")

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "0"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	filter := model.LessonGetAll{
		Title:   title,
		Content: content,
		Offset:  offset,
		Limit:   limit,
	}

	lessons, err := h.Lesson.GetAllLessons(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

// filterlab lessonlarni chiqarish
func (h *Handler) GetAllLessonsFiltered(c *gin.Context) {
	var filter model.LessonGetAll

	if err := c.BindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	lessons, err := h.Lesson.GetAllLessons(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

// lessonni yangilash
func (h *Handler) UpdateLesson(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid lesson ID"})
		return
	}

	var lesson model.LessonsUpdate
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}
	lesson.LessonId = id
	resp, err := h.Lesson.UpdateLesson(lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update lesson: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// ID boyicha lessonni delete qilish
func (h *Handler) DeleteLesson(c *gin.Context) {
	lessonID := c.Param("id")

	err := h.Lesson.DeleteLesson(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete lesson: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Lesson deleted successfully"})
}

func (h *Handler) GetLessonsByCourseID(c *gin.Context) {
	courseId := c.Param("course_id")

	resp, err := h.Lesson.GetLessonsByCourseID(courseId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not find lesson"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
