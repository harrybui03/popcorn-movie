package service

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/repository"
	"PopcornMovie/service/auth"
	"PopcornMovie/service/user"
	"go.uber.org/zap"
)

type Registry interface {
	User() user.Service
	Auth() auth.Service
}

type impl struct {
	user user.Service
	auth auth.Service
}

func New(entClient *ent.Client, logger *zap.Logger, appConfig config.AppConfig) Registry {
	repositoryRegistry := repository.New(entClient)

	return &impl{
		user: user.New(repositoryRegistry, logger, appConfig),
		auth: auth.New(repositoryRegistry, logger, appConfig),
	}
}

func (i impl) User() user.Service {
	//TODO implement me
	panic("implement me")
}

func (i impl) Auth() auth.Service {
	return i.auth
}
