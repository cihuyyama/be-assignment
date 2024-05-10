package domain

import (
	"be-assignment/dto"
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type UserRepository interface {
	FindAll() ([]User, error)
	FindByID(id string) (User, error)
	FindByEmail(email string) (User, error)
	Insert(user User) error
	Update(user User) error
	Delete(id string) error
}

type UserService interface {
	GetAllUsers() ([]User, error)
	GetUser(ctx context.Context) (*User, error)
	Register(userReq dto.RegisterRequest) error
	Login(userReq dto.LoginRequest) (dto.LoginResponse, error)
}
