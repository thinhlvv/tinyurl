package service

import (
	"github.com/labstack/echo"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/counter"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/repository"
)

// Service is interface to strict service methods.
type Service interface {
	ShortenLink(c echo.Context) error
	GetLongLink(c echo.Context) error
}

type service struct {
	linkRepo repository.Linker
	app      *model.App
	counter  *counter.Counter
}

func New(linkRepo repository.Linker, counter *counter.Counter, app *model.App) Service {
	return &service{
		linkRepo: linkRepo,
		counter:  counter,
		app:      app,
	}
}
