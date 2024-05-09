package user

import "be-assignment/domain"

type service struct {
	repo domain.UserRepository
}

func NewService(repo domain.UserRepository) domain.UserService {
	return &service{
		repo,
	}
}

// CreateUser implements domain.UserService.
func (s *service) Register(user domain.User) error {
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
