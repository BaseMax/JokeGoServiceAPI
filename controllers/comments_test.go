package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/BaseMax/JokeGoServiceAPI/models"
)

func TestCreateJokeComment(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        "1",
		Issuer:    FAKE_USER,
		ExpiresAt: EXPTIME,
	})
	bearer, _ := token.SignedString([]byte(os.Getenv("JWT_KET")))

	expectedResult := models.CommentRequest{ID: 1, Content: "Comment", Author: FAKE_USER}
	var actualResult models.CommentRequest

	models.CreateJoke(&models.JokeRequest{Content: "Joke content", Author: FAKE_USER})

	e := echo.New()
	data, _ := json.Marshal(expectedResult)
	req := httptest.NewRequest(http.MethodPost, "/jokes/1/comments", bytes.NewBuffer(data))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1")
	req.Header.Set("Authorization", "Bearer "+bearer)
	if assert.NoError(t, CreateJokeComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		assert.Equal(t, expectedResult, actualResult)
	}

	req = httptest.NewRequest(http.MethodPost, "/jokes/1/comments", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1")
	req.Header.Set("Authorization", "Bearer "+bearer)
	if err := CreateJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodPost, "/jokes/badid/comments", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("badid")
	req.Header.Set("Authorization", "Bearer "+bearer)
	if err := CreateJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodPost, "/jokes/100/comments", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("100")
	req.Header.Set("Authorization", "Bearer "+bearer)
	if err := CreateJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestGetJokeComment(t *testing.T) {
	expectedResult := models.CommentRequest{ID: 1, Content: "Comment", Author: FAKE_USER}
	var actualResult models.CommentRequest

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/comments/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("1")
	if assert.NoError(t, GetJokeComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		assert.Equal(t, expectedResult, actualResult)
	}

	req = httptest.NewRequest(http.MethodGet, "/jokes/badid/comments", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("badid")
	if err := GetJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodGet, "/jokes/1500/comments", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("1500")
	if err := GetJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestGetJokeComments(t *testing.T) {
	var expectedResult []models.CommentRequest
	var actualResult []models.CommentRequest
	expectedResult = append(expectedResult, models.CommentRequest{ID: 1, Content: "Comment", Author: FAKE_USER})

	for i := 2; i <= 10; i++ {
		comment := models.CommentRequest{ID: uint(i), Content: fmt.Sprint("Comment ", i), Author: FAKE_USER}
		expectedResult = append(expectedResult, comment)
		assert.NoError(t, models.CreateComment(1, &comment))
	}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/1/comments", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1")
	if assert.NoError(t, GetJokeComments(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		assert.Equal(t, expectedResult, actualResult)
	}

	req = httptest.NewRequest(http.MethodGet, "/jokes/badid/comments", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("badid")
	if err := GetJokeComments(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodGet, "/jokes/1500/comments", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1500")
	if err := GetJokeComments(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestEditJokeComment(t *testing.T) {
	expectedResult := models.CommentRequest{ID: 1, Content: "Updated comment", Author: FAKE_USER}
	var actualResult models.CommentRequest

	e := echo.New()
	data, _ := json.Marshal(expectedResult)
	req := httptest.NewRequest(http.MethodPut, "/jokes/comments/1", bytes.NewBuffer(data))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("1")
	if assert.NoError(t, EditJokeComment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		assert.Equal(t, expectedResult, actualResult)
	}

	req = httptest.NewRequest(http.MethodPut, "/jokes/comments/1500", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("1500")
	if err := EditJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}

	req = httptest.NewRequest(http.MethodPut, "/jokes/comments/badid", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("badid")
	if err := EditJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	data, _ = json.Marshal(map[string]any{"content": 1})
	req = httptest.NewRequest(http.MethodPut, "/jokes/comments/1", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("1")
	if err := EditJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}
}

func TestDeleteJokeComment(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/jokes/comments/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("1")
	if assert.NoError(t, DeleteJokeComment(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		assert.Empty(t, rec.Body.String())
	}

	req = httptest.NewRequest(http.MethodPut, "/jokes/comments/badid", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("badid")
	if err := DeleteJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodPut, "/jokes/comments/1500", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("comment_id")
	c.SetParamValues("1500")
	if err := DeleteJokeComment(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestCleanUp(t *testing.T) {
	db.TruncateTable("comments", "jokes")
}
