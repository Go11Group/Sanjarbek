package models

type Balance struct {
	UserId string `json:"id"`
	Amound float64 `json:"amount"`
}

type BalanceResp struct {
	Balance      float64 `json:"balance"`
	BalanceStatus string  `json:"balance_status"`
}