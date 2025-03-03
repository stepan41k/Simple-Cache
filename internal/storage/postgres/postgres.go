package postgres

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	pool *pgxpool.Pool
	mu sync.Mutex
}

func New(ctx context.Context, storagePath string) (*Storage, error) {
	pool, err := pgxpool.Connect(ctx ,storagePath)

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &Storage{
		pool: pool,
		mu: sync.Mutex{},
	}, nil
}