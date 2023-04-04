package database

import (
	"fmt"
	"log"

	"github.com/B-danik/SecondTopic/config"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(cfg *config.Config) (*sqlx.DB, error) {
	log.Println("Connect Database ", cfg.PGdbname)
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.PGhost, cfg.Port, cfg.PGusername, cfg.PGdbname, cfg.PGpassword, cfg.PGsslmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
