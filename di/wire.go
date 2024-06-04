//go:build wireinject
// +build wireinject

package di

import (
	"awesomeProject/api"
	"awesomeProject/api/handler"
	"awesomeProject/config"
	"awesomeProject/db"
	"awesomeProject/repository"
	"awesomeProject/usecase"
	"awesomeProject/utils/logger"
	"github.com/google/wire"
)

func InitServer() (*api.HttpServer, error) {
	wire.Build(
		api.ProvideAddrString,
		config.GetConfig,
		db.ConnectDb,
		logger.NewLogger,
		repository.NewForestRepo,
		usecase.NewForestService,

		handler.NewForestHandler,
		api.NewHttpServer,
	)
	return nil, nil

}
func Co() *config.Config {
	return &config.Config{}
}
