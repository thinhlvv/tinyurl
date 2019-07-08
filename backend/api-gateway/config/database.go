package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func MustInitDB(dataSourceName string) *sql.DB {
	var err error
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}
