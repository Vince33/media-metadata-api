package handlers

import (
	"net/http"
	"path/filepath"

	"os"

	"github.com/Vince33/media-metadata-api/utils"
	"github.com/gin-gonic/gin"
)

// ExtractHandler handles file uploads and returns  a dummy JSON for now
func ExtractHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}
	// Save the file to a temporary location
	savePath := filepath.Join("media", file.Filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Validate the file type after saving
	allowedTypes := []string{"video/mp4", "video/mpeg", "video/quicktime"}
	if !utils.IsValidMimeType(savePath, allowedTypes) {
		os.Remove(savePath) // Clean up the saved file
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file type"})
		return
	}

	// TODO: Extract metadata using ffprobe
	metadata, err := utils.ExtractMetadata(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract metadata", "details": err.Error()})
		return
	}
	c.JSON(http.StatusOK, metadata)
}
