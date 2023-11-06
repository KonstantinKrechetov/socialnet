package login

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go.octolab.org/pointer"
	"log"
	"net/http"
	"socialnet/api"
	"socialnet/handler"
	"socialnet/models"
	"time"
)

const CookieSessionToken = "session_token"
const CookieSessionTokenTTL = 120 * time.Second

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

func (h *Handler) PostLogin(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req api.PostLoginJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		handler.SendError(w, http.StatusBadRequest, "invalid request")
		return
	}

	userID, err := uuid.Parse(*req.Id)
	if err != nil {
		log.Println(fmt.Errorf("failed to parse userID to UUID: %w", err))
		handler.SendError(w, http.StatusBadRequest, "failed to parse userID to UUID")
		return
	}

	user, err := h.db.SelectUserById(ctx, userID)
	if err != nil {
		log.Println(err)
		handler.SendError(w, http.StatusBadRequest, "SelectUserById failed")
		return
	}

	password := req.Password
	if !h.encryptor.ComparePasswordAndHash(*password, user.PasswordHash) {
		handler.SendError(w, http.StatusNotFound, "invalid userID or password")
		return
	}

	// пример проставления токена
	sessionToken := uuid.NewString()
	out := api.PostLogin200JSONResponse{
		Token: pointer.ToString(sessionToken),
	}

	http.SetCookie(w, &http.Cookie{
		Name:    CookieSessionToken,
		Value:   sessionToken,
		Expires: time.Now().Add(CookieSessionTokenTTL),
	})

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(out)
	return
}
