package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true
    },
}

func handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println(err)
        return
    }

    clientIP := getClientIP(c.Request)

    clients[clientIP] = conn

    defer func() {
        delete(clients, clientIP)
        conn.Close()
    }()

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        for ip, client := range clients {
            if ip == clientIP || client == conn { 
                continue
            }

            if err := client.WriteMessage(messageType, p); err != nil {
                log.Println(err)
                return
            }
        }
    }
}
