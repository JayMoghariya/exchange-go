package handlers

import (
    "net/http"
    "trading-system-go/db"
    "trading-system-go/models"
    "github.com/gin-gonic/gin"
)

type RegisterInput struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
    Role     string `json:"role" binding:"required"`
}

func RegisterUser(c *gin.Context) {
    var input RegisterInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var role models.Role
    if err := db.DB.Where("name = ?", input.Role).First(&role).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
        return
    }

    user := models.User{
        Username: input.Username,
        Password: input.Password, //TODO: Hash in production!
        RoleID:   role.ID,
    }
    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"user": user})
}