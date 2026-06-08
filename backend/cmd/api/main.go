package main

import (
	"fmt"
	"log"
	"os"

	"study-topics-cicd/internal/db"
	"study-topics-cicd/internal/handlers"
	"study-topics-cicd/internal/repository"
	"study-topics-cicd/internal/routes"
	"study-topics-cicd/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "topics.db"
	}

	conn, err := db.Open(dbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer conn.Close()

	topicRepo := repository.NewTopicRepository(conn)
	topicService := service.NewTopicService(topicRepo)
	topicHandler := handlers.NewTopicHandler(topicService)

	router := gin.Default()

	router.Use(cors.Default())

	routes.SetupRoutes(router, topicHandler)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)

	router.Run(":" + port)
}