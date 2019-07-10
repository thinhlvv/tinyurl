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
func (ctrl *service) ShortenLink(c echo.Context) error {
	req := ShortenLinkRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	response := ShortenLinkResponse{}
	// check long link exist -> return short link
	link, err := ctrl.linkRepo.GetByLongLink(req.LongLink)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	if link != nil {
		response.ShortLink = link.ShortLink
		return c.JSON(http.StatusOK, response)
	}

	// Ask counter to get next id
	// shortenedCode := utils.EncodeBase()
	// shortLink:= fmt.Sprintf("%s/%s",c.Request().Host, shortenedCode)
	// if long link not exist -> hash and create short link then save DB
	return c.String(http.StatusOK, "ok")
}

// GetLongLink will redirect shortened link to long link
// This API will be heavy traffic. Handle cache here
func (ctrl *service) GetLongLink(c echo.Context) error {
	shortLink := c.QueryParam("id")
	fmt.Println(shortLink)
	// cache request here
	// count link
	// redirect to long link
	return c.String(http.StatusOK, "ok")
}
