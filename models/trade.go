
package models

import "time"

type Trade struct {
    ID         uint      `gorm:"primaryKey"`
    BuyOrderID uint
    SellOrderID uint
    Price      float64
    Quantity   float64
    CreatedAt  time.Time
}
