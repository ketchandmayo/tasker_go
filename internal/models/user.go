package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Status    string `gorm:"default:'user'"`
	CreatedAt time.Time
}
