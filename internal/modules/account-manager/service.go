package accountmanager

import (
	"be-assignment/domain"
	"be-assignment/dto"
	"be-assignment/internal/util"
	"context"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type service struct {
	userRepo    domain.UserRepository
	accountRepo domain.AccountRepository
}

func NewService(userRepo domain.UserRepository, accountRepo domain.AccountRepository) domain.AccountManagerService {
	return &service{
		userRepo,
		accountRepo,
	}
}

// GetUser implements domain.UserService.
func (s *service) GetUser(ctx context.Context) (*domain.User, error) {
	claims := ctx.Value("x-user").(jwt.MapClaims)

	user, err := s.userRepo.FindByID(claims["id"].(string))
	if err != nil {
		return &domain.User{}, err
	}
	return &user, nil
}

// CreateUser implements domain.UserService.
func (s *service) Register(userReq dto.RegisterRequest) error {
	_, err := s.userRepo.FindByEmail(userReq.Email)
	if err == nil {
		return domain.ErrUserAlreadyExists
	}

	var user domain.User
	user.ID = uuid.New().String()
	user.Email = userReq.Email

	hashedPassword, err := util.HashPassword(userReq.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	if err := s.userRepo.Insert(user); err != nil {
		return err
	}

	return nil
}

// Login implements domain.UserService.
func (s *service) Login(userReq dto.LoginRequest) (dto.LoginResponse, error) {
	userRepo, err := s.userRepo.FindByEmail(userReq.Email)
	if err != nil {
		return dto.LoginResponse{}, domain.ErrUserNotFound
	}

	if _, err := util.CheckPasswordHash(userReq.Password, userRepo.Password); err != nil {
		return dto.LoginResponse{}, domain.ErrInvalidPassword
	}

	token, err := util.GenerateToken(&userRepo)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	return dto.LoginResponse{
		Token: token,
	}, nil
}

// CreateAccount implements domain.AccountManagerService.
func (s *service) CreateAccount(ctx context.Context, accountReq dto.CreateAccountRequest) error {
	_, err := s.accountRepo.FindByAccountNumber(accountReq.AccountNumber)
	if err == nil {
		return domain.ErrAccountAlreadyExists
	}

	var account domain.PaymentAccount
	account.ID = uuid.New().String()
	account.AccountNumber = accountReq.AccountNumber
	account.AccountType = accountReq.AccountType
	account.Balance = accountReq.Balance

	claims := ctx.Value("x-user").(jwt.MapClaims)
	account.UserID = claims["id"].(string)

	if err := s.accountRepo.Insert(account); err != nil {
		return err
	}

	return nil
}

// GetAccount implements domain.AccountManagerService.
func (s *service) GetAllAccount(ctx context.Context) (*[]domain.PaymentAccount, error) {
	claims := ctx.Value("x-user").(jwt.MapClaims)

	accounts, err := s.accountRepo.FindByUserID(claims["id"].(string))
	if err != nil {
		return &[]domain.PaymentAccount{}, err
	}

	return &accounts, nil
}
