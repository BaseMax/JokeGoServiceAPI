package controllers

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func decodeBody(c echo.Context, i any) error {
	if c.Request().Body == nil {
		return echo.ErrBadRequest
	}
	if err := json.NewDecoder(c.Request().Body).Decode(i); err != nil {
		return echo.ErrBadRequest
	}
	return nil
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

func strToUint(param string) (uint, error) {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
