package dto

import (
	"strings"
	"tasker_go/internal/models"
)

func (r CreateUserRequest) ToModel() models.User {
	return models.User{
		Email:    strings.TrimSpace(r.Email),
		Password: r.Password,
		Status:   "user",
	}
}
