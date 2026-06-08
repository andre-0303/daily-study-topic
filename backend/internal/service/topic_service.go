package service

import (
	"math/rand"
	"study-topics-cicd/internal/models"
)

var topics = []models.Topic{
	{
		ID:          1,
		Title:       "Structs em Go",
		Difficulty:  "intermediário",
		Description: "Aprenda composição de dados",
	},
	{
		ID:          2,
		Title:       "Interfaces em Go",
		Difficulty:  "avançado",
		Description: "Entenda polimorfismo em Go",
	},
	{
		ID:          3,
		Title:       "Slices",
		Difficulty:  "iniciante",
		Description: "Manipulação dinâmica de arrays",
	},
}

func GetRandomTopic() models.Topic {
	randomIndex := rand.Intn(len(topics))
	return topics[randomIndex]
}