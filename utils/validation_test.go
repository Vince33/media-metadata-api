package utils

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsValidMimeType(t *testing.T) {
	// Ensure testdata directory exists
	err := os.MkdirAll("utils/testdata", 0755)
	require.NoError(t, err)

	testFile := "utils/testdata/sample_video.mp4"

	// Safer FFmpeg command with full output logging
	cmd := exec.Command("ffmpeg", "-y",
		"-f", "lavfi", "-i", "testsrc=duration=1:size=128x128:rate=24",
		"-c:v", "libx264", "-t", "1", "-pix_fmt", "yuv420p",
		testFile,
	)
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "failed to generate test video: %s", string(output))

	// Run MIME type validations
	valid := IsValidMimeType(testFile, []string{"video/mp4"})
	require.True(t, valid, "expected valid MIME type")

	invalid := IsValidMimeType(testFile, []string{"image/png"})
	require.False(t, invalid, "expected invalid MIME type")

	missing := IsValidMimeType("does_not_exist.mp4", []string{"video/mp4"})
	require.False(t, missing, "expected false for nonexistent file")

	t.Cleanup(func() {
		os.Remove(testFile)
		os.Remove("utils/testdata")
	})
}
