package main

import (
	"github.com/labstack/echo"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/config"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/repository"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/service"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// Load config.
	cfgPath := kingpin.Flag("config", "Path to config file").Short('c').Default().String()
	kingpin.Parse()
	cfg := config.MustLoad(*cfgPath)

	// Init db
	db := config.MustInitDB(cfg.Postgres.ConnectionString())

	// Init Repository
	linkRepo := repository.NewLinkRepo(db)

	// Define handler.
	e := echo.New()
	service := service.New(linkRepo)

	e.POST("/shorten-link", service.ShortenLink())
	e.GET("/:id", service.GetLongLink())

	e.Logger.Fatal(e.Start(cfg.HTTP.ConnectionURL()))
}
