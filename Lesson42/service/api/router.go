package api

import (
	"module/service/api/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateServer(handler *handler.Handler) *http.Server {
	router := gin.Default()

	usersRoute := router.Group("/users")

	usersRoute.GET("/", handler.GetUsers)
	usersRoute.GET("/:id", handler.GetUsers)
	usersRoute.GET("/:id/courses", handler.GetUsers)
	usersRoute.POST("/create", handler.CreateUsers)
	usersRoute.PUT("/:id/update", handler.UpdateUsers)
	usersRoute.DELETE("/:id/delete", handler.DeleteUser)

	coursesRoute := router.Group("/courses")

	coursesRoute.GET("/all", handler.GetCourses)
	coursesRoute.GET("/:id", handler.GetCourses)
	coursesRoute.GET("/popular", handler.GetCourses)
	coursesRoute.GET("/:id/students", handler.GetCourses)
	coursesRoute.POST("/create", handler.CreateCourse)
	coursesRoute.PUT("/:id/update", handler.UpdateCourse)
	coursesRoute.DELETE("/:id/delete", handler.DeleteCourse)

	enrollmentsRoute := router.Group("/enrollments")

	enrollmentsRoute.GET("/all", handler.GetEnrollment)
	enrollmentsRoute.GET("/:id", handler.GetEnrollment)
	enrollmentsRoute.POST("/create", handler.CreateEnrollment)
	enrollmentsRoute.PUT("/:id/update", handler.UpdateEnrollment)
	enrollmentsRoute.DELETE("/:id/delete", handler.DeleteEnrollment)

	lessonsRoute := router.Group("/lessons")

	lessonsRoute.GET("/all", handler.GetLessons)
	lessonsRoute.GET("/:id", handler.GetLessons)
	lessonsRoute.POST("/create", handler.CreateLesson)
	lessonsRoute.PUT("/:id/update", handler.UpdateLesson)
	lessonsRoute.DELETE("/:id/delete", handler.DeleteLesson)

	return &http.Server{Addr: ":8081", Handler: router}

}
