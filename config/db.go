package config

import (
    "fmt"
    "log"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbHost := os.Getenv("DB_HOST")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPass, dbHost, dbName)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Failed to connect to database:", err)
    }

    log.Println("✅ Database connected successfully")
}

