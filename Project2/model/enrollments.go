package model

type Enrollments struct {
	EnrollmentId string
	UserId string
	CourseId string
	EnrollmentDate string
}

type EnrollmentGetAll struct {
	UserId string
	CourseId string
	EnrollmentDate string
	Offset int
	Limit int
}