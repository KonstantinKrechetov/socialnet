package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"socialnet/api"
)

func SendError(w http.ResponseWriter, code int, message string) {
	petErr := api.N5xx{
		Code:    &code,
		Message: message,
	}
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(petErr)
}

func GetCookie(r *http.Request, cookieName string) (*string, error) {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		switch {
		case errors.Is(err, http.ErrNoCookie):
			return nil, fmt.Errorf(fmt.Sprintf("cookie %s is not set", cookieName))
		default:
			log.Println(err)
			return nil, fmt.Errorf("internal error %w", err)
		}
	}

	// проверяем актуальна ли сессия
	//if cookie.Expires.Before(cookieExpirationFromCache) {
	//	return nil, fmt.Errorf("cookie is expired")
	//}

	return &cookie.Value, nil
}
