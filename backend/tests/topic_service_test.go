package tests

import (
	"study-topics-cicd/internal/service"
	"testing"
)

func TestGetRandomTopic(t *testing.T) {
	topic := service.GetRandomTopic()

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
