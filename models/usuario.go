package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"size:150;not null"`
	Email        string    `gorm:"size:150;unique;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	RegisteredAt time.Time `gorm:"autoCreateTime"`
}
