# gin-limits
Gin framework middleware that currently implements connections throttling through the use of `limits.MaxConnections()`

## Usage/Example


```go
package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/mikeee/gin-limits"
)

func main() {
    r := gin.Default()
    r.Use(limits.MaxConnections(10))

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world",
        })
	})
	
	r.Run(":8080")
}
```
