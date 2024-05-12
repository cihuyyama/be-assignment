package paymentmanager

import (
	"be-assignment/domain"

	"gorm.io/gorm"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(con *gorm.DB) domain.TransactionRepository {
	return &transactionRepository{
		db: con,
	}
}

// FindAll implements domain.TransactionRepository.
func (t *transactionRepository) FindAll() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	tx := t.db.Order("transaction_date_time desc").Find(&transactions)
	if tx.Error != nil {
		return []domain.Transaction{}, tx.Error
	}
	return transactions, nil
}

// FindByID implements domain.TransactionRepository.
func (t *transactionRepository) FindByID(id string) (domain.Transaction, error) {
	var transaction domain.Transaction
	tx := t.db.Where("id = ?", id).First(&transaction)
	if tx.Error != nil {
		return domain.Transaction{}, tx.Error
	}
	return transaction, nil
}

// Insert implements domain.TransactionRepository.
func (t *transactionRepository) Insert(transaction *domain.Transaction) error {
	tx := t.db.Create(&transaction)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements domain.TransactionRepository.
func (t *transactionRepository) Update(transaction *domain.Transaction) error {
	tx := t.db.Save(&transaction)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
