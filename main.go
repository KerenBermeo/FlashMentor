package main

import (
	"fmt"
	"log"

	"github.com/KerenBermeo/FlashMentor/db"
	"github.com/KerenBermeo/FlashMentor/middleware"
	"github.com/KerenBermeo/FlashMentor/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := db.ConnectionBD()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Conexión exitosa a la base de datos", db)

	// Crear el router de Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(middleware.SetupCORS())

	// Configurar las rutas
	routes.SetupRoutes(r, db)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8080")
}
