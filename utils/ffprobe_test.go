package utils

import (
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestExtractMetadata_InvalidFile checks that it fails gracefully on a nonexistent file.
func TestExtractMetadata_InvalidFile(t *testing.T) {
	_, err := ExtractMetadata("nonexistent.mp4")
	require.Error(t, err, "expected error for invalid file")
}

// TestExtractMetadata_ValidFile runs ffprobe on a real file if available.
func TestExtractMetadata_ValidFile(t *testing.T) {
	// Use a temporary directory for the test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "sample_for_utils_test.mp4")

	// Generate the sample video file with FFmpeg
	cmd := exec.Command("ffmpeg", "-y", "-f", "lavfi", "-i", "testsrc=duration=1:size=128x128:rate=24", testFile)
	err := cmd.Run()
	require.NoError(t, err, "ffmpeg video generation failed")

	metadata, err := ExtractMetadata(testFile)
	require.NoError(t, err, "metadata extraction should not fail on valid video")
	require.NotNil(t, metadata, "metadata should not be nil")
	require.NotEmpty(t, metadata.Format.Duration, "duration should be populated")
}
