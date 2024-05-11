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

	var transaction domain.Transaction
	transaction.ID = uuid.New().String()
	transaction.Name = "transfer"
	transaction.SofNumber = transactionReq.SofNumber
	transaction.DofNumber = transactionReq.DofNumber
	transaction.Amount = transactionReq.Amount
	transaction.Status = "pending"
	transaction.Currency = "USD"
	transaction.TransactionDateTime = time.Now()

	if err := s.transactionRepo.Insert(transaction); err != nil {
		return err
	}

	userID := ctx.Value("x-user").(jwt.MapClaims)["id"].(string)

	stop := make(chan bool)

	go func() {
		for {
			select {

			case <-stop:
				return
			default:
				time.Sleep(10 * time.Second)

				transaction.Status = "success"

				sourceAccount, err := s.accountRepo.FindByAccountNumber(transactionReq.SofNumber)
				if err != nil {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					stop <- true
				}

				if sourceAccount.UserID != userID {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					stop <- true
				}

				destinationAccount, err := s.accountRepo.FindByAccountNumber(transactionReq.DofNumber)
				if err != nil {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					stop <- true

				}

				if sourceAccount.Balance < transactionReq.Amount {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					stop <- true
				}

				sourceAccount.Balance -= transactionReq.Amount
				if err = s.accountRepo.UpdateBalanceWithPessimisticLock(sourceAccount); err != nil {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					log.Printf("error updating source account: %v", err)
					stop <- true
				}

				destinationAccount.Balance += transactionReq.Amount
				if err = s.accountRepo.UpdateBalanceWithPessimisticLock(destinationAccount); err != nil {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					log.Printf("error updating destination account: %v", err)
					stop <- true
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
			}
		}

	}()

	return nil

}

// Withdraw implements domain.PaymentManagerService.
func (s *service) Withdraw(ctx context.Context, transactionReq dto.WithdrawRequest) error {

	var transaction domain.Transaction
	transaction.ID = uuid.New().String()
	transaction.Name = "withdraw"
	transaction.SofNumber = transactionReq.SofNumber
	transaction.DofNumber = "-"
	transaction.Amount = transactionReq.Amount
	transaction.Currency = "USD"
	transaction.TransactionDateTime = time.Now()

	err := s.transactionRepo.Insert(transaction)
	if err != nil {
		return err
	}

	userID := ctx.Value("x-user").(jwt.MapClaims)["id"].(string)

	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				return
			default:
				time.Sleep(10 * time.Second)

				transaction.Status = "success"

				sourceAccount, err := s.accountRepo.FindByAccountNumber(transactionReq.SofNumber)
				if err != nil {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					stop <- true
				}

				if sourceAccount.UserID != userID {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					stop <- true
				}

				if sourceAccount.Balance < transactionReq.Amount {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					stop <- true
				}

				sourceAccount.Balance -= transactionReq.Amount
				err = s.accountRepo.UpdateBalanceWithPessimisticLock(sourceAccount)
				if err != nil {
					transaction.Status = "failed"
					if err := s.transactionRepo.Update(transaction); err != nil {
						log.Printf("error updating transaction status: %v", err)
					}
					log.Printf("error updating source account: %v", err)
					stop <- true
				}

				var sourcePaymentHistory domain.PaymentHistory
				sourcePaymentHistory.ID = uuid.New().String()
				sourcePaymentHistory.PaymentAccountID = sourceAccount.ID
				sourcePaymentHistory.TransactionID = transaction.ID
				sourcePaymentHistory.Amount = transactionReq.Amount
				sourcePaymentHistory.TransactionType = "debit"

				if err := s.historyRepo.Insert(sourcePaymentHistory); err != nil {
					log.Printf("error inserting payment history: %v", err)
				}

				if err := s.transactionRepo.Update(transaction); err != nil {
					log.Printf("error updating transaction status: %v", err)
				}
			}
		}
	}()

	return nil
}
