package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/handler"
)

func main() {
	router := gin.Default()

	healthHandler := handler.NewHealthHandler()
	router.GET("/health", healthHandler.HealthCheck)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
