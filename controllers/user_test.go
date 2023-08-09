package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/BaseMax/JokeGoServiceAPI/migration"
	"github.com/BaseMax/JokeGoServiceAPI/models"
)

const (
	FAKE_USER = "user"
	FAKE_PASS = "pass"
)

var (
	token string
)

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	db.Init()
	migration.Init()

	models.RegisterUser(&models.User{Username: FAKE_USER, Password: FAKE_PASS})

	code := m.Run()

	models.DeleteUserByName("user")

	os.Exit(code)
}

func TestRegister(t *testing.T) {
	data, _ := json.Marshal(map[string]any{
		"username": "newuser",
		"password": "pass",
	})

	e := echo.New()
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

	req, _ = http.NewRequest(http.MethodPost, "/register", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if err := Register(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	data, _ = json.Marshal(map[string]any{"username": 1})
	req, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if err := Register(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	models.DeleteUserByName("newuser")
}

func TestLogin(t *testing.T) {
	data, _ := json.Marshal(map[string]any{
		"username": FAKE_USER,
		"password": FAKE_PASS,
	})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(data))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if assert.NoError(t, Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var resData map[string]string
		json.NewDecoder(rec.Body).Decode(&resData)
		assert.NotEmpty(t, resData["bearer"])
	}

	req = httptest.NewRequest(http.MethodPost, "/login", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if err := Login(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	data, _ = json.Marshal(map[string]any{"username": "nouser", "pass": "pass"})
	req = httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if err := Login(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestRefresh(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/refresh", nil)
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

	req = httptest.NewRequest(http.MethodPost, "/refresh", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if err := Refresh(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}
}
