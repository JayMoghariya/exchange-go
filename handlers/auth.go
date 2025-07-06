package handlers

import (
    "net/http"
	"trading-system-go/db"
	"trading-system-go/models"
	"trading-system-go/utils"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
    var input LoginInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := db.DB.Preload("Role").Where("username = ?", input.Username).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    if !utils.CheckPassword(user.Password, input.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := utils.GenerateJWT(user.ID, user.Username, user.Role.Name)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}