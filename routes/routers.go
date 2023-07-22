package routes

import (
	"idea-generator/controller"

	"github.com/gin-gonic/gin"
)

func IdeaRoute(router *gin.Engine) {
	router.GET("/idea", controller.GetAllIdeas)
	router.POST("/idea", controller.CreateIdea)
	router.DELETE("/idea/:id", controller.DeleteIdea)
	router.PUT("/idea/:id", controller.UpdateIdea)
	router.GET("/idea/:id", controller.GetIdeaByID)
}
