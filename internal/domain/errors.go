package domain

import "errors"

var (
	ErrNotFound      = errors.New("not found")
	ErrForbidden     = errors.New("forbidden")
	ErrAlreadyExists = errors.New("already exist")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrInvalidToken  = errors.New("invalid token")
	ErrValidation  = errors.New("validation error")
)
