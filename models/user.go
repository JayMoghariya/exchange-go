package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username   string  `gorm:"uniqueIndex;not null"`
    Password   string  `gorm:"not null"` // Store hashed passwords!
    RoleID     uint
    Role       Role
}

type Role struct {
    gorm.Model
    Name        string       `gorm:"uniqueIndex;not null"`
    Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
    gorm.Model
    Name string `gorm:"uniqueIndex;not null"`
}