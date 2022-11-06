package handler

import (
	_ "Api/config"
	"Api/pkg/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	testUser, _ = json.Marshal(map[string]interface{}{"Email": "test@mail.com", "Password": "test", "Name": "test user"})
	authHandler = NewAuthHandler(service.BcryptHasher{Cost: 5}, service.Jwt{}, service.Oauth{})
)

func TestGoogleSignIn(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sign-in/google")

	if assert.NoError(t, authHandler.GoogleSignIn(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignUp(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testUser))
	req.Header.Set(echo.HeaderContentType, "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sign-up")

	if assert.NoError(t, authHandler.SignUp(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestSignIn(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(testUser))
	req.Header.Set(echo.HeaderContentType, "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/sign-in")

	if assert.NoError(t, authHandler.SignIn(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
