package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSanitizeFilename(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"../secret/file.mp4", "file.mp4"},
		{"../../../etc/passwd", "passwd"},
		{"video with spaces!.mp4", "video_with_spaces_.mp4"},
		{"clean-file_name.mp4", "clean-file_name.mp4"},
		{"bad@chars#here$.mp4", "bad_chars_here_.mp4"},
		{"/deep/path/to/file.mp4", "file.mp4"},
	}

	for _, tc := range tests {
		actual := SanitizeFilename(tc.input)
		require.Equal(t, tc.expected, actual, "failed on input: %s", tc.input)
	}
}
