package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"socialnet/models"
)

// сделать методы БД
func (pg *postgres) SelectUser(ctx context.Context) (models.User, error) {
	query := `SELECT * FROM users LIMIT 1`

	rows, err := pg.conn.Query(ctx, query)
	if err != nil {
		return models.User{}, fmt.Errorf("unable to query users: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
}
