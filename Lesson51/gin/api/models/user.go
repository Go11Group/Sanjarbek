package models

type UserRes struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}
