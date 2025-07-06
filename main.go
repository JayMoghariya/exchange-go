package main

import (
	"log"
	"trading-system-go/db"
	"trading-system-go/engine"
	"trading-system-go/handlers"
    "trading-system-go/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init_db() {
	// Initialize the database and run migrations
	db.Init()

	// Seed roles and permissions
	db.SeedRolesAndPermissions()

	// Seed admin user
	db.SeedAdminUser()
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or could not be loaded")
	}

	init_db()

	engine.StartPersistenceWorker()

	r := gin.Default()
    r.POST("/orders", middleware.JWTAuthMiddleware(), handlers.PlaceOrder)
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.Login)
	r.GET("/ws", handlers.WebSocketHandler)
	r.Run(":8080")
}
