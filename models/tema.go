package models

type Temas struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	NombreTema string `gorm:"type:varchar(150);not null"`
	UsuarioID  uint   `gorm:"not null"`
}
