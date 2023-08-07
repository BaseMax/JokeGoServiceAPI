package controllers

import (
	"net/http"

	"github.com/BaseMax/JokeGoServiceAPI/models"
	"github.com/labstack/echo/v4"
)

func CreateJokeComment(c echo.Context) error {
	var comment models.CommentRequest
	if err := decodeBody(c, &comment); err != nil {
		return echo.ErrBadRequest
	}
	id, err := strToUint(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	err = models.CreateComment(id, &comment)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

func GetJokeComment(c echo.Context) error {
	id, err := strToUint(c.Param("comment_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	comment, err := models.FetchCommentById(id)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, comment)
}

func GetJokeComments(c echo.Context) error {
	id, err := strToUint(c.Param("joke_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	comments, err := models.FetchAllComments(id)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, comments)
}

func EditJokeComment(c echo.Context) error {
	var comment models.CommentRequest
	if err := decodeBody(c, &comment); err != nil {
		return echo.ErrBadRequest
	}
	id, err := strToUint(c.Param("comment_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	err = models.UpdateComment(id, &comment)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, comment)
}

func DeleteJokeComment(c echo.Context) error {
	id, err := strToUint(c.Param("comment_id"))
	if err != nil {
		return echo.ErrBadRequest
	}

	err = models.DeleteComment(id)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
