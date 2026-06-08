package main

import (
	"github.com/gin-gonic/gin"
	"study-topics-cicd/internal/routes"
)

func main() {
	router := gin.Default()
	routes.SetupRoutes(router)
	router.Run(":8080")
}
