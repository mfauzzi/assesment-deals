package main

import (
    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "log"
    "assesment-deals/config"
    "assesment-deals/routes"
)

func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize database
    config.InitDB()

    // Create Echo instance
    e := echo.New()

    // Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    // Routes
    routes.InitRoutes(e)

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}