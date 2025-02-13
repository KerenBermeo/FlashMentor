package repositories

import "gorm.io/gorm"

// Interfaz Genérica para el CRUD
type Repository[T any] interface {
	Create(entity *T) error
	GetByID(id uint) (*T, error)
	GetAll() ([]T, error)
	Update(entity *T) error
	Delete(id uint) error
}

// Estructura base que implementa el CRUD genérico
type BaseRepository[T any] struct {
	db *gorm.DB
}

// Constructor del repositorio
func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{db: db}
}

// Método para crear un registro
func (r *BaseRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

// Método para obtener un registro por ID
func (r *BaseRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return &entity, err
}

// Método para obtener todos los registros
func (r *BaseRepository[T]) GetAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

// Método para actualizar un registro
func (r *BaseRepository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

// Método para eliminar un registro por ID
func (r *BaseRepository[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}
