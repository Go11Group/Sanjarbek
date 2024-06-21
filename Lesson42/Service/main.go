package main

import (
	"log"
	"module/handler"
	"module/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	userRepo := postgres.NewUserRepo(db)
	courseRepo := postgres.NewCourseRepo(db)
	lessonRepo := postgres.NewLessonRepo(db)
	enrollmentRepo := postgres.NewEnrollmentRepo(db)

	h := handler.Handler{User: *userRepo, Course: *courseRepo, Lesson: *lessonRepo, Enrollment: *enrollmentRepo}

	handler.Router(&h)

}
