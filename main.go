package main

import (
    "trading-system-go/db"
    "trading-system-go/engine"
    "trading-system-go/handlers"
    "github.com/gin-gonic/gin"
)

func main() {
    db.Init()
    engine.StartPersistenceWorker()

    r := gin.Default()
    r.POST("/orders", handlers.PlaceOrder)
    r.POST("/register", handlers.RegisterUser)
    r.GET("/ws", handlers.WebSocketHandler)
    r.Run(":8080")
}
