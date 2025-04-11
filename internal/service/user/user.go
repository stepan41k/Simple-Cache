package cache

import (
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

type UserService struct {
	userPG UserPG
	userRedis UserRedis 
	log *logrus.Logger
}

func New(userPG UserPG, userRedis UserRedis, log *logrus.Logger) *UserService {
	return &UserService{
		userPG: userPG,
		userRedis: userRedis,
		log: log,
	}
}

func (us *UserService) CreateUser(user models.User) (models.Card, error) {
	
}

func (us *UserService) GetUser(name string) (models.Card, error) {
	
}

