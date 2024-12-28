package config

import (
    "fmt"
    "log"
    "os"
    "time"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/joho/godotenv"
)

var Db *gorm.DB

func Connect() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Printf("Error loading .env file: %v", err)
    }

    // Fetch environment variables
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    // Construct the DSN (Data Source Name)
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, password, host, port, dbname)

    // Retry logic for database connection
    for i := 0; i < 5; i++ {
        Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
        if err == nil {
            break
        }
        log.Printf("Database connection failed: %v. Retrying in 5 seconds...", err)
        time.Sleep(5 * time.Second)
    }

    if err != nil {
        log.Fatalf("Failed to connect to database after retries: %v", err)
    }

    fmt.Println("Connected to the MySQL database.")
}
