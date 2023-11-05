package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"socialnet/models"
)

func (pg *postgres) SelectUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `
		SELECT * 
		FROM users
		WHERE id = $1
		`

	rows, err := pg.conn.Query(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("SelectUserById failed: %w", err)
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.User])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("SelectUserById pgx.CollectOneRow failed: %w", err)
	}

	return &user, nil
}
