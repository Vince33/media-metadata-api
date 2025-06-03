package main

import (
	"log"
	"net/http"

	"github.com/Vince33/media-metadata-api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	r := gin.Default()

	// healthcheck endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	//TODO: Add more endpoints here
	//New extract endpoint
	r.POST("/extract", handlers.ExtractHandler)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
