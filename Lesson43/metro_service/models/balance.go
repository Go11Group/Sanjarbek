package models

type Balance struct {
	UserID string  `json:"id"`
	Amount float64 `json:"amount"`
}

type BalanceResponse struct {
	Balance       float64 `json:"balance"`
	BalanceStatus string  `json:"balance_status"`
}
