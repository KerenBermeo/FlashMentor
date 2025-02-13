package models

import "time"

type Caja struct {
	ID                      uint      `gorm:"primaryKey;autoIncrement"`
	Caja                    int       `gorm:"type:enum('1','2','3','4','5');not null"`
	UsuarioID               uint      `gorm:"not null"`
	PreguntaID              uint      `gorm:"not null"`
	TipoPregunta            string    `gorm:"type:enum('ABIERTA','CERRADA');not null"`
	TemaID                  uint      `gorm:"not null"`
	NivelID                 uint      `gorm:"not null"`
	RespondidaCorrectamente bool      `gorm:"not null"`
	FechaRespuesta          time.Time `gorm:"autoCreateTime"`

	// Relaciones con otras tablas
	Usuario Usuario `gorm:"foreignKey:UsuarioID"`
	Tema    Temas   `gorm:"foreignKey:TemaID"`
	Nivel   Nivel   `gorm:"foreignKey:NivelID"`
}
