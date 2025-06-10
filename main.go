package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/Vince33/media-metadata-api/handlers"
	"github.com/Vince33/media-metadata-api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Set max upload size to 10MiB for now
	const maxUploadSize = 10 << 20 // 10 MiB

	// healthcheck endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	//TODO: Add more endpoints here
	//New extract endpoint
	r.POST("/extract", middleware.UploadValidationMiddleware(maxUploadSize), handlers.ExtractHandler)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
