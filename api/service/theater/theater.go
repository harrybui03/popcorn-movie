package theater

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/theater"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"go.uber.org/zap"
)

// Service is define for theater interface
type Service interface {
	ListTheaters(ctx context.Context, input model.ListTheatersInput) ([]*ent.Theater, int, error)
}

// impl is implement for theater service
type impl struct {
	logger     *zap.Logger
	repository repository.Registry
	config     config.AppConfig
}

func (i impl) ListTheaters(ctx context.Context, input model.ListTheatersInput) ([]*ent.Theater, int, error) {
	query := i.repository.Theater().TheaterQuery()
	if input.Filter != nil {
		if input.Filter.Name != "" {
			query.Where(theater.NameContains(input.Filter.Name))
		}

		if input.Filter.Address != "" {
			query.Where(theater.AddressContains(input.Filter.Address))
		}
	}

	count, err := i.repository.Theater().CountTheaters(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, err
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	theaters, err := i.repository.Theater().GetAllTheaters(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, err
	}

	return theaters, *count, nil
}

// New is a function init Theater Service
func New(registry repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: registry,
		logger:     logger,
		config:     appConfig,
	}
}
