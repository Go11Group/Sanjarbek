package model

type Courses struct {
	CourseId    string `json:"course_id"`
	Title       string
	Description string
	CreatedAt   string
}

type CourseGetAll struct {
	Title string
	Description string
	Offset int
	Limit int
}