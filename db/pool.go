package db

import (
	"context"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
)

const pathToMigrations = "db/migrations"

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

func Migrate(cfg PoolConfig) error {
	migrateDsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DbName)
	migrateSrc := fmt.Sprintf("file://%s", pathToMigrations)

	m, err := migrate.New(migrateSrc, migrateDsn)
	if err != nil {
		return fmt.Errorf("error creating migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrate up failed %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		return fmt.Errorf("failed to get migrate version: %w", err)
	}

	log.Println(fmt.Sprintf("Applied migration: %d, dirty: %t", version, dirty))
	return nil
}
