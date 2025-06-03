package utils

import (
	"encoding/json"
	"os/exec"
)

type FFProbeFormat struct {
	Filename       string            `json:"filename"`
	FormatName     string            `json:"format_name"`
	FormatLongName string            `json:"format_long_name"`
	Duration       string            `json:"duration"`
	Size           string            `json:"size"`
	BitRate        string            `json:"bit_rate"`
	Tags           map[string]string `json:"tags"`
}

type FFProbeStream struct {
	CodecName     string            `json:"codec_name"`
	CodecType     string            `json:"codec_type"`
	Width         int               `json:"width,omitempty"`
	Height        int               `json:"height,omitempty"`
	BitRate       string            `json:"bit_rate,omitempty"`
	Duration      string            `json:"duration,omitempty"`
	Channels      int               `json:"channels,omitempty"`
	SampleRate    string            `json:"sample_rate,omitempty"`
	ChannelLayout string            `json:"channel_layout,omitempty"`
	Tags          map[string]string `json:"tags"`
}

type FFProbeOutput struct {
	Format  FFProbeFormat   `json:"format"`
	Streams []FFProbeStream `json:"streams"`
}

// ExtractMetadata runs ffprobe and parses JSON output
func ExtractMetadata(filePath string) (*FFProbeOutput, error) {
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

	var metadata FFProbeOutput
	if err := json.Unmarshal(output, &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}
