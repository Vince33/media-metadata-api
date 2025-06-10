package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UploadValidationMiddleware(maxSize int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxSize)

		contentType := c.GetHeader("Content-Type")
		if contentType == "" || !strings.HasPrefix(contentType, "multipart/form-data") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be multipart/form-data"})
			return
		}

		c.Next()
	}
}
