package accountmanager

import (
	"be-assignment/domain"

	"gorm.io/gorm"
)

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(con *gorm.DB) domain.AccountRepository {
	return &accountRepository{
		db: con,
	}
}

// FindByAccountNumber implements domain.AccountRepository.
func (a *accountRepository) FindByAccountNumber(accountNumber string) (domain.PaymentAccount, error) {
	var account domain.PaymentAccount
	tx := a.db.Where("account_number = ?", accountNumber).First(&account)
	if tx.Error != nil {
		return domain.PaymentAccount{}, tx.Error
	}
	return account, nil
}

// FindByID implements domain.AccountRepository.
func (a *accountRepository) FindByID(id string) (domain.PaymentAccount, error) {
	var account domain.PaymentAccount
	tx := a.db.Where("id = ?", id).First(&account)
	if tx.Error != nil {
		return domain.PaymentAccount{}, tx.Error
	}
	return account, nil
}

// FindByUserID implements domain.AccountRepository.
func (a *accountRepository) FindByUserID(userID string) ([]domain.PaymentAccount, error) {
	var accounts []domain.PaymentAccount
	tx := a.db.Preload("PaymentHistory").Where("user_id = ?", userID).Find(&accounts)
	if tx.Error != nil {
		return []domain.PaymentAccount{}, tx.Error
	}
	return accounts, nil
}

// Insert implements domain.AccountRepository.
func (a *accountRepository) Insert(account domain.PaymentAccount) error {
	tx := a.db.Create(&account)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Update implements domain.AccountRepository.
func (a *accountRepository) Update(account domain.PaymentAccount) error {
	tx := a.db.Save(&account)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements domain.AccountRepository.
func (a *accountRepository) Delete(accountNumber string) error {
	panic("unimplemented")
}
