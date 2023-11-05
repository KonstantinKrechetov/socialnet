package login

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.octolab.org/pointer"
	"log"
	"socialnet/api"
	"socialnet/models"
)

type storage interface {
	SelectUserById(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type encryptor interface {
	ComparePasswordAndHash(password, hash string) bool
}

type Handler struct {
	db        storage
	encryptor encryptor
}

func NewHandler(db storage, encryptor encryptor) *Handler {
	return &Handler{db: db, encryptor: encryptor}
}

func (h *Handler) PostLogin(ctx context.Context, request api.PostLoginRequestObject) (api.PostLoginResponseObject, error) {
	userID, err := uuid.Parse(*request.Body.Id)
	if err != nil {
		log.Println(fmt.Errorf("failed to parse userID to UUID: %w", err))
		return nil, fmt.Errorf("failed to parse userID to UUID: %w", err)
	}

	user, err := h.db.SelectUserById(ctx, userID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	password := request.Body.Password
	if h.encryptor.ComparePasswordAndHash(*password, user.PasswordHash) {
		return api.PostLogin200JSONResponse{
			Token: pointer.ToString("some token"),
		}, nil
	}

	return api.PostLogin404Response{}, nil
}
