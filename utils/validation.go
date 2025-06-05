package utils

import (
	"net/http"
	"os"
)

// IsValidMimeType checks if the provided MIME type is of a file matches one of the allowed types.
func IsValidMimeType(filePath string, allowedTypes []string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		return false
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return false
	}
	mimeType := http.DetectContentType(buffer)

	for _, t := range allowedTypes {
		if mimeType == t {
			return true
		}
	}
	return false
}
