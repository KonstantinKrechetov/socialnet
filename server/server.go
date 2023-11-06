package server

import (
	"net/http"
	"socialnet/api"
)

// Server implements all the handlers in the ServerInterface
type Server struct {
	LoginHandler        LoginHandler
	UserGetIdHandler    UserGetIdHandler
	UserRegisterHandler UserRegisterHandler
}

// Make sure we conform to StrictServerInterface
var _ api.ServerInterface = (*Server)(nil)

func (s *Server) PostLogin(w http.ResponseWriter, r *http.Request) {
	s.LoginHandler.PostLogin(w, r)
}

func (s *Server) GetUserGetId(w http.ResponseWriter, r *http.Request, id api.UserId) {
	s.UserGetIdHandler.GetUserGetId(w, r, id)
}

func (s *Server) PostUserRegister(w http.ResponseWriter, r *http.Request) {
	s.UserRegisterHandler.PostUserRegister(w, r)
}
