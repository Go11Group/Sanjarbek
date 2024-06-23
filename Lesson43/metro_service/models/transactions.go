package models

type Transaction struct {
    Id              string
    CardId          string
    Amount          float64
    TerminalId      string
    TransactionType string
}