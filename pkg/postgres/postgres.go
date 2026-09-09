package postgres

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func New(ctx context.Context, dsn string) *Postgres {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("ошибка создания пула подключений: %s", err)
		os.Exit(1)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("ошибка подключения к БД: %s", err)
		os.Exit(1)
	}

	return &Postgres{
		Pool: pool,
	}
}

func (p *Postgres) Close() {
	p.Pool.Close()
}
