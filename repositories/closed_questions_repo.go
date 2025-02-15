package repositories

import (
	"github.com/KerenBermeo/FlashMentor/models"
	"gorm.io/gorm"
)

type ClosedQuestionsRepo struct {
	*BaseRepository[models.ClosedQuestions]
}

func NewClosedQuestionsRepository(db *gorm.DB) *ClosedQuestionsRepo {
	return &ClosedQuestionsRepo{
		BaseRepository: NewBaseRepository[models.ClosedQuestions](db),
	}
}
