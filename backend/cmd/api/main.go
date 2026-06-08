package main

import (
	"study-topics-cicd/internal/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	routes.SetupRoutes(router)

	router.Run(":8080")
}