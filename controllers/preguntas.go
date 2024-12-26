package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOpenQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Obteniendo todas las preguntas abiertas"})
}

func CreateOpenQuestion(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Pregunta abierta creada"})
}

func GetOpenQuestionByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Obteniendo pregunta abierta"})
}

func UpdateOpenQuestion(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Pregunta abierta actualizada"})
}

func DeleteOpenQuestion(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Pregunta abierta eliminada"})
}

func GetAllClosedQuestions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Obteniendo todas las preguntas cerradas"})
}

func CreateClosedQuestion(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Pregunta cerrada creada"})
}

func GetClosedQuestionByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Obteniendo pregunta cerrada"})
}

func UpdateClosedQuestion(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Pregunta cerrada actualizada"})
}

func DeleteClosedQuestion(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "Pregunta cerrada eliminada"})
}
