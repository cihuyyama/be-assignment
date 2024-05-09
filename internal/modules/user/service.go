package user

import (
	"be-assignment/domain"

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
func (s *service) GetUserByID(id string) (domain.User, error) {
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
func (s *service) Register(user domain.User) error {
	_, err := s.repo.FindByID(user.ID)
	if err == nil {
		return domain.ErrUserAlreadyExists
	}

	user.ID = uuid.New().String()

	if err := s.repo.Insert(user); err != nil {
		return err
	}

	return nil
}
