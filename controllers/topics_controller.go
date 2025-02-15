package controllers

import (
	"net/http"
	"strconv"

	"github.com/KerenBermeo/FlashMentor/models"
	service "github.com/KerenBermeo/FlashMentor/services"
	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	topicsService *service.TopicsService
}

// Constructor
func NewTopicsController(topicsService *service.TopicsService) *TopicsController {
	return &TopicsController{topicsService: topicsService}
}

func (tc *TopicsController) CreateTopic(c *gin.Context) {
	var topic models.Topics
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	if err := tc.topicsService.CreateTopic(&topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, topic)
}

func (tc *TopicsController) GetTopicByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	topic, err := tc.topicsService.GetTopicByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (tc *TopicsController) GetAllTopics(c *gin.Context) {
	topics, err := tc.topicsService.GetAllTopics()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topics)
}

func (tc *TopicsController) UpdateTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var topic models.Topics
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	topic.ID = uint(id)
	if err := tc.topicsService.UpdateTopic(&topic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, topic)
}

func (tc *TopicsController) DeleteTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := tc.topicsService.DeleteTopic(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tema eliminado correctamente"})
}
