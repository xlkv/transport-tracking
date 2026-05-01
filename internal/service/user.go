package service

import (
	"context"
	"errors"
	"fmt"

	"tracking.xlkv.com/internal/domain"
	"tracking.xlkv.com/internal/hash"
)

type UserRepository interface {
	Create(ctx context.Context, driver *domain.Driver) error
	GetByID(ctx context.Context, id int64) (*domain.Driver, error)
	GetByUsername(ctx context.Context, username string) (*domain.Driver, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, username, name, password string) (*domain.Driver, error) {

	if username == "" {
		return nil, fmt.Errorf("username is empty")
	}

	if len(username) > 16 {
		return nil, fmt.Errorf("username is so big")
	}

	if len(password) < 8 {
		return nil, fmt.Errorf("passwrod should be longer than 8 characters")
	}

	if len(password) > 16 {
		return nil, fmt.Errorf("password is so big")
	}

	if len(name) < 2 {
		return nil, fmt.Errorf("name should be longer than 3 characters")
	}

	hashedPassword, err := hash.HashPassword(password)

	if err != nil {
		return nil, fmt.Errorf("password is invalid")
	}

	driver := domain.Driver{
		UserName: username,
		Name:     name,
		Password: string(hashedPassword),
	}

	err = s.repo.Create(
		ctx, &driver,
	)

	if err != nil {
		if errors.Is(err, domain.ErrAlreadyExists) {
			return nil, fmt.Errorf("username is already exists")
		}
		return nil, fmt.Errorf("server error")
	}

	return &driver, err
}

func (s *UserService) Login(ctx context.Context, username, password string) (*domain.Driver, error) {

	if len(password) < 3 || len(password) > 16 {
		return nil, fmt.Errorf("password is invalid")
	}

	driver, err := s.repo.GetByUsername(ctx, username)

	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("server error")
	}

	err = hash.ComparePassword(driver.Password, password)

	if err != nil {
		return nil, domain.ErrUnauthorized
	}

	return driver, err
}
