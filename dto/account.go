package dto

type CreateAccountRequest struct {
	AccountNumber string `json:"account_number"`
	Type          string `json:"type"`
	Balance       int    `json:"balance"`
}
