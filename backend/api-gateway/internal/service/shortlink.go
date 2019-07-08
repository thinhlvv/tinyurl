package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ShortenLinkRequest struct {
	LongLink string `json:"long_link" form:"long_link"`
}
type ShortenLinkResponse struct {
	ShortLink string `json:"short_link" form:"short_link"`
}

// ShortenLink ...
func (ctrl *service) ShortenLink() func(c echo.Context) error {
	return func(c echo.Context) error {
		// get long link
		req := ShortenLinkRequest{}
		if err := c.Bind(&req); err != nil {
			return err
		}

		// check long link exist -> return short link
		// if long link not exist -> hash link and save DB
		return c.String(http.StatusOK, "ok")
	}
}

// GetLongLink will redirect shortened link to long link
// This API will be heavy traffic. Handle cache here
func (ctrl *service) GetLongLink() func(c echo.Context) error {
	return func(c echo.Context) error {
		shortLink := c.QueryParam("id")
		fmt.Println(shortLink)
		// cache request here
		// count link
		// redirect to long link
		return c.String(http.StatusOK, "ok")
	}
}
