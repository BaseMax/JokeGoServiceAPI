package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/BaseMax/JokeGoServiceAPI/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}

	err := models.RegisterUser(&user)
	if err := dbErrorToHttp(err); err != nil {
		return err
	}
	if err != nil {
		return echo.ErrInternalServerError
	}

	bearer, err := createToken(user.ID, user.Username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"bearer": bearer,
	})
}

func Login(c echo.Context) error {
	var user models.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return echo.ErrBadRequest
	}
	if err := models.LoginUser(&user); err != nil {
		return echo.ErrNotFound
	}

	bearer, err := createToken(user.ID, user.Username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"bearer": bearer,
	})
}

func Refresh(c echo.Context) error {
	bearer := c.Request().Header.Get("Authorization")
	bearer = bearer[len("Bearer "):]

	token, _, err := new(jwt.Parser).ParseUnverified(bearer, jwt.MapClaims{})
	if err != nil {
		return echo.ErrBadRequest
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return echo.ErrBadRequest
	}
	bearer, err = createToken(claims["jti"], claims["iss"].(string))
	if err != nil {
		return echo.ErrBadRequest
	}

	return c.JSON(http.StatusOK, map[string]string{
		"bearer": bearer,
	})
}
