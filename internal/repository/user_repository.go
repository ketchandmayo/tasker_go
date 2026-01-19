package repository

import "tasker_go/internal/models"

type UserRepository interface {
	FindByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}
