package service

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/repository"
	"PopcornMovie/service/auth"
	"PopcornMovie/service/theater"
	"PopcornMovie/service/user"
	"go.uber.org/zap"
)

type Registry interface {
	User() user.Service
	Auth() auth.Service
	Theater() theater.Service
}

type impl struct {
	user    user.Service
	auth    auth.Service
	theater theater.Service
}

func (i impl) Theater() theater.Service { return i.theater }

func (i impl) User() user.Service {
	//TODO implement me
	return i.user
}

func (i impl) Auth() auth.Service {
	return i.auth
}

func New(entClient *ent.Client, logger *zap.Logger, appConfig config.AppConfig) Registry {
	repositoryRegistry := repository.New(entClient)

	return &impl{
		user:    user.New(repositoryRegistry, logger, appConfig),
		auth:    auth.New(repositoryRegistry, logger, appConfig),
		theater: theater.New(repositoryRegistry, logger, appConfig),
	}
}
