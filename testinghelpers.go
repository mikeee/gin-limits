package limits

import (
	"net/http"
	"net/http/httptest"
)

// performRequest has been borrowed from the Gin repo - https://github.com/gin-gonic/gin/blob/master/routes_test.go
func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}
