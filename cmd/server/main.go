package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/handler"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/middleware"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/service"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/service/memory"
	"github.com/work-kumar-rajesh/go-url-shortner/internal/service/postgres"
	"golang.org/x/time/rate"
)

func main() {
	router := gin.Default()

	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Retrieve environment variables
	servicePort := os.Getenv("SERVICE_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	// database service
	dbConnStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbName, dbSSLMode)
	db, err := sqlx.Connect("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	var urlDbService service.URLService = postgres.New(db)
	fmt.Print(urlDbService)

	//memory service
	urlService := memory.NewURLService()

	healthHandler := handler.NewHealthHandler()
	router.GET("/health", healthHandler.HealthCheck)

	rateLimiter := middleware.NewRateLimiter(rate.Every(6*time.Second), 5) // ~5 req/minute if used continuously.
	router.Use(rateLimiter.RateLimitMiddleware())

	urlHandler := handler.NewURLHandler(urlService) // here we can use memory service or database service as required
	router.POST("/shorten", urlHandler.Shorten)
	router.GET("/:code", urlHandler.Redirect)
	router.GET("/analytics/:code", urlHandler.GetAnalytics)

	log.Println("Starting server on :8080")
	if err := router.Run(":" + servicePort); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
