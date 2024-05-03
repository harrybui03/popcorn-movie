package movie

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/movie"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"go.uber.org/zap"
)

type Service interface {
	ListMovies(ctx context.Context, input model.ListMovieInput) ([]*ent.Movie, int, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.AppConfig
}

func (i impl) ListMovies(ctx context.Context, input model.ListMovieInput) ([]*ent.Movie, int, error) {
	query := i.repository.Movie().MovieQuery()
	if input.Filter != nil {
		if input.Filter.Status != nil {
			query.Where(movie.StatusEQ(movie.Status(*input.Filter.Status)))
		}
	}

	count, err := i.repository.Movie().CountMovies(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	movies, err := i.repository.Movie().GetAllMovie(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return movies, *count, nil
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}
