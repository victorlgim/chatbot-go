package main

import (
    "net/http"
)

func handleStatic(c *gin.Context) {
    http.ServeFile(c.Writer, c.Request, "web/index.html")
}
