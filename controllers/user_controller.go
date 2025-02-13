package controllers

import (
	"net/http"
	"strconv"

	"github.com/KerenBermeo/FlashMentor/models"
	"github.com/KerenBermeo/FlashMentor/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UsuarioController estructura del controlador
type UsuarioController struct {
	repo *repositories.BaseRepository[models.Usuario]
}

// NewUsuarioController crea una nueva instancia del controlador
func NewUsuarioController(db *gorm.DB) *UsuarioController {
	return &UsuarioController{repo: repositories.NewBaseRepository[models.Usuario](db)}
}

// CrearUsuario maneja la solicitud para crear un usuario
func (uc *UsuarioController) CrearUsuario(c *gin.Context) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.repo.Create(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

// ObtenerUsuarioPorID maneja la solicitud para obtener un usuario por ID
func (uc *UsuarioController) ObtenerUsuarioPorID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	usuario, err := uc.repo.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// ObtenerTodosLosUsuarios maneja la solicitud para obtener todos los usuarios
func (uc *UsuarioController) ObtenerTodosLosUsuarios(c *gin.Context) {
	usuarios, err := uc.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

// ActualizarUsuario maneja la solicitud para actualizar un usuario
func (uc *UsuarioController) ActualizarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var datosActualizados models.Usuario

	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	datosActualizados.ID = uint(id) // Asegurar que el ID es correcto
	if err := uc.repo.Update(&datosActualizados); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado"})
}

// EliminarUsuario maneja la solicitud para eliminar un usuario
func (uc *UsuarioController) EliminarUsuario(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := uc.repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
