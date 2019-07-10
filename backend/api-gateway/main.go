package main

// TODO: separate db in struct DB to handle DBWithContext, timeout...

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	// Migrate
	if err := config.MustMigrate(db); err != nil {
		log.Fatal(err)
	}

	// Init Repository
	linkRepo := repository.NewLinkRepo(db)

	// Define handler.
	e := echo.New()
	e.Use(middleware.Logger())
	service := service.New(linkRepo)

	e.POST("/shorten-link", service.ShortenLink)
	e.GET("/:id", service.GetLongLink)

	e.Logger.Fatal(e.Start(cfg.HTTP.ConnectionURL()))
}
