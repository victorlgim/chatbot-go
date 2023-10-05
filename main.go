package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/ws", handleWebSocket)
    r.GET("/", handleStatic)

    r.Run(":8081")
}
