package models

import "gorm.io/gorm"

type Tema struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	NombreTema string `gorm:"type:varchar(150);not null"`
}

// Crear un nuevo tema
func (t *Tema) CreateTema(db *gorm.DB) error {
	return db.Create(&t).Error
}

// Obtener todos los temas
func GetAllTemas(db *gorm.DB) ([]Tema, error) {
	var temas []Tema
	if err := db.Find(&temas).Error; err != nil {
		return nil, err
	}
	return temas, nil
}

// Actualizar el nombre de un tema
func (t *Tema) UpdateTema(db *gorm.DB) error {
	return db.Save(&t).Error
}

// Borrar un tema y sus preguntas asociadas
func (t *Tema) DeleteTema(db *gorm.DB) error {
	// Borrar las preguntas asociadas
	if err := db.Where("tema_id = ?", t.ID).Delete(&PreguntaAbierta{}).Error; err != nil {
		return err
	}
	if err := db.Where("tema_id = ?", t.ID).Delete(&PreguntaCerrada{}).Error; err != nil {
		return err
	}
	// Borrar el tema
	return db.Delete(&t).Error
}
