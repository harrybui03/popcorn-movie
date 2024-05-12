package show_time

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/room"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	ListShowTimes(ctx context.Context, input model.ListShowTimeInput) ([]*ent.ShowTime, int, error)
}

type impl struct {
	logger     *zap.Logger
	repository repository.Registry
	config     config.AppConfig
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.AppConfig) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		config:     appConfig,
	}
}

func (i impl) ListShowTimes(ctx context.Context, input model.ListShowTimeInput) ([]*ent.ShowTime, int, error) {
	query := i.repository.ShowTime().ShowTimeQuery().WithMovie().WithRoom(func(roomQuery *ent.RoomQuery) {
		roomQuery.WithTheater()
	})
	if input.Filter != nil {
		if input.Filter.TheaterID != nil {
			theaterID, err := uuid.Parse(*input.Filter.TheaterID)
			if err != nil {
				i.logger.Error(err.Error())
				return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
			}

			roomIDs, err := i.repository.Room().RoomQuery().Where(room.TheaterID(theaterID)).IDs(ctx)
			if err != nil {
				i.logger.Error(err.Error())
				return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
			}
			query.Where(showtime.HasRoomWith(room.IDIn(roomIDs...)))
		}

		if input.Filter.MovieID != nil {
			movieID, err := uuid.Parse(*input.Filter.MovieID)
			if err != nil {
				i.logger.Error(err.Error())
				return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
			}

			query.Where(showtime.MovieID(movieID))
		}

	}

	count, err := i.repository.ShowTime().CountShowTime(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	showTimes, err := i.repository.ShowTime().GetAllShowTime(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return showTimes, *count, nil
}
