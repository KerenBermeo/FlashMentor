package repositories

import (
	"github.com/KerenBermeo/FlashMentor/models"
	"gorm.io/gorm"
)

type OpenQuestionsRepo struct {
	*BaseRepository[models.OpenQuestions]
}

// Repository constructor
func NewOpenQuestionsRepository(db *gorm.DB) *OpenQuestionsRepo {
	return &OpenQuestionsRepo{
		BaseRepository: NewBaseRepository[models.OpenQuestions](db),
	}
}
