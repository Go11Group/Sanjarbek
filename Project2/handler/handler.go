package handler

import (
	"module/storage/postgres"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	User   postgres.UserRepo
	Course postgres.CourseRepo
	Lesson postgres.LessonRepo
	Enrollment postgres.EnrollmentRepo
}

func NewHandler(userRepo postgres.UserRepo) *Handler {
	return &Handler{User: userRepo}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {

	users := r.Group("/users")
	users.POST("/create", h.CreateUser)
	users.GET("/get", h.GetUser)
	users.GET("/get/:id", h.GetUserId)
	users.GET("/c", h.GetAllUsers)
	users.PUT("/update/:id", h.UpdateUser)
	users.DELETE("/delete/:id", h.DeleteUser)



	courses := r.Group("/courses")
	courses.POST("/create", h.CreateCourse)
	courses.POST("/get", h.GetCourse)
	courses.POST("/get/:id", h.GetCourseId)
	courses.POST("/c", h.GetAllCourses)
	courses.POST("/updata/:id", h.UpdateCourse)
	courses.POST("/delete/:id", h.DeleteCourse)



	lessons := r.Group("/lessons")
	lessons.POST("/create", h.Createlesson)
	lessons.POST("/get", h.Getlesson)
	lessons.POST("/get/:id", h.GetlessonId)
	lessons.POST("/c", h.GetAlllessons)
	lessons.POST("/updata/:id", h.Updatelesson)
	lessons.POST("/delete/:id", h.Deletelesson)



	enrollment := r.Group("/lessons")
	enrollment.POST("/create", h.Createlesson)
	enrollment.POST("/get", h.Getlesson)
	enrollment.POST("/get/:id", h.GetlessonId)
	enrollment.POST("/c", h.GetAlllessons)
	enrollment.POST("/updata/:id", h.Updatelesson)
	enrollment.POST("/delete/:id", h.Deletelesson)

}
