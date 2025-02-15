package models

type Level string

const (
	Basic        Level = "BASIC"
	Intermediate Level = "INTERMEDIATE"
	Advanced     Level = "ADVANCED"
)

type Levels struct {
	ID        uint  `gorm:"primaryKey"`
	LevelName Level `gorm:"type:enum('BASIC','INTERMEDIATE','ADVANCED');not null"`
}
