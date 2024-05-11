package domain

import "errors"

var (
	ErrUserAlreadyExists          = errors.New("user already exists")
	ErrUserNotFound               = errors.New("user not found")
	ErrInvalidPassword            = errors.New("invalid password")
	ErrInvalidToken               = errors.New("invalid token")
	ErrUnauthorized               = errors.New("unauthorized")
	ErrAccountAlreadyExists       = errors.New("account already exists")
	ErrSourceAccountNotFound      = errors.New("source account not found")
	ErrDestinationAccountNotFound = errors.New("destination account not found")
	ErrInsufficientBalance        = errors.New("insufficient balance")
	ErrUnauthorizedAccount        = errors.New("unauthorized account")
)
