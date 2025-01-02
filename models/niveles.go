package models

import "gorm.io/gorm"

type Nivel string

const (
	Basico     Nivel = "BASICO"
	Intermedio Nivel = "INTERMEDIO"
	Avanzado   Nivel = "AVANZADO"
)

type Niveles struct {
	ID          uint  `gorm:"primaryKey"`
	NombreNivel Nivel `gorm:"type:enum('BASICO','INTERMEDIO','AVANZADO');not null"`
}

func GetAllNiveles(db *gorm.DB) ([]Nivel, error) {
	var niveles []Nivel
	if err := db.Find(&niveles).Error; err != nil {
		return nil, err
	}
	return niveles, nil
}
