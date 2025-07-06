package db

import (
	"log"
	"os"
	"trading-system-go/models"
	"trading-system-go/utils"
)

func SeedRolesAndPermissions() {
	// Example permissions
	permNames := []string{"manage_users", "place_order", "view_orders"}
	var perms []models.Permission
	for _, name := range permNames {
		var p models.Permission
		DB.FirstOrCreate(&p, models.Permission{Name: name})
		perms = append(perms, p)
	}

	// Admin role
	var adminRole models.Role
	DB.FirstOrCreate(&adminRole, models.Role{Name: "admin"})
	DB.Model(&adminRole).Association("Permissions").Replace(perms)

	// Trader role
	var traderRole models.Role
	DB.FirstOrCreate(&traderRole, models.Role{Name: "trader"})
	DB.Model(&traderRole).Association("Permissions").Replace([]models.Permission{perms[1], perms[2]})
}

func SeedAdminUser() {
	var admin_password string = os.Getenv("ADMIN_PASSWORD");
	var admin_username string = os.Getenv("ADMIN_USERNAME");
	if admin_password == "" || admin_username == "" {
		log.Fatal("ADMIN_PASSWORD and ADMIN_USERNAME must be set in the environment variables")
	}
	var adminRole models.Role
	if err := DB.Where("name = ?", admin_username).First(&adminRole).Error; err != nil {
		log.Fatal("Admin role not found. Seed roles first.")
	}

	// Check if admin user already exists
	var user models.User
	if err := DB.Where("username = ?", admin_username).First(&user).Error; err == nil {
		log.Println("Admin user already exists, skipping creation.")
		return // Admin user already exists
	}

	// Check if the password is strong enough
	if !utils.IsStrongPassword(admin_password) {
		log.Fatal("Admin password must be at least 8 characters long, contain at least one uppercase letter, one lowercase letter, one number, and one special character")
	}
	
	// Hash the password
	hashedPassword, err := utils.HashPassword(admin_password)
	if err != nil {
		log.Fatalf("Failed to hash password: %v", err)
	}
	
	adminUser := models.User{
		Username: "admin",
		Password: string(hashedPassword),
		RoleID:   adminRole.ID,
	}
	DB.Create(&adminUser)
}
