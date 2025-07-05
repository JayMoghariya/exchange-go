package db

import (
    "log"
    "trading-system-go/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    dsn := "host=db user=postgres password=postgres dbname=trading port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Order{}, &models.Trade{})
}
