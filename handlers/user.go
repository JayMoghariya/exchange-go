package handlers

import (
    "net/http"
    "trading-system-go/db"
    "trading-system-go/models"
    "github.com/gin-gonic/gin"
	"trading-system-go/utils"
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

	// Check if the password is strong enough
	if !utils.IsStrongPassword(input.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character"})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	// Check if the user already exists
	var existingUser models.User
	if err := db.DB.Where("username = ?", input.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

    user := models.User{
        Username: input.Username,
        Password: hashedPassword,
        RoleID:   role.ID,
    }

	// Create the user in the database
    if err := db.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
        "id":         user.ID,
        "username":   user.Username,
        "role":       role.Name,
        "created_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
        "updated_at": user.UpdatedAt.Format("2006-01-02 15:04:05"),
    },
    "message": "User registered successfully!!",
	})
}