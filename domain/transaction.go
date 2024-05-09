package domain

type Transaction struct {
	ID              string `json:"id"`
	AccountID       string `json:"account_id"`
	SofNumber       string `json:"sof_number"`
	DofNumber       string `json:"dof_number"`
	TransactionType string `json:"transaction_type"`
	Amount          int    `json:"amount"`
}

type TransactionRepository interface {
}
