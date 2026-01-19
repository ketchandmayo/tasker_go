package gorm

import (
	"tasker_go/internal/models"
	"tasker_go/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	return nil, nil
}

func (r *userRepository) Create(user *models.User) error {
	return nil
}
