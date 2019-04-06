package limits

import (
	"github.com/gin-gonic/gin"
)

// MaxConnections creates an artificial limit on connections.
func MaxConnections(count int) gin.HandlerFunc {
	semaphore := make(chan struct{}, count)
	acquire := func() {
		semaphore <- struct{}{}
	}
	relase := func() {
		<-semaphore
	}
	return func(context *gin.Context) {
		acquire()      // before request
		defer relase() // after request
		context.Next()
	}
}
