package dto

type CreateAccountRequest struct {
	AccountNumber string `json:"account_number" validate:"required"`
	AccountType   string `json:"type" validate:"required"`
	Balance       int    `json:"balance" validate:"required"`
}
