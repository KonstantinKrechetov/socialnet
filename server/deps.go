package server

import (
	"context"
	"socialnet/api"
)

type LoginHandler interface {
	PostLogin(ctx context.Context, request api.PostLoginRequestObject) (api.PostLoginResponseObject, error)
}

type UserGetIdHandler interface {
	GetUserGetId(ctx context.Context, request api.GetUserGetIdRequestObject) (api.GetUserGetIdResponseObject, error)
}

type UserRegisterHandler interface {
	PostUserRegister(ctx context.Context, request api.PostUserRegisterRequestObject) (api.PostUserRegisterResponseObject, error)
}
