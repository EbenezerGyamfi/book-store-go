package config

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Db *gorm.DB

func Connect() {
    var db *gorm.DB
    var err error

    // Fetch environment variables
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, password, host, port, dbname)

    // Retry logic
    for i := 0; i < 5; i++ {
        db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        log.Printf("Database connection failed: %v. Retrying in 5 seconds...", err)
        time.Sleep(5 * time.Second)
    }

    if err != nil {
        log.Fatalf("Failed to connect to database after retries: %v", err)
    }

    Db = db
    fmt.Println("Connected to database.")
}
