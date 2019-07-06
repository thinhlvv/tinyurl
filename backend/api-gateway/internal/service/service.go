package service

import "github.com/labstack/echo"

// Service is interface to strict service methods.
// Why return function instead of error. Easier to change
// service interface.
type Service interface {
	ShortenLink() func(c echo.Context) error
	GetLongLink() func(c echo.Context) error
}

type service struct {
	// services
	// repo
}

func New() Service {
	return &service{}
}
