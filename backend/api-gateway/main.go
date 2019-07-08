package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/config"
	"github.com/thinhlvv/tinyurl/backend/api-gateway/internal/service"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// Load config.
	cfgPath := kingpin.Flag("config", "Path to config file").Short('c').Default().String()
	kingpin.Parse()
	cfg := config.MustLoad(*cfgPath)

	// Init db
	fmt.Println(cfg.Postgres.ConnectionString())
	db := config.MustInitDB(cfg.Postgres.ConnectionString())
	fmt.Println(db)

	// Define handler.
	e := echo.New()
	service := service.New()

	e.POST("/shorten-link", service.ShortenLink())
	e.GET("/:id", service.GetLongLink())

	e.Logger.Fatal(e.Start(cfg.HTTP.ConnectionURL()))
}
