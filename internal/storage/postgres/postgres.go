package postgres

import (
	"sync"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	mu sync.Mutex
	pool *pgxpool.Pool
}

func New(storagePath string) (*Storage)