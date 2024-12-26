package routes

import (
	"github.com/KerenBermeo/FlashMentor/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Grupo de rutas para preguntas abiertas
	openQuestionsGroup := r.Group("/preguntas_abiertas")
	{
		openQuestionsGroup.GET("/", controllers.GetAllOpenQuestions)
		openQuestionsGroup.POST("/", controllers.CreateOpenQuestion)
		openQuestionsGroup.GET("/:id", controllers.GetOpenQuestionByID)
		openQuestionsGroup.PUT("/:id", controllers.UpdateOpenQuestion)
		openQuestionsGroup.DELETE("/:id", controllers.DeleteOpenQuestion)
	}

	// Grupo de rutas para preguntas cerradas
	closedQuestionsGroup := r.Group("/preguntas_cerradas")
	{
		closedQuestionsGroup.GET("/", controllers.GetAllClosedQuestions)
		closedQuestionsGroup.POST("/", controllers.CreateClosedQuestion)
		closedQuestionsGroup.GET("/:id", controllers.GetClosedQuestionByID)
		closedQuestionsGroup.PUT("/:id", controllers.UpdateClosedQuestion)
		closedQuestionsGroup.DELETE("/:id", controllers.DeleteClosedQuestion)
	}
}
