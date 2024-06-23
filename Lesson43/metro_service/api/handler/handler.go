package handler

import (
	"database/sql"
	"metro_service/storage/postgres"
)

type handler struct {
	Station      *postgres.StationRepo
	Card         *postgres.CardRepo
	Terminal     *postgres.TerminalRepo
	Transactions *postgres.TransactionRepo
}

func NewHandler(db *sql.DB) *handler {
	return &handler{
		Station:      postgres.NewStationRepo(db),
		Card:         postgres.NewCardRepo(db),
		Terminal:     postgres.NewTerminalRepo(db),
		Transactions: postgres.NewTransactionRepo(db),
	}
}
