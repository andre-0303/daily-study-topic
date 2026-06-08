package handlers

import (
	"net/http"
	"study-topics-cicd/internal/service"

	"github.com/gin-gonic/gin"
)

func GetTopic(c *gin.Context) {
	topic := service.GetRandomTopic()

	c.JSON(http.StatusOK, topic)
}