package models

import "time"

type Usuario struct {
	ID            uint      `gorm:"primaryKey"`
	Nombre        string    `gorm:"size:150;not null"`
	Correo        string    `gorm:"size:150;unique;not null"`
	PasswordHash  string    `gorm:"size:255;not null"`
	FechaRegistro time.Time `gorm:"autoCreateTime"`
}
