// limits is a gin middleware handler that allows for connection throttling.
package limits

import (
	"github.com/gin-gonic/gin"
)

// MaxConnections creates an artificial limit on connections.
func MaxConnections(limit int) gin.HandlerFunc {
	semaphore := make(chan struct{}, limit)
	acquire := func() {
		semaphore <- struct{}{}
	}
	release := func() {
		<-semaphore
	}
	return func(context *gin.Context) {
		acquire()
		defer release()
		context.Next()
	}
}
