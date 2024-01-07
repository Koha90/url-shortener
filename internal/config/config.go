package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env"          env-default:"local"`
	StoragePath string `yaml:"storage_path"                     env-required:"true"`
	HTTPServer  `       yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address"      env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout"      env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	conigPath := os.Getenv("CONFIG_PATH")
	if conigPath == "" {
		log.Fatal("No path to config. Set CONFIG_PATH.")
	}

	// Check if file exists
	if _, err := os.Stat(conigPath); os.IsNotExist(err) {
		log.Fatalf("config file  does not exist: %s", conigPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(conigPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
