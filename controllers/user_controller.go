package controllers

import (
	"net/http"
	"strconv"

	"github.com/KerenBermeo/FlashMentor/models"
	service "github.com/KerenBermeo/FlashMentor/services"
	"github.com/gin-gonic/gin"
)

// UsuarioController estructura del controlador
type UsuarioController struct {
	userService *service.UserService
}

// NewUsuarioController crea una nueva instancia del controlador
func NewUsuarioController(userService *service.UserService) *UsuarioController {
	return &UsuarioController{userService: userService}
}

// CrearUsuario maneja la solicitud para crear un usuario
func (uc *UsuarioController) CrearUsuario(c *gin.Context) {
	var usuario models.User
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if err := uc.userService.RegisterUser(&usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

// ObtenerUsuarioPorID maneja la solicitud para obtener un usuario por ID
func (uc *UsuarioController) ObtenerUsuarioPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	usuario, err := uc.userService.GetUser(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

// ObtenerTodosLosUsuarios maneja la solicitud para obtener todos los usuarios
func (uc *UsuarioController) ObtenerTodosLosUsuarios(c *gin.Context) {
	usuarios, err := uc.userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, usuarios)
}

// ActualizarUsuario maneja la solicitud para actualizar un usuario
func (uc *UsuarioController) ActualizarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var datosActualizados models.User
	if err := c.ShouldBindJSON(&datosActualizados); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	datosActualizados.ID = uint(id)
	if err := uc.userService.UpdateUser(&datosActualizados); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, datosActualizados)
}

// EliminarUsuario maneja la solicitud para eliminar un usuario
func (uc *UsuarioController) EliminarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := uc.userService.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar usuario: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado correctamente"})
}
