package routes

import (
	"net/http"
	"strconv"

	m "github.com/KerenBermeo/FlashMentor/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	temas := r.Group("/temas")
	{
		temas.POST("", func(c *gin.Context) {
			var tema m.Tema
			if err := c.ShouldBindJSON(&tema); err == nil {
				if err := tema.CreateTema(db); err == nil {
					c.JSON(http.StatusOK, tema)
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				}
			}
		})

		temas.GET("", func(c *gin.Context) {
			temas, err := m.GetAllTemas(db)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, temas)
		})

		temas.PUT("/:id", func(c *gin.Context) {
			var tema m.Tema
			idStr := c.Param("id")
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				c.JSON(400, gin.H{"error": "ID inválido"})
				return
			}

			if err := c.ShouldBindJSON(&tema); err == nil {
				tema.ID = uint(id)
				if err := tema.UpdateTema(db); err == nil {
					c.JSON(http.StatusOK, tema)
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				}
			}
		})

		temas.DELETE("/:id", func(c *gin.Context) {
			var tema m.Tema
			idStr := c.Param("id")
			id, err := strconv.ParseUint(idStr, 10, 64)
			if err != nil {
				c.JSON(400, gin.H{"error": "ID inválido"})
				return
			}
			tema.ID = uint(id)
			if err := tema.DeleteTema(db); err == nil {
				c.JSON(http.StatusOK, gin.H{"message": "Tema eliminado"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		})
	}

	niveles := r.Group("niveles")
	{
		niveles.GET("", func(c *gin.Context) {
			nivel, err := m.GetAllNiveles(db)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, nivel)
		})
	}

	preguntas := r.Group("preguntas/:tema_id/:nivel_id/:tipo")
	{
		preguntas.GET("", func(c *gin.Context) {
			// Obtener parámetros
			temaID, err1 := strconv.Atoi(c.Param("tema_id"))
			nivelID, err2 := strconv.Atoi(c.Param("nivel_id"))
			tipo, err3 := strconv.Atoi(c.Param("tipo"))

			// Validar parámetros
			if err1 != nil || err2 != nil || err3 != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Los parámetros deben ser numéricos"})
				return
			}

			var preguntas []interface{}

			// Validar tipo y consultar la tabla correspondiente
			switch tipo {
			case 1: // Tipo 1 para preguntas abiertas
				if err := db.Table("preguntas_abiertas").
					Where("tema_id = ? AND nivel_id = ?", temaID, nivelID).
					Find(&preguntas).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener preguntas abiertas"})
					return
				}
			case 2: // Tipo 2 para preguntas cerradas
				if err := db.Table("preguntas_cerradas").
					Where("tema_id = ? AND nivel_id = ?", temaID, nivelID).
					Find(&preguntas).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener preguntas cerradas"})
					return
				}
			default:
				c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo inválido. Use 1 para abiertas o 2 para cerradas"})
				return
			}

			// Retornar preguntas
			c.JSON(http.StatusOK, preguntas)
		})
	}

}
