package service

type AuthService interface {
	Register(email, password string) error
	Login(email, password string) (string, error) // JWT
}
