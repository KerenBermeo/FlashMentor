package db

import (
	"fmt"
	"log"
	"os"

	"github.com/KerenBermeo/FlashMentor/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ConnectionBD() (*gorm.DB, error) {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Obtener las variables de entorno
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Crear el DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)

	// Abrir la conexión con GORM sin pluralizar los nombres de las tablas
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Desactivar la pluralización de los nombres de las tablas
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	// 🛠 Ejecutar migraciones automáticas
	err = db.AutoMigrate(&models.Usuario{}, &models.Temas{}, &models.Niveles{}, &models.PreguntasAbiertas{}, &models.PreguntasCerradas{}, &models.Caja{}) // Agrega aquí todos tus modelos
	if err != nil {
		log.Fatal("Error en la migración:", err)
	}

	log.Println("Migraciones ejecutadas correctamente ✅")

	return db, nil
}
