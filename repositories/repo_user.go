package repositories

import (
	"github.com/KerenBermeo/FlashMentor/models"
	"gorm.io/gorm"
)

type UsuarioRepository struct {
	*BaseRepository[models.Usuario]
}

// Constructor
func NewUsuarioRepository(db *gorm.DB) *UsuarioRepository {
	return &UsuarioRepository{
		BaseRepository: NewBaseRepository[models.Usuario](db),
	}
}
