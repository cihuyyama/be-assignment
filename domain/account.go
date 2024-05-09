package domain

type Account struct {
	ID            string `json:"id"`
	UserID        string `json:"user_id"`
	AccountNumber string `json:"account_number"`
	Type          string `json:"type"`
	Balance       int    `json:"balance"`
}

type AccountRepository interface {
	FindByID(id string) (Account, error)
	FindByUserID(userID string) ([]Account, error)
	FindByAccountNumber(accountNumber string) (Account, error)
	Insert(account Account) error
	Update(account Account) error
	Delete(accountNumber string) error
}

type AccountService interface {
}
