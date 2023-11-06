package user_get_id

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
	SelectUserById(ctx context.Context, id uuid.UUID) (*models.User, error)
}

type Handler struct {
	db storage
}

func NewHandler(db storage) *Handler {
	return &Handler{db: db}
}

func (h *Handler) GetUserGetId(w http.ResponseWriter, r *http.Request, id api.UserId) {
	ctx := r.Context()

	// пример получения сессии
	//if _, err := handler.GetCookie(r, login.CookieSessionToken); err != nil {
	//	log.Println(fmt.Errorf("session is empty or expired: %w", err))
	//	handler.SendError(w, http.StatusBadRequest, "session is empty or expired")
	//	return
	//}

	userID, err := uuid.Parse(id)
	if err != nil {
		log.Println(fmt.Errorf("failed convert userId to UUID: %w", err))
		handler.SendError(w, http.StatusBadRequest, "failed convert userId to UUI")
		return
	}

	user, err := h.db.SelectUserById(ctx, userID)
	if err != nil {
		log.Println(err)
		handler.SendError(w, http.StatusInternalServerError, "SelectUserById failed")
		return
	}

	if user == nil {
		handler.SendError(w, http.StatusNotFound, "invalid userID or password")
		return
	}

	birthdate := api.BirthDate{
		Time: user.Birthdate,
	}
	userIdStr := user.Id.String()
	out := api.GetUserGetId200JSONResponse{
		Biography:  &user.Biography,
		Birthdate:  &birthdate,
		City:       &user.City,
		FirstName:  &user.FirstName,
		Id:         &userIdStr,
		SecondName: &user.SecondName,
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(out)
	return
}
