package service

import (
	"study-topics-cicd/internal/models"
	"study-topics-cicd/internal/repository"
)

type TopicService struct {
	repo *repository.TopicRepository
}

func NewTopicService(repo *repository.TopicRepository) *TopicService {
	return &TopicService{repo: repo}
}

func (s *TopicService) GetRandomTopic() (models.Topic, error) {
	return s.repo.GetRandom()
}
