package movie

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/movie"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	ListMovies(ctx context.Context, input model.ListMovieInput) ([]*ent.Movie, int, error)
	GetMovieByID(ctx context.Context, id string) (*ent.Movie, error)
	CreateMovie(ctx context.Context, input model.CreateMovieInput) (*ent.Movie, error)
	UpdateMovie(ctx context.Context, input model.UpdateMovieInput) (*ent.Movie, error)
	DeleteMovie(ctx context.Context, id string) (string, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.Configurations
}

func (i impl) UpdateMovie(ctx context.Context, input model.UpdateMovieInput) (*ent.Movie, error) {
	// Find Movie
	movieID, err := uuid.Parse(input.ID)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeBadRequest)
	}

	movieRecord, _ := i.repository.Movie().GetMovieByID(ctx, movieID)
	if movieRecord == nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "Movie"), utils.ErrorCodeNotFound)
	}

	// Update Movie
	movieUpdate := movieRecord.Update()

	if input.Status != nil {
		movieUpdate.SetStatus(movie.Status(*input.Status))
	}

	if input.Title != nil {
		movieUpdate.SetTitle(*input.Title)
	}

	if input.Cast != nil {
		movieUpdate.SetCast(*input.Cast)
	}

	if input.Director != nil {
		movieUpdate.SetDirector(*input.Director)
	}

	if input.Duration != nil {
		movieUpdate.SetDuration(*input.Duration)
	}

	if input.Genre != nil {
		movieUpdate.SetGenre(*input.Genre)

	}

	if input.Language != nil {
		movieUpdate.SetLanguage(*input.Language)
	}

	if input.OpeningDay != nil {
		movieUpdate.SetOpeningDay(*input.OpeningDay)
	}

	if input.Poster != nil {
		movieUpdate.SetPoster(*input.Poster)
	}

	if input.Rated != nil {
		movieUpdate.SetRated(*input.Rated)
	}

	if input.Story != nil {
		movieUpdate.SetStory(*input.Story)
	}

	if input.Trailer != nil {
		movieUpdate.SetTrailer(*input.Trailer)
	}

	movieRecordUpdated, err := movieUpdate.Save(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return movieRecordUpdated, nil
}

func (i impl) DeleteMovie(ctx context.Context, id string) (string, error) {
	movieID, err := uuid.Parse(id)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeBadRequest)

	}
	_, err = i.repository.Movie().MovieDelete().Where(movie.ID(movieID)).Exec(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return "Movie deleted successfully", nil
}

func (i impl) CreateMovie(ctx context.Context, input model.CreateMovieInput) (*ent.Movie, error) {
	movieRecord, err := i.repository.
		Movie().MovieCreate().
		SetStatus(movie.Status(input.Status)).
		SetTitle(input.Title).
		SetCast(input.Cast).
		SetDirector(input.Director).
		SetDuration(input.Duration).
		SetGenre(input.Genre).
		SetLanguage(input.Language).
		SetOpeningDay(input.OpeningDay).
		SetPoster(input.Poster).
		SetPoster(input.Poster).
		SetRated(input.Rated).
		SetStory(input.Story).
		SetTrailer(input.Trailer).Save(ctx)

	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return movieRecord, nil
}

func (i impl) GetMovieByID(ctx context.Context, id string) (*ent.Movie, error) {
	idUUID, err := uuid.Parse(id)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeBadRequest)
	}
	movieRecord, err := i.repository.Movie().GetMovieByID(ctx, idUUID)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return movieRecord, nil
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

func New(repository repository.Registry, logger *zap.Logger, appConfig config.Configurations) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}
