//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/api"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/api/handler"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/config"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/db"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/repository"
	"github.com/sgokul961/echo-hub-auth-svc/pkg/usecase"
)

func InitApi(cfg config.Config) (*api.ServerHTTP, error) {
	wire.Build(db.Init,
		repository.NewUserRepo,
		usecase.NewUserUseCase,
		handler.NewUserHandler,

		repository.NewAdminRepo,
		usecase.NewAdminUseCse,

		api.NewServerHttp)
	return &api.ServerHTTP{}, nil
}
