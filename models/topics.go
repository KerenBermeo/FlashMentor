package models

type Topics struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	TopicName string `gorm:"type:varchar(150);not null"`
	UserID    uint   `gorm:"not null"`
}
