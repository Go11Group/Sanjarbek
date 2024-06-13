package model

type Users struct {
	UserId    string
	Name      string
	Email     string
	Birthday  string
	Password  string
	CreatedAt string
	UpdatedAt string
}

type UserGetAll struct{
	Name string
	Birthday string
	Offset int
	Limit int
}
