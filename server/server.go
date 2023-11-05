package server

import (
	"context"
	"socialnet/api"
)

// Server implements all the handlers in the ServerInterface
type Server struct {
	LoginHandler        LoginHandler
	UserGetIdHandler    UserGetIdHandler
	UserRegisterHandler UserRegisterHandler
}

// Make sure we conform to StrictServerInterface
var _ api.StrictServerInterface = (*Server)(nil)

func (s *Server) PostLogin(ctx context.Context, request api.PostLoginRequestObject) (api.PostLoginResponseObject, error) {
	return s.LoginHandler.PostLogin(ctx, request)
}

func (s *Server) GetUserGetId(ctx context.Context, request api.GetUserGetIdRequestObject) (api.GetUserGetIdResponseObject, error) {
	return s.UserGetIdHandler.GetUserGetId(ctx, request)
}

func (s *Server) PostUserRegister(ctx context.Context, request api.PostUserRegisterRequestObject) (api.PostUserRegisterResponseObject, error) {
	return s.UserRegisterHandler.PostUserRegister(ctx, request)
}
