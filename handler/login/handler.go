package login

import (
	"context"
	"fmt"
	"socialnet/api"
	"socialnet/models"
)

type storage interface {
	SelectUser(ctx context.Context) (models.User, error)
}

type Handler struct {
	db storage
}

func NewHandler(db storage) *Handler {
	return &Handler{db: db}
}

// реализовать хендлеры
func (h *Handler) PostLogin(ctx context.Context, request api.PostLoginRequestObject) (api.PostLoginResponseObject, error) {
	user, err := h.db.SelectUser(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return api.PostLogin200JSONResponse{
		Token: &user.FirstName,
	}, nil
}
