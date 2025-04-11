package user

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/stepan41k/MyRest/internal/domain/models"
)

type UserPG interface {
	CreateUser(user models.User) (models.Card, error)
	GetUser(name string) (models.Card, error)
}

type UserRedis interface {
	Get(name string) error
	NewCard(user models.Card) error
}

type UserHandler struct {
	userPG UserPG
	userRedis UserRedis
	log *logrus.Logger
}

func New(userPG UserPG, userRedis UserRedis, log *logrus.Logger) UserHandler {
	return UserHandler{
		userPG: userPG,
		userRedis: userRedis,
		log: log,
	}
}

func (uh *UserHandler) CreateUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (uh *UserHandler) GetUser(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
	}
}