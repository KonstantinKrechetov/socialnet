package user_get_id

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"socialnet/api"
	"socialnet/models"
)

type storage interface {
	SelectUserById(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type Handler struct {
	db storage
}

func NewHandler(db storage) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetUserGetId(ctx context.Context, request api.GetUserGetIdRequestObject) (api.GetUserGetIdResponseObject, error) {
	userID, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, fmt.Errorf("failed convert userId to UUID: %w", err)
	}

	user, err := h.db.SelectUserById(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	birthdate := api.BirthDate{
		Time: user.Birthdate,
	}
	userIdStr := user.Id.String()
	return api.GetUserGetId200JSONResponse{
		Biography:  &user.Biography,
		Birthdate:  &birthdate,
		City:       &user.City,
		FirstName:  &user.FirstName,
		Id:         &userIdStr,
		SecondName: &user.SecondName,
	}, nil
}
