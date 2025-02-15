package repositories

import (
	"github.com/KerenBermeo/FlashMentor/models"
	"gorm.io/gorm"
)

type TopicsRepo struct {
	*BaseRepository[models.Topics]
}

// Constructor
func NewTopicRepository(db *gorm.DB) *TopicsRepo {
	return &TopicsRepo{
		BaseRepository: NewBaseRepository[models.Topics](db),
	}
}
