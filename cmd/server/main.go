package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/handler"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/service/memory"
)

func main() {
	router := gin.Default()

	healthHandler := handler.NewHealthHandler()
	router.GET("/health", healthHandler.HealthCheck)

	urlService := memory.NewURLService()
	urlHandler := handler.NewURLHandler(urlService)
	router.POST("/shorten", urlHandler.Shorten)
	router.GET("/:code", urlHandler.Redirect)

	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
