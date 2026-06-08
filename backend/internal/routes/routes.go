package routes

import (
	"study-topics-cicd/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, topicHandler *handlers.TopicHandler) {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/api/topic", topicHandler.GetTopic)
}
