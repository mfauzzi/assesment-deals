package handlers

import (
    "database/sql"
    "net/http"
    "os"
    "time"
    "assesment-deals/config"
    "assesment-deals/models"
    "github.com/dgrijalva/jwt-go"
    "github.com/labstack/echo/v4"
    "golang.org/x/crypto/bcrypt"
)

type JWTCustomClaims struct {
    ID int `json:"id"`
    jwt.StandardClaims
}

func SignUp(c echo.Context) error {
    username := c.FormValue("username")
    email := c.FormValue("email")
    password := c.FormValue("password")

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
    }

    err = models.CreateUser(config.DB, username, string(hashedPassword), email)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
}

func Login(c echo.Context) error {
    email := c.FormValue("email")
    password := c.FormValue("password")

    user, err := models.GetUserByEmail(config.DB, email)
    if err != nil {
        if err == sql.ErrNoRows {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
        }
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
    }

    claims := &JWTCustomClaims{
        user.ID,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Internal Server Error"})
    }

    return c.JSON(http.StatusOK, map[string]string{"token": t})
}