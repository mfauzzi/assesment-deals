package config

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "os"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    var err error
    connStr := "user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " sslmode=require"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal(err)
    }
}
