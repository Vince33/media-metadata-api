package handlers

import (
	"io"
	"net/http"
	// "path/filepath"

	"os"

	"github.com/Vince33/media-metadata-api/utils"
	"github.com/gin-gonic/gin"
)

// ExtractHandler handles the file upload and metadata extraction
// It expects a file to be uploaded with the key "file" in the form data.
// The file is saved to a temporary location, validated for type, and metadata is extracted using ffprobe.
func ExtractHandler(c *gin.Context) {
	// Parse uploaded file
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	defer file.Close()

	// Create a temporary file
	tempFile, err := os.CreateTemp("", "uploaded-*.mp4")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create temp file"})
		return
	}
	defer os.Remove(tempFile.Name()) // Ensure cleanup
	defer tempFile.Close()

	// Copy uploaded content to the temp file
	if _, err := io.Copy(tempFile, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	// Extract metadata
	metadata, err := utils.ExtractMetadata(tempFile.Name())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to extract metadata"})
		return
	}

	c.JSON(http.StatusOK, metadata)
}
