package redis

import (
	"context"
	"time"

	"github.com/stepan41k/MyRest/internal/domain/models"
)

func (rs *RedisStorage) Get(name string) error {
	if err := rs.client.Get(context.Background(), name).Err(); err != nil {
		return err
	}

	return nil
}

func (rs *RedisStorage) NewCard(user models.Card) error {
	if err := rs.client.Set(context.Background(), user.Name, user.Data, 1*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}
