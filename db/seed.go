package db

import (
    "trading-system-go/models"
)

func SeedRolesAndPermissions() {
    // Example permissions
    perms := []models.Permission{
        {Name: "manage_users"},
        {Name: "place_order"},
        {Name: "view_orders"},
    }
    for _, p := range perms {
        DB.FirstOrCreate(&p, models.Permission{Name: p.Name})
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