package models

type PreguntasCerradas struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	TemaID    uint    `gorm:"not null"`
	NivelID   uint    `gorm:"not null"`
	Pregunta  string  `gorm:"type:text;not null"`
	Respuesta bool    `gorm:"not null"`
	UsuarioID uint    `gorm:"not null"`
	Tema      Temas   `gorm:"foreignKey:TemaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Nivel     Nivel   `gorm:"foreignKey:NivelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Usuario   Usuario `gorm:"foreignKey:UsuarioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
