package main

import (
	"fmt"
	"os"
	"study-topics-cicd/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	routes.SetupRoutes(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)

	router.Run(":" + port)
}