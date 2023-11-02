package login

import (
	"context"
	"fmt"
	"socialnet/api"
	"socialnet/db/models"
)

type storage interface {
	SelectItem(ctx context.Context) (models.Item, error)
}

type Handler struct {
	db storage
}

func NewHandler(db storage) *Handler {
	return &Handler{db: db}
}

func (h *Handler) PostLogin(ctx context.Context, request api.PostLoginRequestObject) (api.PostLoginResponseObject, error) {
	item, err := h.db.SelectItem(ctx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return api.PostLogin200JSONResponse{
		Token: &item.Description,
	}, nil
}
