package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/stepan41k/MyRest/internal/config"
	"github.com/stepan41k/MyRest/internal/storage/redis"
	"github.com/stepan41k/MyRest/internal/storage/postgres"
)

const (
	envlocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting app", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
	log.Error("error messages are enabled")

	storagePath := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.PSQL.Host, cfg.PSQL.Port, cfg.PSQL.Username, cfg.PSQL.DBName, os.Getenv("DB_PASSWORD"), cfg.PSQL.SSLMode)

	storage, err := postgres.New(context.Background(), storagePath)
	if err != nil {
		panic(err)
	}
	
	redis, err := redis.New(context.Background(), cfg.Redis)
	if err != nil {
		panic(err)
	}
}

func setupLogger(env string) *slog.Logger{
	var log *slog.Logger

	switch env {
	case envlocal:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log

}