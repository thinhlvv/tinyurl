package main

// TODO: separate db in struct DB to handle DBWithContext, timeout...

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/internalcache"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/zookeeperctl"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/config"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/model"
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

	// Init internal cache engine
	cache := internalcache.New()
	// Init zookeeperctl
	zookeeperctl, err := zookeeperctl.New([]string{"localhost:2181"})
	if err != nil {
		log.Fatal("Can't connect zookeeper:", err)
	}

	// Define service
	app := &model.App{
		InternalCache: cache,
		Zookeeper:     zookeeperctl,
	}
	service := service.New(linkRepo, app)

	// Define handler.
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/shorten-link", service.ShortenLink)
	e.GET("/:id", service.GetLongLink)

	e.Logger.Fatal(e.Start(cfg.HTTP.ConnectionURL()))
}
