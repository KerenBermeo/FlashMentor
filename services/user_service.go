package service

import (
	"errors"

	"github.com/KerenBermeo/FlashMentor/models"
	"github.com/KerenBermeo/FlashMentor/repositories"
	"golang.org/x/crypto/bcrypt"
)

// User Service
type UserService struct {
	repo *repositories.UserRepo
}

// Constructor
func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{repo: repo}
}

// Specific user methods
func (s *UserService) RegisterUser(user *models.User) error {
	// Validate the user before registering
	if err := s.validateUser(user); err != nil {
		return err
	}

	// Encrypt the password before saving
	hashedPassword, err := s.encryptPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	return s.repo.Create(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) UpdateUser(user *models.User) error {
	// Validate the user before updating
	if err := s.validateUser(user); err != nil {
		return err
	}

	// Encrypt the password before updating
	hashedPassword, err := s.encryptPassword(user.PasswordHash)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)

	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}

// validateUser performs basic user validations
func (s *UserService) validateUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("user name is required")
	}
	if user.Email == "" {
		return errors.New("user email is required")
	}
	if user.PasswordHash == "" {
		return errors.New("user password is required")
	}
	return nil
}

// encryptPassword encrypts the user's password
func (s *UserService) encryptPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword checks if the provided password matches the encrypted password
func (s *UserService) VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
