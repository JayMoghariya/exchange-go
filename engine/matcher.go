
package engine

import (
    "sync"
    "trading-system-go/models"
)

type OrderBook struct {
    BuyOrders  []*models.Order
    SellOrders []*models.Order
    mu         sync.Mutex
}

var orderBook = OrderBook{
    BuyOrders:  make([]*models.Order, 0),
    SellOrders: make([]*models.Order, 0),
}

func PlaceOrder(newOrder *models.Order) (trades []models.Trade) {
    orderBook.mu.Lock()
    defer orderBook.mu.Unlock()

    if newOrder.Type == models.BUY {
        trades = matchBuyOrder(newOrder)
    } else {
        trades = matchSellOrder(newOrder)
    }

    if newOrder.Quantity > 0 {
        if newOrder.Type == models.BUY {
            orderBook.BuyOrders = append(orderBook.BuyOrders, newOrder)
        } else {
            orderBook.SellOrders = append(orderBook.SellOrders, newOrder)
        }
    } else {
        newOrder.IsFilled = true
    }

    SendToPersistence(newOrder, trades)
    return trades
}

func matchBuyOrder(buy *models.Order) (trades []models.Trade) {
    i := 0
    for i < len(orderBook.SellOrders) && buy.Quantity > 0 {
        sell := orderBook.SellOrders[i]
        if sell.Price > buy.Price {
            break
        }
        matchedQty := min(buy.Quantity, sell.Quantity)
        trade := models.Trade{
            BuyOrderID:  buy.ID,
            SellOrderID: sell.ID,
            Price:       sell.Price,
            Quantity:    matchedQty,
        }
        trades = append(trades, trade)

        buy.Quantity -= matchedQty
        sell.Quantity -= matchedQty

        if sell.Quantity <= 0 {
            sell.IsFilled = true
            orderBook.SellOrders = append(orderBook.SellOrders[:i], orderBook.SellOrders[i+1:]...)
        } else {
            i++
        }
    }
    return
}

func matchSellOrder(sell *models.Order) (trades []models.Trade) {
    i := 0
    for i < len(orderBook.BuyOrders) && sell.Quantity > 0 {
        buy := orderBook.BuyOrders[i]
        if buy.Price < sell.Price {
            break
        }
        matchedQty := min(sell.Quantity, buy.Quantity)
        trade := models.Trade{
            BuyOrderID:  buy.ID,
            SellOrderID: sell.ID,
            Price:       buy.Price,
            Quantity:    matchedQty,
        }
        trades = append(trades, trade)

        sell.Quantity -= matchedQty
        buy.Quantity -= matchedQty

        if buy.Quantity <= 0 {
            buy.IsFilled = true
            orderBook.BuyOrders = append(orderBook.BuyOrders[:i], orderBook.BuyOrders[i+1:]...)
        } else {
            i++
        }
    }
    return
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}
