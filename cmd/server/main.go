package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/handler"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/middleware"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/service/memory"
	"golang.org/x/time/rate"
)

func main() {
	router := gin.Default()

	healthHandler := handler.NewHealthHandler()
	router.GET("/health", healthHandler.HealthCheck)

	rl := middleware.NewRateLimiter(rate.Every(6*time.Second), 5) // ~5 req/minute if used continuously.
	router.Use(rl.RateLimitMiddleware())

	urlService := memory.NewURLService()
	urlHandler := handler.NewURLHandler(urlService)
	router.POST("/shorten", urlHandler.Shorten)
	router.GET("/:code", urlHandler.Redirect)
	router.GET("/analytics/:code", urlHandler.GetAnalytics)


	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
