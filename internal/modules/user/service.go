package user

import (
	"be-assignment/domain"
	"be-assignment/dto"
	"be-assignment/internal/util"

	"github.com/google/uuid"
)

type service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) domain.UserService {
	return &service{
		repo,
	}
}

// GetUserByID implements domain.UserService.
func (s *service) GetUser(id string) (domain.User, error) {
	panic("unimplemented")
}

// GetAllUsers implements domain.UserService.
func (s *service) GetAllUsers() ([]domain.User, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return []domain.User{}, err
	}
	return users, nil
}

// CreateUser implements domain.UserService.
func (s *service) Register(userReq dto.RegisterRequest) error {
	_, err := s.repo.FindByEmail(userReq.Email)
	if err == nil {
		return domain.ErrUserAlreadyExists
	}

	var user domain.User
	user.ID = uuid.New().String()
	user.Username = userReq.Username
	user.Email = userReq.Email

	hashedPassword, err := util.HashPassword(userReq.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	if err := s.repo.Insert(user); err != nil {
		return err
	}

	return nil
}

// Login implements domain.UserService.
func (s *service) Login(userReq dto.LoginRequest) (dto.LoginResponse, error) {
	userRepo, err := s.repo.FindByEmail(userReq.Email)
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
