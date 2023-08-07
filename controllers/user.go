package controllers

import (
	"net/http"

	"github.com/BaseMax/JokeGoServiceAPI/models"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var u models.User
	if err := decodeBody(c, &u); err != nil {
		return echo.ErrBadRequest
	}

	err := models.RegisterUser(&u)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	if err != nil {
		return echo.ErrInternalServerError
	}

	bearer, err := createToken(u.ID, u.Username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"bearer": bearer,
	})
}

func Login(c echo.Context) error {
	var u models.User
	if err := decodeBody(c, &u); err != nil {
		return echo.ErrBadRequest
	}
	if err := models.LoginUser(&u); err != nil {
		return echo.ErrNotFound
	}

	bearer, err := createToken(u.ID, u.Username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"bearer": bearer,
	})
}

func Refresh(c echo.Context) error {
	claims, err := getClaims(c)
	if err != nil {
		return echo.ErrBadRequest
	}

	bearer, err := createToken(claims.ID, claims.Issuer)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"bearer": bearer,
	})
}
