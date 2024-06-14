package model

import (
	"time"

)

type Enrollments struct {
	EnrollmentId string `json:"enrollment_id"`
	UserId       string `json:"user_id"`
	CourseId     string `json:"course_id"`
}

type EnrollmentGetAll struct {
	UserId   string `json:"user_id"`
	CourseId string `json:"course_id"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type EnrollmentsUpdate struct {
	EnrolmentId    string  `json:"enrollment_id"`
	UserId         *int       `json:"user_id"`
	CourseId       *int       `json:"course_id"`
	EnrollmentDate *time.Time `json:"enrollment_date"`
}
