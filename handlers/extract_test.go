package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

//Summary:
//This test simulates uploading a valid video file to your /extract endpoint
// and checks that the handler responds with HTTP 200 OK. It uses Go’s
// standard testing tools and Gin’s test mode to run everything in memory—no real HTTP server is started.

func TestExtractHandler_ValidFile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/extract", ExtractHandler)

	// Use a temporary directory for the test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "sample_for_extract_test.mp4")

	// Generate the sample video file with FFmpeg
	cmd := exec.Command(
		"ffmpeg", "-y", "-f", "lavfi", "-i", "testsrc=duration=1:size=1280x720:rate=24", testFile,
	)
	output, err := cmd.CombinedOutput()
	require.NoError(t, err, "ffmpeg video generation failed: %s", string(output))

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	file, err := os.Open(testFile)
	require.NoError(t, err)
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(testFile))
	require.NoError(t, err)
	_, err = io.Copy(part, file)
	require.NoError(t, err)
	writer.Close()

	req := httptest.NewRequest("POST", "/extract", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)
}

func TestExtractHandler_MetadataFields(t *testing.T) {
	t.Parallel()

	// Use a temporary directory for the test file
	tempDir := t.TempDir()
	testFile := filepath.Join(tempDir, "auto_generated.mp4")

	// Run FFmpeg command and capture stderr
	cmd := exec.Command(
		"ffmpeg", "-y", "-f", "lavfi", "-i", "testsrc=duration=1:size=128x128:rate=24", testFile,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("❌ FFmpeg failed to generate test video.\nExit code: %v\nOutput:\n%s", err, string(output))
	}

	// Set up Gin router and handler
	router := gin.Default()
	router.POST("/extract", ExtractHandler)

	// Open test file
	file, err := os.Open(testFile)
	require.NoError(t, err)
	defer file.Close()

	// Create multipart form with the video file
	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", filepath.Base(testFile))
	require.NoError(t, err)
	_, err = io.Copy(part, file)
	require.NoError(t, err)
	writer.Close()

	// Send request
	req := httptest.NewRequest("POST", "/extract", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusOK, resp.Code)

	// Validate returned JSON metadata
	var data map[string]interface{}
	err = json.Unmarshal(resp.Body.Bytes(), &data)
	require.NoError(t, err)

	require.Contains(t, data, "format", "missing 'format' section in metadata")

	format, ok := data["format"].(map[string]interface{})
	require.True(t, ok, "'format' field is not a map")

	require.Contains(t, format, "duration")
	require.Contains(t, format, "filename")
	require.Contains(t, format, "bit_rate")

	require.Contains(t, data, "streams", "missing 'streams' section")
	streams, ok := data["streams"].([]interface{})
	require.True(t, ok, "'streams' field is not a slice")
	require.Greater(t, len(streams), 0, "no streams returned")

	firstStream, ok := streams[0].(map[string]interface{})
	require.True(t, ok, "stream[0] is not a map")

	require.Contains(t, firstStream, "codec_name")
	require.Contains(t, firstStream, "width")
	require.Contains(t, firstStream, "height")
	require.Contains(t, firstStream, "duration")

	t.Logf("Extracted metadata: %+v", data)
}
