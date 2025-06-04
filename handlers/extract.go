package handlers

import (
	"net/http"
	"path/filepath"

	"os"

	"github.com/Vince33/media-metadata-api/utils"
	"github.com/gin-gonic/gin"
)

// ExtractHandler handles the file upload and metadata extraction
// It expects a file to be uploaded with the key "file" in the form data.
// The file is saved to a temporary location, validated for type, and metadata is extracted using ffprobe.
func ExtractHandler(c *gin.Context) {
	const maxUploadSize = 10 << 20 // 10 MiB

	// Ensures the request body does not exceed the maximum upload size
	// This is a security measure to prevent large file uploads that could exhaust server resources.
	// It also allows us to handle the file upload in a controlled manner.
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	if file.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds limit Maximum size is 10 MiB"})
		return
	}

	// Save the file to a temporary location
	savePath := filepath.Join("../media", file.Filename)
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
