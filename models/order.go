package models

import "time"

type OrderType string

const (
	BUY  OrderType = "BUY"
	SELL OrderType = "SELL"
)

type Order struct {
	ID        uint      `gorm:"primaryKey"`
	Price     float64
	Quantity  float64
	Type      OrderType
	CreatedAt time.Time
	IsFilled  bool
}
