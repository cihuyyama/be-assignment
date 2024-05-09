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

// FindByID implements domain.UserRepository.
func (r *repository) FindByID(id string) (domain.User, error) {
	var user domain.User
	tx := r.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}
	return user, nil
}

// FindByEmail implements domain.UserRepository.
func (r *repository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	tx := r.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return domain.User{}, tx.Error
	}
	return user, nil
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

// Update implements domain.UserRepository.
func (r *repository) Update(user domain.User) error {
	tx := r.db.Save(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Delete implements domain.UserRepository.
func (r *repository) Delete(id string) error {
	tx := r.db.Where("id = ?", id).Delete(&domain.User{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
