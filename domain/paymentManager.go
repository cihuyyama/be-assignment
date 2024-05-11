package domain

import (
	"be-assignment/dto"
	"time"
)

type Transaction struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`       // withdraw / transfer
	SofNumber           string    `json:"sof_number"` // Source of Fund Number
	DofNumber           string    `json:"dof_number"` // Destination of Fund Number
	Currency            string    `json:"currency" gorm:"default:'USD'"`
	Amount              int       `json:"amount"`
	Status              string    `json:"status" gorm:"default:'pending'"`
	TransactionDateTime time.Time `json:"transaction_date_time" gorm:"autoCreateTime"`
}

type TransactionRepository interface {
	FindAll() ([]Transaction, error)
	FindByID(id string) (Transaction, error)
	Update(transaction Transaction) error
	Insert(transaction Transaction) error
}

type PaymentManagerService interface {
	GetAllTransaction() ([]Transaction, error)
	Withdraw(transaction Transaction) error
	Transfer(transactionReq dto.TransferRequest) error
}
