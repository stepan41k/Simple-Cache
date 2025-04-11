package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"github.com/stepan41k/MyRest/internal/config"
)

type App struct {
	log *logrus.Logger
	httpServer *http.Server
}

func New(log *logrus.Logger, cfg *config.Config, router chi.Router) *App {
	httpServer := http.Server{
		Addr: cfg.Port,
		Handler: router,
		ReadTimeout: cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout: cfg.IdleTimeout,
	}

	return &App{log:log, httpServer: &httpServer}
}

func (a *App) Run() error {
	const op = "httpapp.Run"

	log := a.log.WithFields(logrus.Fields{
		"op": op,
		"port": a.httpServer.Addr,
	})

	log.Info("starting http server")

	if err := a.httpServer.ListenAndServe(); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("http server started")

	return nil
}

func (a *App) Stop(ctx context.Context) {
	const op = "httpapp.Stop"

	log := a.log.WithFields(logrus.Fields{
		"op": op,
	})

	log.Info("stoping http server")

	a.httpServer.Shutdown(ctx)
}
