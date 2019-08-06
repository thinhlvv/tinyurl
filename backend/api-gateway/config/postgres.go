package config

import "fmt"

type Postgres struct {
	Host            string `yaml:"host"`
	Port            int    `yaml:"port"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	Database        string `yaml:"database"`
	SSLMode         string `yaml:"ssl_mode"`
	Timeout         int
	Protocol        string
	GoogleAuthFile  string
	MaxOpenConn     int
	MaxIdleConn     int
	ConnMaxLifeTime int
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

func (p *Postgres) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", p.Username, p.Password, p.Host, p.Port, p.Database)
}
