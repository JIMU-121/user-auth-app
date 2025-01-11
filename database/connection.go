// database/connection.go

package database

import (
    "fmt"
    "log"
    "user-auth-app/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("./userauth.db"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

    // Auto-migrate the User table
    DB.AutoMigrate(&models.User{})
    fmt.Println("Database connected successfully")
}
