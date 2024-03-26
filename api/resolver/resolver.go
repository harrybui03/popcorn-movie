package resolver

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	generated "PopcornMovie/graphql"
	"PopcornMovie/service"
	"github.com/99designs/gqlgen/graphql"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service service.Registry
	logger  *zap.Logger
}

// NewExecutableSchema creates an ExecutableSchema instance.
func NewExecutableSchema(client *ent.Client, validator *validator.Validate, validationTranslator ut.Translator, logger *zap.Logger, appConfig config.AppConfig) graphql.ExecutableSchema {
	service := service.New(client, logger, appConfig)

	config := generated.Config{
		Resolvers: &Resolver{service: service, logger: logger},
	}

	return generated.NewExecutableSchema(config)
}
