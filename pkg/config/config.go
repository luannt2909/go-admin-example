package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cast"
	"os"
)

const DevelopmentMode = false
const MaximumReminder = 5

type DBConfig struct {
	DBClient        string `env:"DB_CLIENT"`
	DBConnectionURI string `env:"DB_CONNECTION_URI"`
}

type Config struct {
	DBConfig
	Port            string `env:"PORT" envDefault:"2909"`
	JwtSigningKey   string `env:"JWT_SIGNING_KEY" envDefault:"CSwS88WnQjKGBAEI"`
	MaximumReminder uint   `env:"MAX_REMINDER" envDefault:"5"`
}

func LoadEnv() (cfg Config, err error) {
	if DevelopmentMode {
		return LoadEnvFromFile()
	}
	err = env.Parse(&cfg)
	if err != nil {
		log.Printf("failed to config load from ENV: %s", err)
		return
	}
	log.Printf("config load from ENV: %+v", cfg)
	return
}

func LoadEnvFromFile() (cfg Config, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Printf("failed to config load from ENV: %s", err)
		return
	}
	cfg = Config{
		DBConfig: DBConfig{
			DBClient:        os.Getenv("DB_CLIENT"),
			DBConnectionURI: os.Getenv("DB_CONNECTION_URI"),
		},
		Port:            os.Getenv("PORT"),
		JwtSigningKey:   os.Getenv("JWT_SIGNING_KEY"),
		MaximumReminder: cast.ToUint(os.Getenv("MAX_REMINDER")),
	}
	if cfg.MaximumReminder == 0 {
		cfg.MaximumReminder = MaximumReminder
	}
	log.Printf("config load from ENV: %+v", cfg)
	return
}
