package model


type Courses struct {
	CourseId    string `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CourseGetAll struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Offset      int    `json:"offset"`
	Limit       int    `json:"limit"`
}

type CoursePopularity struct {
	CourseId        string `json:"course_id"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	CreatedAt       string `json:"created_at"`
	EnrollmentCount int    `json:"enrollment_count"`
}
