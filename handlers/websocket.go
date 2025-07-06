package handlers

import (
	"log"
	"net/http"
	"trading-system-go/notifier"
	"trading-system-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(c *gin.Context) {
	tokenString := c.Query("token")
	if tokenString == "" {
		tokenString = c.GetHeader("Sec-WebSocket-Protocol")
	}
	claims, err := utils.ParseJWT(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		return
	}
	username, _ := claims["username"].(string)
	userIDFloat, _ := claims["user_id"].(float64) // JWT numbers are float64
	userID := uint(userIDFloat)

	log.Printf("WebSocket connection from user: %s (ID: %d) to stream", username, userID)

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	notifier.RegisterClient(ws)
	defer func() {
		notifier.RemoveClient(ws)
		ws.Close()
	}()

	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			notifier.RemoveClient(ws)
			break
		}
	}
}
