package paymentmanager

import (
	"be-assignment/domain"
	"be-assignment/dto"
	"context"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
func (s *service) Transfer(ctx context.Context, transactionReq dto.TransferRequest) error {
	sourceAccount, err := s.accountRepo.FindByAccountNumber(transactionReq.SofNumber)
	if err != nil {
		return domain.ErrSourceAccountNotFound
	}

	userID := ctx.Value("x-user").(jwt.MapClaims)["id"].(string)
	if sourceAccount.UserID != userID {
		return domain.ErrUnauthorizedAccount
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
	transaction.Currency = "USD"
	transaction.TransactionDateTime = time.Now()

	err = s.transactionRepo.Insert(transaction)
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(30 * time.Second)

		transaction.Status = "success"

		sourceAccount.Balance -= transactionReq.Amount
		if err = s.accountRepo.UpdateBalanceWithPessimisticLock(sourceAccount); err != nil {
			transaction.Status = "failed"
			log.Printf("error updating source account: %v", err)
		}

		destinationAccount.Balance += transactionReq.Amount
		if err = s.accountRepo.UpdateBalanceWithPessimisticLock(destinationAccount); err != nil {
			transaction.Status = "failed"
			log.Printf("error updating destination account: %v", err)
		}

		var sourcePaymentHistory domain.PaymentHistory
		sourcePaymentHistory.ID = uuid.New().String()
		sourcePaymentHistory.PaymentAccountID = sourceAccount.ID
		sourcePaymentHistory.TransactionID = transaction.ID
		sourcePaymentHistory.Amount = transactionReq.Amount
		sourcePaymentHistory.TransactionType = "debit"

		if err = s.historyRepo.Insert(sourcePaymentHistory); err != nil {
			log.Printf("error inserting payment history: %v", err)
		}

		var destinationPaymentHistory domain.PaymentHistory
		destinationPaymentHistory.ID = uuid.New().String()
		destinationPaymentHistory.PaymentAccountID = destinationAccount.ID
		destinationPaymentHistory.TransactionID = transaction.ID
		destinationPaymentHistory.Amount = transactionReq.Amount
		destinationPaymentHistory.TransactionType = "credit"

		if err = s.historyRepo.Insert(destinationPaymentHistory); err != nil {
			log.Printf("error inserting payment history: %v", err)
		}

		if err := s.transactionRepo.Update(transaction); err != nil {
			log.Printf("error updating transaction status: %v", err)
		}
	}()

	return nil

}

// Withdraw implements domain.PaymentManagerService.
func (s *service) Withdraw(ctx context.Context, transactionReq dto.WithdrawRequest) error {
	sourceAccount, err := s.accountRepo.FindByAccountNumber(transactionReq.SofNumber)
	if err != nil {
		return domain.ErrSourceAccountNotFound
	}

	userID := ctx.Value("x-user").(jwt.MapClaims)["id"].(string)
	if sourceAccount.UserID != userID {
		return domain.ErrUnauthorizedAccount
	}

	if sourceAccount.Balance < transactionReq.Amount {
		return domain.ErrInsufficientBalance
	}

	var transaction domain.Transaction
	transaction.ID = uuid.New().String()
	transaction.Name = "withdraw"
	transaction.SofNumber = transactionReq.SofNumber
	transaction.DofNumber = "-"
	transaction.Amount = transactionReq.Amount
	transaction.Currency = "USD"
	transaction.TransactionDateTime = time.Now()

	err = s.transactionRepo.Insert(transaction)
	if err != nil {
		return err
	}

	go func() {
		time.Sleep(30 * time.Second)

		transaction.Status = "success"

		sourceAccount.Balance -= transactionReq.Amount
		err = s.accountRepo.UpdateBalanceWithPessimisticLock(sourceAccount)
		if err != nil {
			transaction.Status = "failed"
			log.Printf("error updating source account: %v", err)
		}

		var sourcePaymentHistory domain.PaymentHistory
		sourcePaymentHistory.ID = uuid.New().String()
		sourcePaymentHistory.PaymentAccountID = sourceAccount.ID
		sourcePaymentHistory.TransactionID = transaction.ID
		sourcePaymentHistory.Amount = transactionReq.Amount
		sourcePaymentHistory.TransactionType = "debit"

		err = s.historyRepo.Insert(sourcePaymentHistory)
		if err != nil {
			log.Printf("error inserting payment history: %v", err)
		}

		err := s.transactionRepo.Update(transaction)
		if err != nil {
			log.Printf("error updating transaction status: %v", err)
		}

		if transaction.Status == "success" {
			log.Printf("transaction %s completed", transaction.ID)
		} else {
			log.Printf("transaction %s failed", transaction.ID)
		}
	}()

	return nil
}
