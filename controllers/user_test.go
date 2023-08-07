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

var token string

func TestMain(m *testing.M) {
	godotenv.Load("../.env")
	db.Init()
	migration.Init()

	code := m.Run()

	models.DeleteUserByName("user")
	os.Exit(code)
}

func SendUserPass(t *testing.T, path string, handler func(c echo.Context) error) {
	e := echo.New()

	data, _ := json.Marshal(map[string]any{
		"username": "user",
		"password": "pass",
	})
	req, _ := http.NewRequest(http.MethodPost, path, bytes.NewBuffer(data))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, handler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		var resData map[string]string
		json.NewDecoder(rec.Body).Decode(&resData)
		tkn, ok := resData["bearer"]
		token = tkn
		assert.True(t, ok)
	}
}

func TestRegister(t *testing.T) {
	SendUserPass(t, "/register", Register)
}

func TestLogin(t *testing.T) {
	SendUserPass(t, "/login", Login)
}

func TestRefresh(t *testing.T) {

}
