package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_sever"`
}

type HTTPServer struct {
	Address     string `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"http_sever" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"http_sever" env-default:"60s"`
	User        string `yaml:"http_sever" env-required:"true"`
	Password    string `yaml:"http_sever" env-required:"true" env:"HTTP_SERVER_PASSWORD"`
}

func MustLoad() *Config {
	os.Setenv("CONFIG_PATH", "C:/Users/Stepan/go/go1.22.2/src/sandbox/project4/MyRestTest/config/local.yaml")
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}