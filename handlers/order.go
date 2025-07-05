package handlers

import (
    "net/http"
    "trading-system-go/engine"
    "trading-system-go/models"
    "github.com/gin-gonic/gin"
    "time"
    "strings"
)

type orderInput struct {
    Price    float64 `json:"price"`
    Quantity float64 `json:"quantity"`
    Type     string  `json:"type"`
}

func PlaceOrder(c *gin.Context) {
    var input orderInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    order := models.Order{
        Price:    input.Price,
        Quantity: input.Quantity,
        Type:     models.OrderType(strings.ToUpper(input.Type)),
        CreatedAt: time.Now(),
        IsFilled: false,
    }

    // Validate required fields
    if order.Price <= 0 || order.Quantity <= 0 || (order.Type != "BUY" && order.Type != "SELL") {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid fields: price, quantity, type"})
        return
    }

    trades := engine.PlaceOrder(&order)

    c.JSON(http.StatusOK, gin.H{
        "order":  order,
        "trades": trades,
    })
}
