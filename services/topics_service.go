package service

import (
	"errors"
	"log"

	"github.com/KerenBermeo/FlashMentor/models"
	"github.com/KerenBermeo/FlashMentor/repositories"
)

// TopicsService estructura del servicio de temas
type TopicsService struct {
	repo *repositories.TopicsRepo
}

// Constructor
func NewTopicsService(repo *repositories.TopicsRepo) *TopicsService {
	return &TopicsService{repo: repo}
}

// CreateTopic crea un nuevo tema
func (s *TopicsService) CreateTopic(topic *models.Topics) error {
	// Validar que el nombre del tema no esté vacío
	if topic.TopicName == "" {
		return errors.New("el nombre del tema es requerido")
	}

	// Validar que el UserID no sea cero
	if topic.UserID == 0 {
		return errors.New("el ID del usuario es requerido")
	}

	// Crear el tema en el repositorio
	if err := s.repo.Create(topic); err != nil {
		log.Printf("Error al crear el tema: %v", err)
		return errors.New("error al crear el tema")
	}

	return nil
}

// GetTopicByID obtiene un tema por su ID
func (s *TopicsService) GetTopicByID(id uint) (*models.Topics, error) {
	// Validar que el ID no sea cero
	if id == 0 {
		return nil, errors.New("el ID del tema es requerido")
	}

	// Obtener el tema del repositorio
	topic, err := s.repo.GetByID(id)
	if err != nil {
		log.Printf("Error al obtener el tema: %v", err)
		return nil, errors.New("error al obtener el tema")
	}

	return topic, nil
}

// GetAllTopics obtiene todos los temas
func (s *TopicsService) GetAllTopics() ([]models.Topics, error) {
	// Obtener todos los temas del repositorio
	topics, err := s.repo.GetAll()
	if err != nil {
		log.Printf("Error al obtener todos los temas: %v", err)
		return nil, errors.New("error al obtener los temas")
	}

	return topics, nil
}

// UpdateTopic actualiza un tema existente
func (s *TopicsService) UpdateTopic(topic *models.Topics) error {
	// Validar que el ID del tema no sea cero
	if topic.ID == 0 {
		return errors.New("el ID del tema es requerido")
	}

	// Validar que el nombre del tema no esté vacío
	if topic.TopicName == "" {
		return errors.New("el nombre del tema es requerido")
	}

	// Actualizar el tema en el repositorio
	if err := s.repo.Update(topic); err != nil {
		log.Printf("Error al actualizar el tema: %v", err)
		return errors.New("error al actualizar el tema")
	}

	return nil
}

// DeleteTopic elimina un tema y sus preguntas asociadas en cascada
func (s *TopicsService) DeleteTopic(id uint) error {
	// Validar que el ID no sea cero
	if id == 0 {
		return errors.New("el ID del tema es requerido")
	}

	// Eliminar el tema y sus preguntas asociadas en cascada
	if err := s.repo.Delete(id); err != nil {
		log.Printf("Error al eliminar el tema: %v", err)
		return errors.New("error al eliminar el tema")
	}

	return nil
}
