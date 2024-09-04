/*
Copyright © 2024 Teruaki Sato <andrea.pirlo.0529@gmail.com>
*/
package infra

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

// TITLE:環境変数の読込み

// STRUCT:
type Config struct {
	DebugMode bool `env:"DEBUG" default:"false"`
	Postgres  POSTGRES
}

type POSTGRES struct {
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Hostname string `env:"POSTGRES_HOST_NAME"`
	Port     string `env:"POSTGRES_PORT"`
	Db       string `env:"POSTGRES_DB"`
}

func LeadEnv() *Config {
	// envファイルのロード
	_, err := os.Stat(".env")
	if !os.IsNotExist(err) {
		godotenv.Load()
		log.Print("loaded environment variables from .env file.")
	}

	var config Config
	envconfig.Process(context.Background(), &config)
	return &config
}
