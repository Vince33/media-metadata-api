package utils

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractMetadata_InvalidFile(t *testing.T) {
	_, err := ExtractMetadata("nonexistent.mp4")
	require.Error(t, err, "expected error for invalid file")
}

func TestExtractMetadata_ValidFile(t *testing.T) {
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "sample_for_utils_test.mp4")

	GenerateTestVideo(t, testFile, "128x128", "1")

	metadata, err := ExtractMetadata(testFile)
	require.NoError(t, err, "metadata extraction should not fail on valid video")
	require.NotNil(t, metadata, "metadata should not be nil")
	require.NotEmpty(t, metadata.Format.Duration, "duration should be populated")
}
