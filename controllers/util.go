package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

func isDuplicatedKeyError(err error) bool {
	var me *mysql.MySQLError
	return errors.As(err, &me) && me.Number == 1062
}
