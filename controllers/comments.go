package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	"github.com/BaseMax/JokeGoServiceAPI/models"
)

func CreateJokeComment(c echo.Context) error {
	var comment models.CommentRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&comment); err != nil {
		return echo.ErrBadRequest
	}
	joke_id, err := strconv.Atoi(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}
	bearer := c.Request().Header.Get("Authorization")
	token, _, _ := new(jwt.Parser).ParseUnverified(bearer[len("Bearer "):], jwt.MapClaims{})
	claims := token.Claims.(jwt.MapClaims)
	author, _ := claims.GetIssuer()

	comment.Author = author
	err = models.CreateComment(uint(joke_id), &comment)
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, comment)
}

func GetJokeComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	comment, err := models.FetchCommentById(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}

	return c.JSON(http.StatusOK, comment)
}

func GetJokeComments(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	comments, err := models.FetchAllComments(uint(id))
	if err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, comments)
}

func EditJokeComment(c echo.Context) error {
	var comment models.CommentRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&comment); err != nil {
		return echo.ErrBadRequest
	}
	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	if models.UpdateComment(uint(id), &comment) != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, comment)
}

func DeleteJokeComment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("comment_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	if models.DeleteComment(uint(id)) != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}
