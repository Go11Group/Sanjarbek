package model

type Lessons struct {
	LessonId string `json:"lesson_id"`
	CourseId string	`json:"course_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type LessonGetAll struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Offset  int `json:"offset"`
	Limit   int `json:"limit"`
}

type LessonsUpdate struct {
    LessonId string `json:"lesson_id"`
    Title    string `json:"title"`
    Content  string `json:"content"`
}