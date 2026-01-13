package models

import "time"

type Task struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	UserId      uint   `gorm:"default:1"`
	Description string
	Status      string `gorm:"default:'new'"`
	CreatedAt   time.Time
}
