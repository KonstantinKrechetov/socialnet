package user_register

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
	CreateUser(ctx context.Context, user models.User) (*uuid.UUID, error)
}

type encryptor interface {
	HashPassword(password string) (string, error)
}

type Handler struct {
	db        storage
	encryptor encryptor
}

func NewHandler(db storage, encryptor encryptor) *Handler {
	return &Handler{db: db, encryptor: encryptor}
}

func (h *Handler) PostUserRegister(ctx context.Context, request api.PostUserRegisterRequestObject) (api.PostUserRegisterResponseObject, error) {
	passwordHashed, err := h.encryptor.HashPassword(*request.Body.Password)
	if err != nil {
		log.Println(fmt.Errorf("failed to hash password: %w", err))
		return api.PostUserRegister500JSONResponse{
			N5xxJSONResponse: api.N5xxJSONResponse{
				Body: struct {
					Code      *int    `json:"code,omitempty"`
					Message   string  `json:"message"`
					RequestId *string `json:"request_id,omitempty"`
				}{
					Code:    pointer.ToInt(500),
					Message: "failed to hash password",
				},
			},
		}, nil
	}

	userID, err := h.db.CreateUser(ctx, models.User{
		FirstName:    *request.Body.FirstName,
		SecondName:   *request.Body.SecondName,
		Birthdate:    request.Body.Birthdate.Time,
		Biography:    *request.Body.Biography,
		City:         *request.Body.City,
		PasswordHash: passwordHashed,
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userIdStr := userID.String()
	return api.PostUserRegister200JSONResponse{
		UserId: &userIdStr,
	}, nil
}
