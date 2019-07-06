package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func (ctrl *service) ShortenLink() func(c echo.Context) error {
	return func(c echo.Context) error {
		// get long link
		// check long link exist -> return short link
		// if long link not exist -> hash link and save DB
		return c.String(http.StatusOK, "ok")
	}
}

func (ctrl *service) GetLongLink() func(c echo.Context) error {
	return func(c echo.Context) error {
		shortLink := c.QueryParam("id")
		fmt.Println(shortLink)
		// count link
		// redirect to long link
		return c.String(http.StatusOK, "ok")
	}
}
