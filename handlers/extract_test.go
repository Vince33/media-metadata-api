package handlers

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
)

//Summary:
//This test simulates uploading a valid video file to your /extract endpoint
// and checks that the handler responds with HTTP 200 OK. It uses Go’s
// standard testing tools and Gin’s test mode to run everything in memory—no real HTTP server is started.

func TestExtractHandler_ValidFile(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/extract", ExtractHandler)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	filePath := "../testdata/SampleVideo_1280x720_1mb.mp4" // Path to a sample video file, make sure this file exists in your test environment. May have to add to repo undecided on how to handle for the time being.
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/extract", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.Code)
	}
}
