package service

import (
	"context"
	"errors"
	"tasker_go/internal/auth"
	"tasker_go/internal/repository"
	"tasker_go/internal/transport/http/dto"
)

type authService struct {
	users repository.UserRepository
}

func NewAuthService(users repository.UserRepository) AuthService {
	return &authService{users: users}
}

func (s *authService) Register(ctx context.Context, req *dto.CreateUserRequest) error {
	var err error
	if req.Password, err = auth.HashPassword(req.Password); err != nil {
		return err
	}

	user := req.ToModel()
	return s.users.Create(ctx, &user)
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.users.FindByEmail(ctx, email)
	if err != nil || !auth.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return auth.GenerateJWT(user.ID)
}
