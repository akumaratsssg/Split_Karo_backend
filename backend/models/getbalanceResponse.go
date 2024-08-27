package models

type getBalancesResponse struct {
	UserName string  `json:"user_name"`
	Balance  float64 `json:"balance"`
}
