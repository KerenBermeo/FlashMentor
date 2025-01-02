package routes

import (
	"github.com/KerenBermeo/FlashMentor/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	QuestionsGroup := r.Group("/question")
	{
		QuestionsGroup.GET("/", controllers.GetQuestions())

	}

}
