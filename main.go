package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

var (
    upgrader = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
    clients = make(map[*websocket.Conn]bool)
)

func main() {
    r := gin.Default()

    r.GET("/ws", func(c *gin.Context) {
        serveWebSocket(c.Writer, c.Request)
    })

    r.GET("/", func(c *gin.Context) {
        http.ServeFile(c.Writer, c.Request, "web/index.html")
    })

    r.Run(":8080")
}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return
    }

    clients[conn] = true

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        
        for client := range clients {
            if err := client.WriteMessage(messageType, p); err != nil {
                log.Println(err)
                return
            }
        }
    }
}
