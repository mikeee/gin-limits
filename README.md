# gin-limits

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Build Status](https://travis-ci.org/mikeee/gin-limits.svg?branch=master)](https://travis-ci.org/mikeee/gin-limits)
[![Go Report Card](https://goreportcard.com/badge/github.com/mikeee/gin-limits)](https://goreportcard.com/report/github.com/mikeee/gin-limits)

Gin framework middleware that currently implements connections throttling through the use of `limits.MaxConnections()`

## Usage/Example

Simply use the middleware with Gin.

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

## Licence

This package is released under the MIT License as available in the [LICENSE](LICENSE) file
