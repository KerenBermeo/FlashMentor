package repositories

import (
	"github.com/KerenBermeo/FlashMentor/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	*BaseRepository[models.User]
}

// Constructor
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		BaseRepository: NewBaseRepository[models.User](db),
	}
}
