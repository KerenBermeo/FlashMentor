package models

import "gorm.io/gorm"

type PreguntaAbierta struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	TemaID    uint   `gorm:"not null"`
	NivelID   uint   `gorm:"not null"`
	Pregunta  string `gorm:"type:varchar(300);not null"`
	Respuesta string `gorm:"type:text;not null"`
	Tema      Tema   `gorm:"foreignKey:TemaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Nivel     Nivel  `gorm:"foreignKey:NivelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Crear una nueva pregunta abierta
func (p *PreguntaAbierta) CreatePreguntaAbierta(db *gorm.DB) error {
	return db.Create(&p).Error
}

// Obtener preguntas abiertas por tema y nivel
func GetPreguntasAbiertasByTemaYNivel(db *gorm.DB, temaID, nivelID uint) ([]PreguntaAbierta, error) {
	var preguntas []PreguntaAbierta
	if err := db.Where("tema_id = ? AND nivel_id = ?", temaID, nivelID).Find(&preguntas).Error; err != nil {
		return nil, err
	}
	return preguntas, nil
}

// Actualizar una pregunta abierta
func (p *PreguntaAbierta) UpdatePreguntaAbierta(db *gorm.DB) error {
	return db.Save(&p).Error
}

// Borrar una pregunta abierta
func (p *PreguntaAbierta) DeletePreguntaAbierta(db *gorm.DB) error {
	return db.Delete(&p).Error
}
