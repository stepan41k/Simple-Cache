package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"github.com/stepan41k/MyRest/internal/app"
	"github.com/stepan41k/MyRest/internal/config"
	userHandler "github.com/stepan41k/MyRest/internal/http/handlers/user"
	userService "github.com/stepan41k/MyRest/internal/service/user"
	"github.com/stepan41k/MyRest/internal/storage/postgres"
	"github.com/stepan41k/MyRest/internal/storage/redis"
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

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)


	storagePath := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.PSQL.Host, cfg.PSQL.Port, cfg.PSQL.Username, cfg.PSQL.DBName, os.Getenv("DB_PASSWORD"), cfg.PSQL.SSLMode)

	storage, err := postgres.New(context.Background(), storagePath)
	if err != nil {
		panic(err)
	}
	redis, err := redis.New(context.Background(), cfg.Redis)
	if err != nil {
		panic(err)
	}

	userService := userService.New(storage, redis, log)
	userHandler := userHandler.New(userService, redis, log)


	router.Route("/app/user", func(r chi.Router) {
		r.Get("/get", userHandler.GetUser(context.Background()))
		r.Post("/new", userHandler.CreateUser(context.Background()))
	})

	application := app.New(log, cfg, router)

	go func() {
		application.HTTPServer.Run()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signal := <-stop

	log.Info("stopping application", slog.String("signal", signal.String()))

	application.HTTPServer.Stop(context.Background())

}

func setupLogger(env string) *logrus.Logger{
	var log *logrus.Logger

	switch env {
	case envlocal:
		log.Level = logrus.DebugLevel
		log.Out = os.Stdout
	case envDev:
		log.Level = logrus.DebugLevel
		log.Out = os.Stdout
	case envProd:
		log.Level = logrus.InfoLevel
		log.Out = os.Stdout
	}

	return log

}