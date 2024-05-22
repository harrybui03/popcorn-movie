package resolver

import (
	"PopcornMovie/cmd/middleware"
	generated "PopcornMovie/graphql"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/service"
	"context"
	"github.com/99designs/gqlgen/graphql"
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
func NewExecutableSchema(service service.Registry, logger *zap.Logger) graphql.ExecutableSchema {
	config := generated.Config{
		Resolvers: &Resolver{service: service, logger: logger},
	}

	config.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		payload := middleware.GetPayload(ctx)
		if payload == nil {
			return nil, utils.WrapGQLError(ctx, string(utils.ErrorUnauthorizedRequest), utils.ErrorCodeUnauthorized)
		}
		return next(ctx)
	}

	config.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []model.Role) (res interface{}, err error) {
		tokenData := middleware.GetPayload(ctx)
		if tokenData == nil {
			return nil, utils.WrapGQLError(ctx, string(utils.ErrorUnauthorizedRequest), utils.ErrorCodeUnauthorized)
		}

		if !utils.Contains(roles, model.Role(tokenData.Role)) {
			return nil, utils.WrapGQLError(ctx, string(utils.ErrorUnauthorizedRequest), utils.ErrorCodeUnauthorized)
		}

		return next(ctx)
	}
	return generated.NewExecutableSchema(config)
}
