package main

import (
	"github.com/KerenBermeo/FlashMentor/middleware"
	"github.com/KerenBermeo/FlashMentor/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Crear el router de Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(middleware.SetupCORS())

	// Configurar las rutas
	routes.SetupRoutes(r)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
