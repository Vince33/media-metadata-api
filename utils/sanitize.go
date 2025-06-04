package utils

import (
	"path/filepath"
	"regexp"
)

// SanitizeFilename cleans the filename by removing path elements and
// allowing only letters, numbers, dots, hyphens and underscores.
func SanitizeFilename(name string) string {
	name = filepath.Base(name)
	// Remove any character not allowed
	re := regexp.MustCompile(`[^a-zA-Z0-9._-]`)
	return re.ReplaceAllString(name, "_")
}
