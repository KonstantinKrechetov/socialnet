//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config=server.cfg.yaml ./openapi.json

package api

import (
	"context"
)

// Server implements all the handlers in the ServerInterface
type Server struct{}

// Make sure we conform to StrictServerInterface
var _ StrictServerInterface = (*Server)(nil)

func NewServer() *Server {
	return &Server{}
}

func (p *Server) PostLogin(ctx context.Context, request PostLoginRequestObject) (PostLoginResponseObject, error) {
	resp := "sadas"
	return PostLogin200JSONResponse{
		Token: &resp,
	}, nil
}

func (p *Server) GetUserGetId(ctx context.Context, request GetUserGetIdRequestObject) (GetUserGetIdResponseObject, error) {
	return GetUserGetId404Response{}, nil
}

func (p *Server) PostUserRegister(ctx context.Context, request PostUserRegisterRequestObject) (PostUserRegisterResponseObject, error) {
	return PostUserRegister500JSONResponse{}, nil
}
