package routes

import (
	"github.com/KerenBermeo/FlashMentor/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	usuarioController := controllers.NewUsuarioController(db)

	usuarioRoutes := r.Group("/usuarios")
	{
		usuarioRoutes.POST("/", usuarioController.CrearUsuario)
		usuarioRoutes.GET("/", usuarioController.ObtenerTodosLosUsuarios)
		usuarioRoutes.GET("/:id", usuarioController.ObtenerUsuarioPorID)
		usuarioRoutes.PUT("/:id", usuarioController.ActualizarUsuario)
		usuarioRoutes.DELETE("/:id", usuarioController.EliminarUsuario)
	}

}
