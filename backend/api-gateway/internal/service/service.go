package service

import (
	"github.com/labstack/echo"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/repository"
)

// Service is interface to strict service methods.
// Why return function instead of error. Easier to change
// service interface.
type Service interface {
	ShortenLink() func(c echo.Context) error
	GetLongLink() func(c echo.Context) error
}

type service struct {
	linkRepo repository.Linker
}

func New(linkRepo repository.Linker) Service {
	return &service{linkRepo: linkRepo}
}
