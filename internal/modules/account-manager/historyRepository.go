package accountmanager

import (
	"be-assignment/domain"

	"gorm.io/gorm"
)

type historyRepository struct {
	db *gorm.DB
}

func NewHistoryRepository(con *gorm.DB) domain.PaymentHistoryRepository {
	return &historyRepository{
		db: con,
	}
}

// Insert implements domain.PaymentHistoryRepository.
func (h *historyRepository) Insert(paymentHistory domain.PaymentHistory) error {
	tx := h.db.Create(&paymentHistory)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
