package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/BaseMax/JokeGoServiceAPI/migration"
	"github.com/BaseMax/JokeGoServiceAPI/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	token   string
	data, _ = json.Marshal(map[string]any{
		"username": "user",
		"password": "pass",
	})
	e = echo.New()
)

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	db.Init()
	migration.Init()

	code := m.Run()

	models.DeleteUserByName("user")
	os.Exit(code)
}

func TestRegister(t *testing.T) {

	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(data))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Register(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resData map[string]string
		json.NewDecoder(rec.Body).Decode(&resData)
		token = resData["bearer"]
		assert.NotEmpty(t, token)
	}

	req, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if err := Register(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrConflict, err)
	}
}

func TestLogin(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(data))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resData map[string]string
		json.NewDecoder(rec.Body).Decode(&resData)
		assert.NotEmpty(t, resData["bearer"])
	}
}

func TestRefresh(t *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(http.MethodPost, "/refresh", nil)
	rec := httptest.NewRecorder()
	req.Header.Set("Authorization", "Bearer "+token)
	c := e.NewContext(req, rec)

	if assert.NoError(t, Refresh(c)) {
		assert.Equal(t, rec.Code, http.StatusOK)
		var resData map[string]string
		json.NewDecoder(rec.Body).Decode(&resData)
		_, ok := resData["bearer"]
		assert.True(t, ok)
	}
}
