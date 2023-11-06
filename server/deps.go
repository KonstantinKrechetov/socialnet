package server

import (
	"net/http"
	"socialnet/api"
)

type LoginHandler interface {
	PostLogin(w http.ResponseWriter, r *http.Request)
}

type UserGetIdHandler interface {
	GetUserGetId(w http.ResponseWriter, r *http.Request, id api.UserId)
}

type UserRegisterHandler interface {
	PostUserRegister(w http.ResponseWriter, r *http.Request)
}
