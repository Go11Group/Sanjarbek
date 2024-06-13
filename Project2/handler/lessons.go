package handler

import (
	"fmt"
	"module/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Createlesson(c *gin.Context) {
	var Lesson model.Lessons
	
	if err := c.ShouldBindJSON(&Lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Lesson.CreateLesson(Lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not create user: %v", err)})
		return
	}

	c.JSON(http.StatusCreated, resp)
}


func (h *Handler) Getlesson(c *gin.Context) {
	LessonID := c.Query("id")

	lesson, err := h.Lesson.GetLessonByID(LessonID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", LessonID)})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) GetlessonId(c *gin.Context) {
	LessonID := c.Param("id")

	lesson, err := h.Lesson.GetLessonByID(LessonID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User not found: %s", LessonID)})
		return
	}

	c.JSON(http.StatusOK, lesson)
}

func (h *Handler) GetAlllessons(c *gin.Context) {
	lessons, err := h.Lesson.GetAllLessons()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Internal server error: %v", err)})
		return
	}

	c.JSON(http.StatusOK, lessons)
}

func (h *Handler) Updatelesson(c *gin.Context) {
	var lesson model.Lessons
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %v", err)})
		return
	}

	resp, err := h.Lesson.UpdateLesson(lesson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not update user: %v", err)})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Deletelesson(c *gin.Context) {
	lessonID := c.Param("id")

	err := h.Lesson.DeleteLesson(lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not delete user: %v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Massage":"Succesfully deleted"})

	c.Status(http.StatusOK)
}
