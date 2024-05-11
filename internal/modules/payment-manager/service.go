package paymentmanager

import (
	"be-assignment/domain"
	"be-assignment/dto"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type service struct {
	transactionRepo domain.TransactionRepository
	accountRepo     domain.AccountRepository
	historyRepo     domain.PaymentHistoryRepository
}

func NewService(tr domain.TransactionRepository, ar domain.AccountRepository, hr domain.PaymentHistoryRepository) domain.PaymentManagerService {
	return &service{
		transactionRepo: tr,
		accountRepo:     ar,
		historyRepo:     hr,
	}
}

// GetAllTransaction implements domain.PaymentManagerService.
func (s *service) GetAllTransaction() ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	transactions, err := s.transactionRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// Transfer implements domain.PaymentManagerService.
func (s *service) Transfer(transactionReq dto.TransferRequest) error {
	sourceAccount, err := s.accountRepo.FindByAccountNumber(transactionReq.SofNumber)
	if err != nil {
		return domain.ErrSourceAccountNotFound
	}
	destinationAccount, err := s.accountRepo.FindByAccountNumber(transactionReq.DofNumber)
	if err != nil {
		return domain.ErrDestinationAccountNotFound
	}

	if sourceAccount.Balance < transactionReq.Amount {
		return domain.ErrInsufficientBalance
	}

	var transaction domain.Transaction
	transaction.ID = uuid.New().String()
	transaction.Name = "transfer"
	transaction.SofNumber = transactionReq.SofNumber
	transaction.DofNumber = transactionReq.DofNumber
	transaction.Amount = transactionReq.Amount

	err = s.transactionRepo.Insert(transaction)
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(30 * time.Second)

		sourceAccount.Balance -= transactionReq.Amount
		err = s.accountRepo.Update(sourceAccount)
		if err != nil {
			transaction.Status = "failed"
			fmt.Printf("error updating source account: %v", err)
		}

		destinationAccount.Balance += transactionReq.Amount
		err = s.accountRepo.Update(destinationAccount)
		if err != nil {
			transaction.Status = "failed"
			fmt.Printf("error updating destination account: %v", err)
		}

		var sourcePaymentHistory domain.PaymentHistory
		sourcePaymentHistory.ID = uuid.New().String()
		sourcePaymentHistory.PaymentAccountID = sourceAccount.ID
		sourcePaymentHistory.TransactionID = transaction.ID
		sourcePaymentHistory.Amount = transactionReq.Amount
		sourcePaymentHistory.TransactionType = "debit"

		err = s.historyRepo.Insert(sourcePaymentHistory)
		if err != nil {
			transaction.Status = "failed"
			fmt.Printf("error inserting payment history: %v", err)
		}

		var destinationPaymentHistory domain.PaymentHistory
		destinationPaymentHistory.ID = uuid.New().String()
		destinationPaymentHistory.PaymentAccountID = destinationAccount.ID
		destinationPaymentHistory.TransactionID = transaction.ID
		destinationPaymentHistory.Amount = transactionReq.Amount
		destinationPaymentHistory.TransactionType = "credit"

		err = s.historyRepo.Insert(destinationPaymentHistory)
		if err != nil {
			transaction.Status = "failed"
			fmt.Printf("error inserting payment history: %v", err)
		}

		transaction.Status = "success"
		err := s.transactionRepo.Update(transaction)
		if err != nil {
			fmt.Printf("error updating transaction status: %v", err)
		}

	}()

	return nil

}

// Withdraw implements domain.PaymentManagerService.
func (s *service) Withdraw(transaction domain.Transaction) error {
	panic("unimplemented")
}
