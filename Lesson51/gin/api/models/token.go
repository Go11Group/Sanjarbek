package models

type TokenReq struct {
	UserId   int
	UserName string
	Email    string
}

type Tokens struct {
	AccesToken string `json:"acces_token"`
}
