package service

import (
	"errors"
	"tasker_go/internal/repository"
)

type authService struct {
	users repository.UserRepository
}

func NewAuthService(users repository.UserRepository) AuthService {
	return &authService{users: users}
}

func (s *authService) Login(email, password string) (string, error) {
	user, err := s.users.FindByEmail(email)
	if err != nil || !CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return GenerateJWT(user.ID)
}

func (s *authService) Register(email, password string) error {
	//TODO implement me
	panic("implement me")
}

func CheckPassword(password string, password2 string) bool {
	//TODO implement me
	panic("implement me")
}

func GenerateJWT(id uint) (string, error) {
	//TODO implement me
	panic("implement me")
}
