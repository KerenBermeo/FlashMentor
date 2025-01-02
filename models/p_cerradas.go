package models

import "gorm.io/gorm"

type PreguntaCerrada struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	TemaID    uint   `gorm:"not null"`
	NivelID   uint   `gorm:"not null"`
	Pregunta  string `gorm:"type:text;not null"`
	Respuesta bool   `gorm:"not null"`
	Tema      Tema   `gorm:"foreignKey:TemaID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Nivel     Nivel  `gorm:"foreignKey:NivelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// Crear una nueva pregunta cerrada
func (p *PreguntaCerrada) CreatePreguntaCerrada(db *gorm.DB) error {
	return db.Create(&p).Error
}

// Obtener preguntas cerradas por tema y nivel
func GetPreguntasCerradasByTemaYNivel(db *gorm.DB, temaID, nivelID uint) ([]PreguntaCerrada, error) {
	var preguntas []PreguntaCerrada
	if err := db.Where("tema_id = ? AND nivel_id = ?", temaID, nivelID).Find(&preguntas).Error; err != nil {
		return nil, err
	}
	return preguntas, nil
}

// Actualizar una pregunta cerrada
func (p *PreguntaCerrada) UpdatePreguntaCerrada(db *gorm.DB) error {
	return db.Save(&p).Error
}

// Borrar una pregunta cerrada
func (p *PreguntaCerrada) DeletePreguntaCerrada(db *gorm.DB) error {
	return db.Delete(&p).Error
}
