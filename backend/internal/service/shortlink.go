package service

import (
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *service) ShortenLink() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "ok")
	}
}
