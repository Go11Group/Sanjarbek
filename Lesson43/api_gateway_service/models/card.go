package models

type Card struct {
	Id     string `json:"card_id"`
	UserId string `json:"user_id"`
	Number string `json:"number"`
}