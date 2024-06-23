package models

type Card struct {
	Id, Number, UserId string
}

type CreateCard struct {
	Number string
	UserId string
}