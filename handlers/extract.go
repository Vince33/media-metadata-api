package handlers

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Vince33/media-metadata-api/utils"
	"github.com/gin-gonic/gin"
)

var ErrBodyTooLarge = errors.New("uploaded file is too large")

// isRequestBodyTooLarge checks whether the given error was caused by a request body
// exceeding the size limit enforced by http.MaxBytesReader.
//
// NOTE: This relies on the exact error message returned by Go's standard library:
//
//	"http: request body too large". This message may change in future versions.
//	If Go ever introduces a typed error for this case, this check should be updated.
func isRequestBodyTooLarge(err error) bool {
	return err != nil && strings.Contains(err.Error(), "http: request body too large")
}

// ExtractHandler handles the file upload and metadata extraction
// It expects a file to be uploaded with the key "file" in the form data.
// The file is saved to a temporary location, validated for type, and metadata is extracted using ffprobe.
func ExtractHandler(c *gin.Context) {
	const maxUploadSize = 10 << 20 // 10 MiB

	// Limit request body size to avoid resource exhaustion
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	file, err := c.FormFile("file")
	if err != nil {
		if isRequestBodyTooLarge(err) {
			// Optional: attach internal sentinel for later detection/testing
			// err = fmt.Errorf("%w", ErrBodyTooLarge)
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "File size exceeds 10 MiB limit"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid file upload or request"})
		return
	}

	// Save the file to a temporary location
	filename := utils.SanitizeFilename(file.Filename)
	savePath := filepath.Join("../media", filename)
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		log.Printf("failed to save uploaded file %s: %v", filename, err)
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

	metadata, err := utils.ExtractMetadata(savePath)
	if err != nil {
		log.Printf("failed to extract metadata for file %s: %v", filename, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to extract metadata", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, metadata)
}
