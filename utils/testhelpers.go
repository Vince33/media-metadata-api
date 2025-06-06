// utils/testhelpers.go
package utils

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

// GenerateTestVideo creates a synthetic video file using FFmpeg for testing purposes.
func GenerateTestVideo(t *testing.T, outputPath, size, duration string) {
	t.Helper()
	cmd := exec.Command(
		"ffmpeg", "-y",
		"-f", "lavfi",
		"-i", "testsrc=duration="+duration+":size="+size+":rate=24",
		outputPath,
	)
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "❌ FFmpeg video generation failed: %s", string(output))
}
