package utils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// NOTE: Ensure GenerateTestVideo is accessible from this package, or move it to a shared testhelpers subpackage.
func TestIsValidMimeType(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "sample_video.mp4")

	GenerateTestVideo(t, testFile, "128x128", "1")

	// Run MIME type validations
	valid := IsValidMimeType(testFile, []string{"video/mp4"})
	require.True(t, valid, "expected valid MIME type")

	invalid := IsValidMimeType(testFile, []string{"image/png"})
	require.False(t, invalid, "expected invalid MIME type")

	missing := IsValidMimeType(filepath.Join(tempDir, "does_not_exist.mp4"), []string{"video/mp4"})
	require.False(t, missing, "expected false for nonexistent file")
}
