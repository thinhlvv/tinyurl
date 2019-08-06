package config

import (
	"io/ioutil"

	"github.com/thinhlvv/tinyurl/backend/api-gateway/common/utils"
	yaml "gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Postgres   *Postgres `yaml:"postgres"`
	HTTP       *HTTP     `yaml:"http"`
	ServerName string    `yaml:"server_name"` // used for identity in zookeeper
}

// DefaultConfig ...
func DefaultConfig() *Config {
	return &Config{
		Postgres:   DefaultPostgres(),
		HTTP:       DefaultHTTP(),
		ServerName: utils.RandomString(6),
	}
}

// MustLoad config file path.
func MustLoad(path string) *Config {
	cfg, err := Load(path)
	if err != nil {
		panic(err)
	}
	return cfg
}

// Load yaml file path.
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
