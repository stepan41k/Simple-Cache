package app

import (
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	httpapp "github.com/stepan41k/MyRest/internal/app/http"
	"github.com/stepan41k/MyRest/internal/config"
)

type App struct {
	HTTPServer *httpapp.App
	log *logrus.Logger
}

func New(log *logrus.Logger, cfg *config.Config, router chi.Router) *App {

	httpApp := httpapp.New(log, cfg, router)

	return &App{
		HTTPServer: httpApp,
		log: log,
	}
}