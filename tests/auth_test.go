package tests

import (
	"assesment-deals/config"
	"assesment-deals/handlers"
	"assesment-deals/routes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func init() {
    config.InitDB()
}

func TestSignUp(t *testing.T) {
    e := echo.New()
    routes.InitRoutes(e)

    body := `{"username":"testuser", "email":"test@example.com", "password":"password"}`
    req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, handlers.SignUp(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Contains(t, rec.Body.String(), "User created successfully")
    }
}

func TestLogin(t *testing.T) {
    e := echo.New()
    routes.InitRoutes(e)

    body := `{"email":"test@example.com", "password":"password"}`
    req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(body))
    req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
    rec := httptest.NewRecorder()
    c := e.NewContext(req, rec)

    if assert.NoError(t, handlers.Login(c)) {
        assert.Equal(t, http.StatusOK, rec.Code)
        assert.Contains(t, rec.Body.String(), "token")
    }
}