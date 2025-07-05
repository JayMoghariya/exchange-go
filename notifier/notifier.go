package notifier

import (
    "github.com/gorilla/websocket"
    "trading-system-go/models"
)

var clients = make(map[*websocket.Conn]bool)

func RegisterClient(ws *websocket.Conn) {
    clients[ws] = true
}

func RemoveClient(ws *websocket.Conn) {
    delete(clients, ws)
}

func BroadcastTrade(trade models.Trade) {
    for client := range clients {
        err := client.WriteJSON(trade)
        if err != nil {
            client.Close()
            RemoveClient(client)
        }
    }
}
