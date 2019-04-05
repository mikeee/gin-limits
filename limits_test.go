package limits

import (
	"github.com/stretchr/testify/assert"
	"time"
	"testing"

	"github.com/gin-gonic/gin"
)

// TestLimit expects a specified limit of 1 to allow one request.
func TestLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(limits.MaxConnections(1))

	r.GET("/", func(*gin.Context) {
		time.Sleep(1 * time.Second)
	})

	// Making a request
	w := performRequest(r, "GET", "/")
	assert.Equal(t, 200, w.Code)
}
