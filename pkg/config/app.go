package config

import (
    "fmt"
    "log"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
    var db *gorm.DB
    var err error
    dsn := "root:root@tcp(mysql:3306)/books?charset=utf8mb4&parseTime=True&loc=Local"

    // Retry logic
    for i := 0; i < 5; i++ {
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        log.Printf("Database connection failed: %v. Retrying in 5 seconds...", err)
        time.Sleep(10 * time.Second) // Initial delay before first connection attempt

    }

    if err != nil {
        log.Fatalf("Failed to connect to database after retries: %v", err)
    }

    Db = db
    fmt.Println("Connected to database.")
}
