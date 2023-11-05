package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"socialnet/models"
)

const createUserQuery = `
	INSERT INTO users (
		first_name,
		second_name,
		birthdate,
		biography,
		city,
		password_hash
	) 
	VALUES (
	        @first_name, 
	        @second_name, 
	        @birthdate, 
	        @biography, 
	        @city, 
	        @password_hash
	)
	RETURNING id`

func (pg *postgres) CreateUser(ctx context.Context, user models.User) (*uuid.UUID, error) {
	args := pgx.NamedArgs{
		"first_name":    user.FirstName,
		"second_name":   user.SecondName,
		"birthdate":     user.Birthdate,
		"biography":     user.Biography,
		"city":          user.City,
		"password_hash": user.PasswordHash,
	}

	row := pg.conn.QueryRow(ctx, createUserQuery, args)

	var userID uuid.UUID
	err := row.Scan(&userID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("no rows returned after inserting user: %w", err)
		}

		return nil, fmt.Errorf("failed inserting user: %w", err)
	}

	return &userID, nil
}
