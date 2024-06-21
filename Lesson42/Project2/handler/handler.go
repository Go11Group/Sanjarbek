package handler

import (
	"module/storage/postgres"

	"github.com/gin-gonic/gin"
)

// repositorylarni oz ichiga olgan handler structi
type Handler struct {
	User       postgres.UserRepo
	Course     postgres.CourseRepo
	Lesson     postgres.LessonRepo
	Enrollment postgres.EnrollmentRepo
}

// turli hil api larni guruhlab ishga tushurich
func Router(h *Handler) *gin.Engine {
	r := gin.Default()

	// user uchun userlar
	users := r.Group("/users")
	users.POST("/create", h.CreateUser)
	users.GET("/get/:id", h.GetUserByID)
	users.GET("/", h.GetAllUsers)
	users.GET("/get", h.GetAllUsersFiltered)
	users.PUT("/update/:id", h.UpdateUser)
	users.DELETE("/delete/:id", h.DeleteUser)
	users.GET("/GetNameEmail", h.SearchUsers)

	// course uchun userlar
	courses := r.Group("/courses")
	courses.POST("/create", h.CreateCourse)
	courses.GET("/get/:id", h.GetCourseByID)
	courses.GET("/", h.GetAllCourses)
	courses.GET("/get", h.GetAllCoursesFiltered)
	courses.PUT("/update/:id", h.UpdateCourse)
	courses.DELETE("/delete/:id", h.DeleteCourse)
	courses.GET("/getPopularCourses", h.GetMostPopularCourses)

	// lesson uchun userlar
	lessons := r.Group("/lessons")
	lessons.POST("/create", h.CreateLesson)
	lessons.GET("/get/:id", h.GetLessonByID)
	lessons.GET("/", h.GetAllLessons)
	lessons.GET("/get", h.GetAllLessonsFiltered)
	lessons.PUT("/update/:id", h.UpdateLesson)
	lessons.DELETE("/delete/:id", h.DeleteLesson)
	lessons.GET("/getCourseId/:course_id", h.GetLessonsByCourseID)

	// enrollment uchun userlar
	enrollments := r.Group("/enrollments")
	enrollments.POST("/create", h.CreateEnrollment)
	enrollments.GET("/get/:id", h.GetEnrollmentByID)
	enrollments.GET("/", h.GetAllEnrollments)
	enrollments.GET("/get", h.GetAllEnrollmentsFiltered)
	enrollments.PUT("/update/:id", h.UpdateEnrollment)
	enrollments.DELETE("/delete/:id", h.DeleteEnrollment)
	enrollments.GET("/getByCourseId/:course_id", h.GetEnrolledUsersByCourseID)

	// serverni 8080 portida ishga tushurish
	r.Run(":8080")

	return r
}
