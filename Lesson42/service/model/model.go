package model

import "time"

type User struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

type Course struct {
	CourseId    string `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Lesson struct {
	LessonId string `json:"lesson_id"`
	CourseId string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type Enrollment struct {
	EnrollmentId   string    `json:"enrollment_id"`
	UserId         string    `json:"user_id"`
	CourseId       string    `json:"course_id"`
	EnrollmentDate time.Time `json:"enrollment_date"`
}

type UserFilter struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type CourseFilter struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Limit       int    `json:"limit"`
	Offset      int    `json:"offset"`
}

type LessonFilter struct {
	CourseId string `json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Limit    int    `json:"limit"`
	Offset   int    `json:"offset"`
}

type EnrollmentFilter struct {
	UserId         string `json:"user_id"`
	CourseId       string `json:"course_id"`
	EnrollmentDate string `json:"enrollment_date"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
}

type UserCourses struct {
	UserId  string       `json:"user_id"`
	Courses []ApiCourses `json:"courses"`
}

type ApiCourses struct {
	CourseId    string `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CourseLesons struct {
	CourseId string      `json:"course_id"`
	Lessons  []ApiLesson `json:"lessons"`
}

type ApiLesson struct {
	LessonId string `json:"lesson_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type CourseUsers struct {
	CourseId     string     `json:"course_id"`
	EnrolledUser []ApiUsers `json:"enrolled_user"`
}

type ApiUsers struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserSearch struct {
	UserId  string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	AgeFrom int    `json:"age_from"`
	AgeTo   int    `json:"age_to"`
}