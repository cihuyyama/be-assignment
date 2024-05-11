package domain

import (
	"be-assignment/dto"
	"context"
	"time"
)

type User struct {
	ID             string           `json:"id"`
	Email          string           `json:"email" gorm:"unique"`
	Password       string           `json:"password"`
	PaymentAccount []PaymentAccount `json:"payment_account"`
	CreatedAt      time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
}

type PaymentAccount struct {
	ID             string           `json:"id"`
	UserID         string           `json:"user_id"`
	AccountNumber  string           `json:"account_number" gorm:"unique"`
	Type           string           `json:"type"`
	Balance        int              `json:"balance"`
	PaymentHistory []PaymentHistory `json:"payment_history"`
	CreatedAt      time.Time        `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time        `json:"updated_at" gorm:"autoUpdateTime"`
}

type PaymentHistory struct {
	ID                  string `json:"id"`
	PaymentAccountID    string `json:"payment_account_id"`
	TransactionID       string `json:"transaction_id"`
	Amount              int    `json:"amount"`
	TransactionType     string `json:"transaction_type"`
	TransactionDateTime string `json:"transaction_date_time" gorm:"autoCreateTime"`
}

type UserRepository interface {
	FindByID(id string) (User, error)
	FindByEmail(email string) (User, error)
	Insert(user User) error
	Update(user User) error
	Delete(id string) error
}

type AccountRepository interface {
	FindByID(id string) (PaymentAccount, error)
	FindByUserID(userID string) ([]PaymentAccount, error)
	FindByAccountNumber(accountNumber string) (PaymentAccount, error)
	Insert(account PaymentAccount) error
	Update(account PaymentAccount) error
	Delete(accountNumber string) error
}

type AccountManagerService interface {
	GetUser(ctx context.Context) (*User, error)
	Register(userReq dto.RegisterRequest) error
	Login(userReq dto.LoginRequest) (dto.LoginResponse, error)

	GetAllAccount(ctx context.Context) (*[]PaymentAccount, error)
	CreateAccount(ctx context.Context, accountReq dto.CreateAccountRequest) error
}
