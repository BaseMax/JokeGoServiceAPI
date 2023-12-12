package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/BaseMax/JokeGoServiceAPI/models"
	"github.com/labstack/echo/v4"
)

func CreateJoke(c echo.Context) error {
	var j models.JokeRequest
	if c.Request().Body == nil {
		return echo.ErrBadRequest
	}
	if err := json.NewDecoder(c.Request().Body).Decode(&j); err != nil {
		return echo.ErrBadRequest
	}

	if models.CreateJoke(&j) != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, j)
}

func GetJoke(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	joke, err := models.FetchAJoke(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, joke)
}

func GetAllJokes(c echo.Context) error {
	limit := 10
	page := 1
	sort := c.QueryParam("sort")

	if u, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		limit = u
	}
	if p, err := strconv.Atoi(c.QueryParam("page")); err == nil {
		page = p
	}
	if sort != "rating" {
		sort = "latest"
	}

	jokes, total, err := models.FetchAllJokes(limit, page, sort)
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, map[string]any{
		"total": total,
		"page":  page,
		"limit": limit,
		"sort":  sort,
		"jokes": jokes,
	})
}

func GetRandomJoke(c echo.Context) error {
	joke, err := models.FetchRandomJoke()
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, joke)
}

func GetTopRatedJoke(c echo.Context) error {
	limit := 10
	if u, err := strconv.Atoi(c.QueryParam("limit")); err == nil {
		limit = u
	}
	jokes, err := models.FetchTopRatedJokes(limit)
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, map[string]any{
		"limit": limit,
		"jokes": jokes,
	})
}

func GetJokeByAuthor(c echo.Context) error {
	jokes, err := models.FetchJokesByAuthor(c.Param("author_name"))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, jokes)
}

func EditJoke(c echo.Context) error {
	var joke models.JokeRequest
	id, err := strconv.Atoi(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := json.NewDecoder(c.Request().Body).Decode(&joke); err != nil {
		return echo.ErrBadRequest
	}

	joke.ID = uint(id)
	if models.UpdateJoke(joke.ID, &joke) != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, joke)
}

func DeleteJoke(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	if err := models.DeleteJokeById(uint(id)); err != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

func RateJoke(c echo.Context) error {
	var joke *models.JokeRequest
	id, err := strconv.Atoi(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := json.NewDecoder(c.Request().Body).Decode(&joke); err != nil {
		return echo.ErrBadRequest
	}

	joke, err = models.RateJoke(uint(id), joke.Rating)
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, joke)
}
