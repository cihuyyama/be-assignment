package domain

import "errors"

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

var (
	ErrInvalidToken = errors.New("invalid token")
)
