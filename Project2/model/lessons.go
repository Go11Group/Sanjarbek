package model

type Lessons struct {
	LessonId string
	CourseId string
	Title string
	Content string
}

type LessonGetAll struct {
	Title string
	Content string
	Offset int
	Limit int
}