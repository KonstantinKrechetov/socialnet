//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ./openapi.json

package api

import (
	"context"
)

type LoginHandler interface {
	PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error)
}

// Server implements all the handlers in the ServerInterface
type Server struct {
	loginHandler LoginHandler
}

// Make sure we conform to StrictServerInterface
var _ StrictServerInterface = (*Server)(nil)

func NewServer(loginHandler LoginHandler) *Server {
	return &Server{
		loginHandler: loginHandler,
	}
}

func (s *Server) PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error) {
	return s.loginHandler.PostLogin(ctx, request)
}

func (s *Server) GetUserGetId(ctx context.Context, request GetUserGetIdRequestObject) (GetUserGetIdResponseObject, error) {
	return GetUserGetId404Response{}, nil
}

func (s *Server) PostUserRegister(ctx context.Context, request PostUserRegisterRequestObject) (PostUserRegisterResponseObject, error) {
	return PostUserRegister500JSONResponse{}, nil
}
