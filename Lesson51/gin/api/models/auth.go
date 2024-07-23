package models

type AuthReq struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginReq struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginRes struct {
	User   UserRes `json:"user_data"`
	Tokens Tokens  `json:"tokens"`
}

