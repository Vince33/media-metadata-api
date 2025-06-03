package utils

import (
	"encoding/json"
	"os/exec"
)

// ExtractMetadata runs ffprobe and parses JSON output
func ExtractMetadata(filePath string) (map[string]interface{}, error) {
	cmd := exec.Command("ffprobe",
		"-v", "quiet",
		"-print_format", "json",
		"-show_format",
		"-show_streams",
		filePath)

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var metadata map[string]interface{}
	if err := json.Unmarshal(output, &metadata); err != nil {
		return nil, err
	}

	return metadata, nil
}
