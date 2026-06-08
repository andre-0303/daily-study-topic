package repository

import (
	"database/sql"

	"study-topics-cicd/internal/models"
)

type TopicRepository struct {
	db *sql.DB
}

func NewTopicRepository(db *sql.DB) *TopicRepository {
	return &TopicRepository{db: db}
}

func (r *TopicRepository) GetRandom() (models.Topic, error) {
	var topic models.Topic

	row := r.db.QueryRow(
		"SELECT id, title, difficulty, description FROM topics ORDER BY RANDOM() LIMIT 1",
	)

	err := row.Scan(&topic.ID, &topic.Title, &topic.Difficulty, &topic.Description)
	if err != nil {
		return models.Topic{}, err
	}

	return topic, nil
}
