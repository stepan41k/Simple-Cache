package postgres

import (
	"context"
	"fmt"

	"github.com/stepan41k/MyRest/internal/domain/models"
)

func (s *Storage) CreateUser(user models.User) (models.Card, error)  {
	const op = "storage.postgres.NewUser"

	row := s.pool.QueryRow(context.Background(), `
		INSERT INTO users (name, data)
		VALUES ($1, $2)
		RETURNING name, data
	`, user.Name, user.Data)

	var card models.Card

	err := row.Scan(&card.Name, &card.Data)
	if err != nil {
		return models.Card{}, fmt.Errorf("%s: %w", op ,err)
	}

	return card, nil
}

func (s *Storage) GetUser(name string) (models.Card, error) {
	const op = "storage.postgres.GetUser"

	row := s.pool.QueryRow(context.Background(), `
		SELECT name, data
		FROM users
		WHERE name = $1
	`, name)

	var card models.Card

	err := row.Scan(&card.Name, &card.Data)
	if err != nil {
		return models.Card{}, fmt.Errorf("%s: %w", op, err)
	}

	return card, nil
}