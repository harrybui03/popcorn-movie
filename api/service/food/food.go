package food

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"go.uber.org/zap"
)

type Service interface {
	ListFood(ctx context.Context, input model.ListFoodInput) ([]*ent.Food, int, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.Configurations
}

func (i impl) ListFood(ctx context.Context, input model.ListFoodInput) ([]*ent.Food, int, error) {
	query := i.repository.Food().FoodQuery()

	count, err := i.repository.Food().CountFoods(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	food, err := i.repository.Food().GetAllFoods(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return food, *count, nil
}

// New is function init Service Rooms
func New(repository repository.Registry, logger *zap.Logger, appConfig config.Configurations) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}
