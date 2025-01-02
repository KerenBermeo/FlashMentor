package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context, type_question string, level int) {
	if type_question == "open_question" {

	} else if type_question == "close_question" {

	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No se envio el tipo de pregunta"})
	}
}
