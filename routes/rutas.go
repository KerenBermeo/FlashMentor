package routes

import (
	"net/http"
	"strconv"

	m "github.com/KerenBermeo/FlashMentor/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	// Rutas para los temas
	r.POST("/temas", func(c *gin.Context) {
		var tema m.Tema
		if err := c.ShouldBindJSON(&tema); err == nil {
			if err := tema.CreateTema(db); err == nil {
				c.JSON(http.StatusOK, tema)
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		}
	})

	r.GET("/temas", func(c *gin.Context) {
		temas, err := m.GetAllTemas(db)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, temas)
	})

	r.PUT("/temas/:id", func(c *gin.Context) {
		var tema m.Tema
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			// Manejar el error en caso de que la conversión falle
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

	r.DELETE("/temas/:id", func(c *gin.Context) {
		var tema m.Tema
		idStr := c.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			// Manejar el error en caso de que la conversión falle
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
