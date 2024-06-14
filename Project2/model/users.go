package model

type Users struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Birthday string `json:"birthday"`
	Password string `json:"password"`
}

type UserGetAll struct {
	Name     string `json:"name"`
	Birthday string `json:"birthday"`
	Offset   int    `json:"offset"` 
	Limit    int    `json:"limit"` 
}

type UsersGet struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	AgeTo   int    `json:"age_to"`
	AgeFrom int    `json:"age_from"`
}

type AdditionalUser struct {
	ID    string `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Result struct {
	Results []AdditionalUser `json:"results"`
}


type CourseDetails struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type GetCoursesbyUsers struct {
	Id     string        `json:"id"`
	Course CourseDetails `json:"course"`
}

type LessonDetails struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type GetLessonsbyCourses struct {
	Id     string       `json:"id"`
	Lesson LessonDetails `json:"lesson"`
}