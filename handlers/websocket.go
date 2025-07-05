package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "trading-system-go/notifier"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(c *gin.Context) {
    ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        return
    }
    defer ws.Close()
    notifier.RegisterClient(ws)

    for {
        _, _, err := ws.ReadMessage()
        if err != nil {
            notifier.RemoveClient(ws)
            break
        }
    }
}
