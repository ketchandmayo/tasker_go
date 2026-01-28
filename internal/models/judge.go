package models

import "time"

type Judge struct {
	ID        uint   `gorm:"primaryKey"`
	TaskID    uint   `gorm:"not null"`
	TaskHash  string `gorm:"not null"`
	Score     uint   `gorm:"default:5"`
	Text      string `gorm:"not null"`
	CreatedAt time.Time
}
