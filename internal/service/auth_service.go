package service

import (
	"context"
	"tasker_go/internal/transport/http/dto"
)

type AuthService interface {
	Register(ctx context.Context, req *dto.CreateUserRequest) error
	Login(ctx context.Context, email, password string) (string, error)
}
