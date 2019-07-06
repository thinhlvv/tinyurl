package main

import (
	"github.com/labstack/echo"
	"github.com/thinhlvv/tinyurl/backend/config"
	"github.com/thinhlvv/tinyurl/backend/internal/service"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	cfgPath := kingpin.Flag("config", "Path to config file").Short('c').Default().String()
	kingpin.Parse()
	cfg := config.MustLoad(*cfgPath)

	e := echo.New()
	service := service.New()

	e.POST("/shorten-link", service.ShortenLink())

	e.Logger.Fatal(e.Start(cfg.HTTP.ConnectionURL()))
}
