package handlers

import (
	"io"
	"net/http"

	// "path/filepath"

	"os"

	"github.com/Vince33/media-metadata-api/utils"
	"github.com/gin-gonic/gin"
)

const maxUploadSize = 10 << 20 // 10 MiB
// ExtractHandler handles the file upload and metadata extraction
// It expects a file to be uploaded with the key "file" in the form data.
// The file is saved to a temporary location, validated for type, and metadata is extracted using ffprobe.

func ExtractHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	defer file.Close()

	// Create temp file
	tempFile, err := os.CreateTemp("", "uploaded-*.mp4")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create temp file"})
		return
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	// Enforce size limit
	limitedReader := &io.LimitedReader{R: file, N: maxUploadSize + 1}
	written, err := io.Copy(tempFile, limitedReader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
		return
	}

	if written > maxUploadSize {
		c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "file too large"})
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
