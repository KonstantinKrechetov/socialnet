package user_register

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
	"socialnet/api"
	"socialnet/handler"
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

func (h *Handler) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req api.PostUserRegisterJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.SendError(w, http.StatusBadRequest, "invalid request")
		return
	}

	passwordHashed, err := h.encryptor.HashPassword(*req.Password)
	if err != nil {
		log.Println(fmt.Errorf("failed to hash password: %w", err))
		handler.SendError(w, http.StatusInternalServerError, "HashPassword failed")
		return
	}

	userID, err := h.db.CreateUser(ctx, models.User{
		FirstName:    *req.FirstName,
		SecondName:   *req.SecondName,
		Birthdate:    req.Birthdate.Time,
		Biography:    *req.Biography,
		City:         *req.City,
		PasswordHash: passwordHashed,
	})
	if err != nil {
		log.Println(err)
		handler.SendError(w, http.StatusInternalServerError, "CreateUser failed")
		return
	}

	userIdStr := userID.String()
	out := api.PostUserRegister200JSONResponse{
		UserId: &userIdStr,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(out)
	return
}
