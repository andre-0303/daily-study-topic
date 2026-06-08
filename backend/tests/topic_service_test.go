package tests

import (
	"testing"

	"study-topics-cicd/internal/db"
	"study-topics-cicd/internal/repository"
	"study-topics-cicd/internal/service"
)

func TestGetRandomTopic(t *testing.T) {
	conn, err := db.Open(":memory:")
	if err != nil {
		t.Fatalf("failed to open test database: %v", err)
	}
	defer conn.Close()

	topicService := service.NewTopicService(repository.NewTopicRepository(conn))

	topic, err := topicService.GetRandomTopic()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if topic.ID == 0 {
		t.Errorf("expected valid topic ID, got %d", topic.ID)
	}

	if topic.Title == "" {
		t.Error("expected topic title, got empty string")
	}

	if topic.Description == "" {
		t.Error("expected topic description, got empty string")
	}
}
