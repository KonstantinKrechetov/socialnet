package login

import (
	"context"
	"socialnet/api"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) PostLogin(ctx context.Context, request api.PostLoginRequestObject) (api.PostLoginResponseObject, error) {
	resp := "success"
	return api.PostLogin200JSONResponse{
		Token: &resp,
	}, nil
}
