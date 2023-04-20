package config
// название файла просто config
import (
	"log"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type Config struct {
	Port int `env:"PORT" envDefault:"3000"`

	PGusername string `env:"PGusername" envDefault:"postgre"`
	PGpassword string `env:"PGpassword"`
	PGhost     string `env:"PGhost" envDefault:"localhost"`
	PGport     int    `env:"PGport" envDefault:"5432"`
	PGdbname   string `env:"PGdbname" envDefault:"postgre"`
	PGsslmode  string `env:"PGsslmode" envDefault:"disable"`
}

func Read() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Printf("%+v\n", err)
	}
	return &cfg, nil
}
