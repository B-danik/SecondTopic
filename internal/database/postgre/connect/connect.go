package database

import (
	"fmt"
	"log"

	"github.com/B-danik/SecondTopic/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// убрать из папки connect
func NewPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	log.Println("Connect to database")
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		cfg.PGhost, cfg.PGport, cfg.PGusername, cfg.PGdbname, cfg.PGpassword, cfg.PGsslmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
