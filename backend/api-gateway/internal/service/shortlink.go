package service

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/utils"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
)

// ShortenLinkRequest ...
type ShortenLinkRequest struct {
	LongLink string `json:"long_link" form:"long_link"`
}

// ShortenLinkResponse ...
type ShortenLinkResponse struct {
	ShortLink string `json:"short_link,omitempty"`
	Error     error  `json:"error,omitempty"`
}

// ShortenLink ...
func (ctrl *service) ShortenLink(c echo.Context) error {
	// if long link not exist -> hash and create short link then save DB
	// Update cache order number value

	req := ShortenLinkRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ShortenLinkResponse{Error: err})
	}

	response := ShortenLinkResponse{}

	link, err := ctrl.linkRepo.GetByLongLink(req.LongLink)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ShortenLinkResponse{Error: err})
	}

	if link != nil {
		response.ShortLink = link.ShortLink
		return c.JSON(http.StatusOK, response)
	}

	// Ask cache counter to get order number
	// Everytime start server need to check order number from zookeeper

	orderNumber, err := ctrl.counter.GetOrderNumber()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ShortenLinkResponse{Error: err})
	}

	shortenedCode := utils.EncodeBase62(uint64(orderNumber))
	shortLink := fmt.Sprintf("%s/%s", c.Request().Host, shortenedCode)

	_, err = ctrl.linkRepo.Create(model.Link{
		LongLink:  req.LongLink,
		ShortLink: shortLink,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, ShortenLinkResponse{Error: err})
	}

	return c.JSON(http.StatusOK, ShortenLinkResponse{ShortLink: shortLink})
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
