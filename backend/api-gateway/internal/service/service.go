package service

import (
	"github.com/labstack/echo"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/repository"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/model"
)

// Service is interface to strict service methods.
type Service interface {
	ShortenLink(c echo.Context) error
	GetLongLink(c echo.Context) error
}

type service struct {
	linkRepo repository.Linker
	app      *model.App
}

func New(linkRepo repository.Linker, app *model.App) Service {
	return &service{
		linkRepo: linkRepo,
		app:      app,
	}
}
