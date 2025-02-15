package models

import "time"

type Boxes struct {
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	Box             int       `gorm:"type:enum('1','2','3','4','5');not null"`
	UserID          uint      `gorm:"not null"`
	QuestionID      uint      `gorm:"not null"`
	TypeQuestion    string    `gorm:"type:enum('OPEN','CLOSE');not null"`
	TopicID         uint      `gorm:"not null"`
	LevelID         uint      `gorm:"not null"`
	ResponseCorrect bool      `gorm:"not null"`
	DateResponse    time.Time `gorm:"autoCreateTime"`

	// Relaciones con otras tablas
	Usuario User   `gorm:"foreignKey:UserID"`
	Tema    Topics `gorm:"foreignKey:TopicID"`
	Nivel   Levels `gorm:"foreignKey:LevelID"`
}
