package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"socialnet/db/models"
)

func (pg *postgres) SelectItem(ctx context.Context) (models.Item, error) {
	query := `SELECT * FROM items LIMIT 1`

	rows, err := pg.conn.Query(ctx, query)
	if err != nil {
		return models.Item{}, fmt.Errorf("unable to query items: %w", err)
	}
	defer rows.Close()

	return pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Item])
}
