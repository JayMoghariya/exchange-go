package engine

import (
    "trading-system-go/db"
    "trading-system-go/notifier"
    "trading-system-go/models"
)

type persistencePayload struct {
    Order  *models.Order
    Trades []models.Trade
}

var persistenceChan = make(chan persistencePayload, 100)

func StartPersistenceWorker() {
    go func() {
        for payload := range persistenceChan {
            db.DB.Create(payload.Order)
            for _, trade := range payload.Trades {
                db.DB.Create(&trade)
                notifier.BroadcastTrade(trade)
            }
        }
    }()
}

func SendToPersistence(order *models.Order, trades []models.Trade) {
    persistenceChan <- persistencePayload{Order: order, Trades: trades}
}
