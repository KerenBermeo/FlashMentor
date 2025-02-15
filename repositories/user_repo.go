package repositories

import (
	"errors"

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

func (r *UserRepo) GetByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.BaseRepository.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil // Si no se encuentra, retornamos nil sin error.
		}
		return nil, result.Error
	}
	return &user, nil
}
