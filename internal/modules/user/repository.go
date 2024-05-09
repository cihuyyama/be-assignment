package user

import (
	"be-assignment/domain"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(con *gorm.DB) domain.UserRepository {
	return &repository{
		db: con,
	}
}

// FindAll implements domain.UserRepository.
func (r *repository) FindAll() ([]domain.User, error) {
	var users *[]domain.User
	tx := r.db.Find(&users)
	if tx.Error != nil {
		return []domain.User{}, tx.Error
	}
	return *users, nil
}

// Insert implements domain.UserRepository.
func (r *repository) Insert(user domain.User) error {
	tx := r.db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
