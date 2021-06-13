package config

import (
	"coursera/creds"
	"database/sql"
	"log"

	pg "gopkg.in/pg.v4"
)

type Config struct {
	PG   *pg.DB
	DB   *sql.DB
	Port string
}

func InitPG() *pg.DB {
	log.Println("[PG] Connecting to " + creds.PG_HOST + "/" + creds.PG_DB_NAME)
	db := pg.Connect(&pg.Options{
		Addr:     creds.PG_HOST,
		User:     creds.PG_USER,
		Password: creds.PG_PASSWORD,
		Database: creds.PG_DB_NAME,
	})
	return db
}
