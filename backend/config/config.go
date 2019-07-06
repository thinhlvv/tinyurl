package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Postgres *Postgres `yaml:"postgres"`
	HTTP     *HTTP     `yaml:"http"`
}

// HTTP represents config for serving files over HTTP
type HTTP struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (h *HTTP) ConnectionURL() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
}

type Postgres struct {
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	SSLMode         string
	Timeout         int
	Protocol        string
	GoogleAuthFile  string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
}

func DefaultHTTP() *HTTP {
	return &HTTP{
		Host: "",
		Port: "1323",
	}
}

// DefaultPostgres returns default Postgres object
func DefaultPostgres() *Postgres {
	return &Postgres{
		Host:            "postgres",
		Port:            5432,
		Username:        "postgres",
		Password:        "postgres",
		Database:        "test",
		SSLMode:         "",
		Timeout:         15,
		Protocol:        "",
		GoogleAuthFile:  "",
		MaxOpenConn:     10,
		MaxIdleConn:     2,
		ConnMaxLifeTime: 1800,
	}
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{
		Postgres: DefaultPostgres(),
		HTTP:     DefaultHTTP(),
	}
}

// Load ...
func Load(path string) (*Config, error) {
	if path == "" {
		return DefaultConfig(), nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(data, &cfg)
	return cfg, err
}

// MustLoad ...
func MustLoad(path string) *Config {
	cfg, err := Load(path)
	if err != nil {
		panic(err)
	}
	return cfg
}
