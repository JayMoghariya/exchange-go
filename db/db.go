package db

import (
    "fmt"
    "log"
    "os"
    "trading-system-go/models"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbname, port,
    )

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Order{}, &models.Trade{})
}
func Close() {
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get database instance: %v", err)
    }
    if err := sqlDB.Close(); err != nil {
        log.Fatalf("Failed to close database connection: %v", err)
    }
}