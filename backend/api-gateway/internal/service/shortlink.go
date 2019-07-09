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
	ShortLink string `json:"short_link"`
}

// ShortenLink ...
func (ctrl *service) ShortenLink() func(c echo.Context) error {
	return func(c echo.Context) error {
		fmt.Println("Get here")
		req := ShortenLinkRequest{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err)
		}

		response := ShortenLinkResponse{}
		// check long link exist -> return short link
		link, err := ctrl.linkRepo.GetByLongLink(req.LongLink)
		if err != nil {
			return err
		}
		fmt.Println(req)
		if link != nil {
			response.ShortLink = link.ShortLink
			return c.JSON(http.StatusOK, response)
		}

		fmt.Println(c.Request().URL)

		// shortened := utils.EncodeBase()
		// shortLink :=
		// shortLink:= fmt.Sprintf("")
		// if long link not exist -> hash and create short link then save DB
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
