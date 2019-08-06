package config

import "fmt"

// HTTP represents config for serving files over HTTP
type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func DefaultHTTP() *HTTP {
	return &HTTP{
		Host: "",
		Port: "1323",
	}
}

func (h *HTTP) ConnectionURL() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}
