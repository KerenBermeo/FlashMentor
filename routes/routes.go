package routes

import (
	"github.com/KerenBermeo/FlashMentor/controllers"
	"github.com/KerenBermeo/FlashMentor/repositories"
	service "github.com/KerenBermeo/FlashMentor/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Inicialización de los repositorios
	userRepo := repositories.NewUserRepo(db)
	temasRepo := repositories.NewTopicRepository(db)

	// Inicialización de los servicios
	userService := service.NewUserService(userRepo)
	temasService := service.NewTopicsService(temasRepo)

	// Inicialización de los controladores con sus respectivos servicios
	usuarioController := controllers.NewUsuarioController(userService)
	temasController := controllers.NewTopicsController(temasService)

	// Prefijo para futuras versiones de la API
	api := r.Group("/api/v1")

	// Grupo de rutas para usuarios
	usuarios := api.Group("/usuarios")
	{
		usuarios.POST("/", usuarioController.CrearUsuario)
		usuarios.GET("/", usuarioController.ObtenerTodosLosUsuarios)
		usuarios.GET("/:id", usuarioController.ObtenerUsuarioPorID)
		usuarios.PUT("/:id", usuarioController.ActualizarUsuario)
		usuarios.DELETE("/:id", usuarioController.EliminarUsuario)
	}

	// Grupo de rutas para temas
	temas := api.Group("/temas")
	{
		temas.POST("/", temasController.CreateTopic)
		temas.GET("/", temasController.GetAllTopics)
		temas.GET("/:id", temasController.GetTopicByID)
		temas.PUT("/:id", temasController.UpdateTopic)
		temas.DELETE("/:id", temasController.DeleteTopic)
	}
}
