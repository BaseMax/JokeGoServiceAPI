package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/BaseMax/JokeGoServiceAPI/db"
	"github.com/BaseMax/JokeGoServiceAPI/models"
)

func TestCreateJoke(t *testing.T) {
	reqJoke := models.JokeRequest{
		Content: "Joke content 1",
		Author:  FAKE_USER,
	}
	expectedJoke := models.JokeRequest{
		ID:      1,
		Content: "Joke content 1",
		Author:  FAKE_USER,
		Rating:  0,
	}
	var resJoke models.JokeRequest

	e := echo.New()
	data, _ := json.Marshal(reqJoke)
	req := httptest.NewRequest(http.MethodPost, "/jokes", bytes.NewBuffer(data))
	rec := httptest.NewRecorder()

	if assert.NoError(t, CreateJoke(e.NewContext(req, rec))) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&resJoke)
		assert.Equal(t, expectedJoke, resJoke)
	}

	reqJoke.Author = "WrongUser"
	data, _ = json.Marshal(reqJoke)
	req, _ = http.NewRequest(http.MethodPost, "/jokes", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()

	if err := CreateJoke(e.NewContext(req, rec)); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}

	data, _ = json.Marshal(map[string]any{"content": 1, "author": nil})
	req, _ = http.NewRequest(http.MethodPost, "/jokes", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	if err := CreateJoke(e.NewContext(req, rec)); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req, _ = http.NewRequest(http.MethodPost, "/jokes", nil)
	rec = httptest.NewRecorder()
	if err := CreateJoke(e.NewContext(req, rec)); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}
}

func TestGetJoke(t *testing.T) {
	expectedJoke := models.JokeRequest{
		ID:      1,
		Content: "Joke content 1",
		Author:  FAKE_USER,
		Rating:  0,
	}
	var resJoke models.JokeRequest

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1")

	if assert.NoError(t, GetJoke(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&resJoke)
		assert.Equal(t, expectedJoke, resJoke)
	}

	req = httptest.NewRequest(http.MethodGet, "/jokes/100", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	if err := GetJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodGet, "/jokes/100", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("100")

	if err := GetJoke(c); assert.Error(t, err) {
		assert.Equal(t, err, echo.ErrNotFound)
	}
}

func TestGetAllJokes(t *testing.T) {
	limit := 26
	page := 6
	total := 200
	target := fmt.Sprint("/jokes?limit=", limit, "&page=", page, "&sort=latest")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	if err := GetAllJokes(c); assert.Error(t, err) {
		assert.Equal(t, err, echo.ErrNotFound)
	}

	var jokes []any
	for i := 2; i <= total; i++ {
		joke := models.JokeRequest{
			Content: fmt.Sprint("Joke content ", i),
			Author:  FAKE_USER,
		}
		assert.NoError(t, models.CreateJoke(&joke))
	}

	expecteResult := map[string]any{
		"total": float64(total),
		"limit": float64(limit),
		"page":  float64(page),
		"sort":  "latest",
		"jokes": jokes,
	}
	var actualResult map[string]any

	start := total - ((page - 1) * limit)
	end := start - limit
	for i := start; i > end; i-- {
		jokes = append(jokes, map[string]any{
			"author": FAKE_USER, "content": fmt.Sprint("Joke content ", i),
			"id": float64(i), "rating": float64(0),
		})
	}
	expecteResult["jokes"] = jokes

	req = httptest.NewRequest(http.MethodGet, target, nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	if assert.NoError(t, GetAllJokes(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		assert.Equal(t, expecteResult, actualResult)
	}
}

func TestGetRandomJoke(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/random", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetRandomJoke(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var res map[string]any
		json.NewDecoder(rec.Body).Decode(&res)
		assert.NotEmpty(t, res)
	}
}

func TestRateJokes(t *testing.T) {
	limit := 5

	var jokes []any
	for i := limit; i > 0; i-- {
		rate := i * 2
		joke_id := i

		jokes = append(jokes, map[string]any{
			"author": FAKE_USER, "content": fmt.Sprint("Joke content ", i),
			"id": float64(joke_id), "rating": float64(rate),
		})
	}
	expectedResult := map[string]any{
		"limit": float64(limit),
		"jokes": jokes,
	}
	var actualResult map[string]any

	e := echo.New()
	for i := 1; i <= 5; i++ {
		rate := i * 2
		joke_id := i

		data, _ := json.Marshal(map[string]any{"rating": rate})
		target := fmt.Sprint("/jokes/", joke_id, "/rating")
		req := httptest.NewRequest(http.MethodPost, target, bytes.NewBuffer(data))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetParamNames("joke_id")
		c.SetParamValues(fmt.Sprint(joke_id))

		assert.NoError(t, RateJoke(c))
	}

	target := fmt.Sprint("/jokes/top-rated?limit=", limit)
	req := httptest.NewRequest(http.MethodGet, target, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, GetTopRatedJoke(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		assert.Equal(t, expectedResult, actualResult)
	}

	data, _ := json.Marshal(map[string]any{"rating": 10})
	req = httptest.NewRequest(http.MethodPost, "/jokes/badid/rating", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("badid")
	if err := RateJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodPost, "/jokes/1/rating", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1")
	if err := RateJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodPost, "/jokes/1500/rating", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1500")
	if err := RateJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestGetJokeByAuthor(t *testing.T) {
	var expectedResult []any
	author := "author"
	models.RegisterUser(&models.User{Username: author, Password: "pass"})

	for i := 0; i < 5; i++ {
		content := fmt.Sprint(author, ": New joke ", i)
		assert.NoError(t, models.CreateJoke(&models.JokeRequest{Content: content, Author: author}))
		expectedResult = append(expectedResult, map[string]any{
			"id":      0,
			"content": content,
			"author":  author,
			"rating":  float64(0),
		})
	}
	var actualResult []any

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/authors/"+author, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("author_name")
	c.SetParamValues(author)

	if assert.NoError(t, GetJokeByAuthor(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		for _, joke := range actualResult {
			joke.(map[string]any)["id"] = 0
		}
		assert.Equal(t, expectedResult, actualResult)
	}

	models.DeleteUserByName(author)
}

func TestEditJoke(t *testing.T) {
	author := "author"
	models.RegisterUser(&models.User{Username: author, Password: "pass"})

	reqJoke := models.JokeRequest{
		Content: "Updated joke",
		Author:  author,
	}
	expectedResult := models.JokeRequest{
		ID:      1,
		Content: "Updated joke",
		Author:  author,
		Rating:  0,
	}
	var actualResult models.JokeRequest

	e := echo.New()
	data, _ := json.Marshal(reqJoke)
	req := httptest.NewRequest(http.MethodPut, "/jokes/1", bytes.NewBuffer(data))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1")
	if assert.NoError(t, EditJoke(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		json.NewDecoder(rec.Body).Decode(&actualResult)
		assert.Equal(t, expectedResult, actualResult)
	}

	req = httptest.NewRequest(http.MethodPut, "/jokes/badid", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("badid")
	if err := EditJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodPut, "/jokes/1500", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1500")
	if err := EditJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodPut, "/jokes/15000", bytes.NewBuffer(data))
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("15000")
	if err := EditJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}

	models.DeleteUserByName(author)
}

func TestDeleteJoke(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/jokes/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1")

	if assert.NoError(t, DeleteJoke(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		_, err := models.FetchAJoke(1)
		assert.Error(t, err)
	}

	req = httptest.NewRequest(http.MethodDelete, "/jokes/badid", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("badid")
	if err := DeleteJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrBadRequest, err)
	}

	req = httptest.NewRequest(http.MethodDelete, "/jokes/1500", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("joke_id")
	c.SetParamValues("1500")
	if err := DeleteJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestTruncateJokes(t *testing.T) {
	db.TruncateTable("comments", "jokes", "users")
}

func TestGetRandomJokeNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/random", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := GetRandomJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestRateJokesNotFount(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/top-rated?limit=10", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := GetTopRatedJoke(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}

func TestGetJokeByAuthorNoyFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/jokes/authors/noauthor", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("author_name")
	c.SetParamValues("noauthor")

	if err := GetJokeByAuthor(c); assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}
}
