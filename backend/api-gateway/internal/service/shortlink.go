package service

import (
	"fmt"
	"math/rand"
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

	// Ask cache counter to get order number
	// Everytime start server need to check order number from zookeeper
	// orderNumber, err := ctrl.counter.GetOrderNumber()
	// if err != nil {
	// 	return c.JSON(http.StatusUnprocessableEntity, err.Error())
	// }
	orderNumber := rand.Intn(1000000)

	shortenedCode := utils.EncodeBase62(uint64(orderNumber))
	shortLink := fmt.Sprintf("%s/%s", c.Request().Host, shortenedCode)

	_, err = ctrl.linkRepo.Create(model.Link{
		LongLink:  req.LongLink,
		ShortLink: shortLink,
	})
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	// if long link not exist -> hash and create short link then save DB
	// Update order number zookeeper
	// Update cache order number value
	return c.String(http.StatusOK, shortLink)
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
