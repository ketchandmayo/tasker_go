package service

import (
	"context"
	"errors"
	"tasker_go/internal/auth"
	"tasker_go/internal/models"
	"tasker_go/internal/repository"
)

type authService struct {
	users repository.UserRepository
}

func NewAuthService(users repository.UserRepository) AuthService {
	return &authService{users: users}
}

func (s *authService) Register(ctx context.Context, email, password string) error {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Password: hash,
	}

	return s.users.Create(ctx, user)
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.users.FindByEmail(ctx, email)
	if err != nil || !auth.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return auth.GenerateJWT(user.ID)
}
