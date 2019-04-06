package limits

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// performRequest has been borrowed from the Gin repo - https://github.com/gin-gonic/gin/blob/master/routes_test.go
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

// TestLimit imposes a requests throttle of 10 per second and against two tests expects the duration to be approximately 1 second given 20 requests.
func TestLimit(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.Use(MaxConnections(10))

	r.GET("/", func(*gin.Context) {
		time.Sleep(time.Second / 2)
	})

	var wg sync.WaitGroup
	var result = true
	var test bool

	var starttime = time.Now()
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			w := performRequest(r, "GET", "/")
			if w.Code != http.StatusOK {
				result = false
			}
		}()
	}
	wg.Wait()
	var timer = time.Since(starttime)
	if (timer < time.Millisecond*1100) && (timer > time.Millisecond*900) && result == true {
		test = true
	}
	assert.Equal(t, true, test)
}
