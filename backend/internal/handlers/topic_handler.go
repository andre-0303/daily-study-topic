package handlers

import (
	"net/http"

	"study-topics-cicd/internal/service"

	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	service *service.TopicService
}

func NewTopicHandler(service *service.TopicService) *TopicHandler {
	return &TopicHandler{service: service}
}

func (h *TopicHandler) GetTopic(c *gin.Context) {
	topic, err := h.service.GetRandomTopic()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch topic"})
		return
	}

	c.JSON(http.StatusOK, topic)
}
