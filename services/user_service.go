package service

import (
	"errors"

	"github.com/KerenBermeo/FlashMentor/models"
	"github.com/KerenBermeo/FlashMentor/repositories"
	"golang.org/x/crypto/bcrypt"
)

// Servicio para Usuarios
type UsuarioService struct {
	repo *repositories.UsuarioRepository
}

// Constructor
func NewUsuarioService(repo *repositories.UsuarioRepository) *UsuarioService {
	return &UsuarioService{repo: repo}
}

// Métodos específicos de usuario
func (s *UsuarioService) RegistrarUsuario(usuario *models.Usuario) error {
	// Validar el usuario antes de registrarlo
	if err := s.validarUsuario(usuario); err != nil {
		return err
	}

	// Encriptar la contraseña antes de guardarla
	hashedPassword, err := s.encriptarContraseña(usuario.PasswordHash)
	if err != nil {
		return err
	}
	usuario.PasswordHash = string(hashedPassword)

	return s.repo.Create(usuario)
}

func (s *UsuarioService) ObtenerUsuario(id uint) (*models.Usuario, error) {
	return s.repo.GetByID(id)
}

func (s *UsuarioService) ObtenerTodosLosUsuarios() ([]models.Usuario, error) {
	return s.repo.GetAll()
}

func (s *UsuarioService) ActualizarUsuario(usuario *models.Usuario) error {
	// Validar el usuario antes de actualizarlo
	if err := s.validarUsuario(usuario); err != nil {
		return err
	}

	// Encriptar la contraseña antes de actualizarla
	hashedPassword, err := s.encriptarContraseña(usuario.PasswordHash)
	if err != nil {
		return err
	}
	usuario.PasswordHash = string(hashedPassword)

	return s.repo.Update(usuario)
}

func (s *UsuarioService) EliminarUsuario(id uint) error {
	return s.repo.Delete(id)
}

// validarUsuario realiza validaciones básicas en el usuario
func (s *UsuarioService) validarUsuario(usuario *models.Usuario) error {
	if usuario.Nombre == "" {
		return errors.New("el nombre del usuario es requerido")
	}
	if usuario.Correo == "" {
		return errors.New("el email del usuario es requerido")
	}
	if usuario.PasswordHash == "" {
		return errors.New("la contraseña del usuario es requerida")
	}
	return nil
}

// encriptarContraseña encripta la contraseña del usuario
func (s *UsuarioService) encriptarContraseña(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerificarContraseña verifica si la contraseña proporcionada coincide con la contraseña encriptada
func (s *UsuarioService) VerificarContraseña(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
