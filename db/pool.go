package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

var pgOnce sync.Once

type PoolConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}

func NewPool(ctx context.Context, cfg PoolConfig) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)

	var db *pgxpool.Pool
	var err error
	pgOnce.Do(func() {
		db, err = pgxpool.New(ctx, dsn)
		if err != nil {
			log.Fatal("error connecting to database: %w", err)
		}
	})

	return db
}
