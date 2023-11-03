//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ./openapi.json

package api

import (
	"context"
)

type LoginHandler interface {
	PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error)
}

type UserGetIdHandler interface {
	GetUserGetId(ctx context.Context, request GetUserGetIdRequestObject) (GetUserGetIdResponseObject, error)
}

type UserRegisterHandler interface {
	PostUserRegister(ctx context.Context, request PostUserRegisterRequestObject) (PostUserRegisterResponseObject, error)
}

// Server implements all the handlers in the ServerInterface
type Server struct {
	loginHandler        LoginHandler
	userGetIdHandler    UserGetIdHandler
	userRegisterHandler UserRegisterHandler
}

// Make sure we conform to StrictServerInterface
var _ StrictServerInterface = (*Server)(nil)

func NewServer(
	loginHandler LoginHandler,
	userGetIdHandler UserGetIdHandler,
	userRegisterHandler UserRegisterHandler,
) *Server {
	return &Server{
		loginHandler:        loginHandler,
		userGetIdHandler:    userGetIdHandler,
		userRegisterHandler: userRegisterHandler,
	}
}

func (s *Server) PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error) {
	return s.loginHandler.PostLogin(ctx, request)
}

func (s *Server) GetUserGetId(ctx context.Context, request GetUserGetIdRequestObject) (GetUserGetIdResponseObject, error) {
	return s.userGetIdHandler.GetUserGetId(ctx, request)
}

func (s *Server) PostUserRegister(ctx context.Context, request PostUserRegisterRequestObject) (PostUserRegisterResponseObject, error) {
	return s.userRegisterHandler.PostUserRegister(ctx, request)
}
