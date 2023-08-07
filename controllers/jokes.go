package controllers

import (
	"net/http"
	"strconv"

	"github.com/BaseMax/JokeGoServiceAPI/models"
	"github.com/labstack/echo/v4"
)

func CreateJoke(c echo.Context) error {
	var j models.JokeRequest
	if err := decodeBody(c, &j); err != nil {
		return echo.ErrBadRequest
	}
	err := models.CreateJoke(&j)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, j)
}

func GetJoke(c echo.Context) error {
	id, err := strToUint(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	joke, err := models.FetchAJoke(id)
	if err := dbErrorToHttp(err); err != nil {
		return err
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
		return err
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
		return err
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
	return c.JSON(http.StatusOK, map[string]any{
		"jokes": jokes,
	})
}

func EditJoke(c echo.Context) error {
	var joke models.JokeRequest
	id, err := strToUint(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := decodeBody(c, &joke); err != nil {
		return echo.ErrBadRequest
	}

	joke.ID = id
	err = models.UpdateJoke(id, &joke)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, joke)
}

func DeleteJoke(c echo.Context) error {
	id, err := strToUint(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	if err := models.DeleteJokeById(id); err != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

func RateJoke(c echo.Context) error {
	var joke models.JokeRequest
	id, err := strToUint(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	if err := decodeBody(c, &joke); err != nil {
		return echo.ErrBadRequest
	}

	jokesResult, err := models.RateJoke(id, joke.Rating)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]any{
		"jokes": jokesResult,
	})
}
