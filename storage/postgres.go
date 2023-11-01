package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type postgres struct {
	conn *pgxpool.Pool
}

func NewPostgres(conn *pgxpool.Pool) *postgres {
	return &postgres{
		conn: conn,
	}
}

func (pg *postgres) Ping(ctx context.Context) error {
	return pg.conn.Ping(ctx)
}

func (pg *postgres) Close() {
	pg.conn.Close()
}
