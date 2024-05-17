package routes

import (
    "github.com/labstack/echo/v4"
    "assesment-deals/handlers"
)

func InitRoutes(e *echo.Echo) {
    e.POST("/signup", handlers.SignUp)
    e.POST("/login", handlers.Login)
}
