package user

import (
	"PopcornMovie/config"
	"PopcornMovie/repository"
	"go.uber.org/zap"
)

// Service is the interface for the user service.
type Service interface {
}

// impl is the implementation of the user service.
type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.AppConfig
}

// New creates a new user service.
func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}
