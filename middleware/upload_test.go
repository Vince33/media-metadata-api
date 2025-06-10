package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestUploadValidationMiddleware_ContentTypeEnforced(t *testing.T) {
	router := gin.New()
	router.POST("/extract",
		UploadValidationMiddleware(10<<20), // 10 MB limit
		func(c *gin.Context) {
			c.String(http.StatusOK, "passed")
		},
	)

	req := httptest.NewRequest("POST", "/extract", strings.NewReader("not multipart"))
	req.Header.Set("Content-Type", "application/json") // invalid content type

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusBadRequest, resp.Code)
	require.Contains(t, resp.Body.String(), "Content-Type must be multipart/form-data")
}

func TestUploadValidationMiddleware_SizeLimit(t *testing.T) {
	router := gin.New()
	router.POST("/extract",
		UploadValidationMiddleware(5), // 5 bytes
		func(c *gin.Context) {
			// Force the form to be parsed — triggers size limit enforcement
			if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
				c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "too big"})
				return
			}
			c.String(http.StatusOK, "passed")
		},
	)

	// Craft a tiny but still-too-big multipart request
	var buf bytes.Buffer
	buf.WriteString("--boundary\r\n")
	buf.WriteString("Content-Disposition: form-data; name=\"file\"; filename=\"a.txt\"\r\n")
	buf.WriteString("Content-Type: text/plain\r\n\r\n")
	buf.WriteString("123456") // 6 bytes > 5
	buf.WriteString("\r\n--boundary--\r\n")

	req := httptest.NewRequest("POST", "/extract", &buf)
	req.Header.Set("Content-Type", "multipart/form-data; boundary=boundary")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	require.Equal(t, http.StatusRequestEntityTooLarge, resp.Code)
}
