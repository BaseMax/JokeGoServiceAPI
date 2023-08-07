package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var jwtKey []byte

func decodeBody(c echo.Context, i any) error {
	if err := json.NewDecoder(c.Request().Body).Decode(i); err != nil {
		return echo.ErrBadRequest
	}
	return nil
}

func createToken(id any, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ID:        fmt.Sprint(id),
		Issuer:    name,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 30)),
	})

	if jwtKey == nil {
		godotenv.Load(".env")
		jwtKey = []byte(os.Getenv("JWT_KET"))
	}

	bearer, err := token.SignedString(jwtKey)
	if err != nil {
		return "", nil
	}

	return bearer, nil
}

func dbErrorToHttp(err error) *echo.HTTPError {
	var me *mysql.MySQLError
	switch {
	case err == nil:
		return nil
	case errors.As(err, &me) && me.Number == 1062:
		return echo.ErrConflict
	case err == gorm.ErrRecordNotFound:
		return echo.ErrNotFound
	}
	return echo.ErrInternalServerError
}

func getClaims(c echo.Context) (*jwt.RegisteredClaims, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, echo.ErrBadRequest
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, echo.ErrBadRequest
	}
	return claims, nil
}

func strToUint(param string) (uint, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
