package room

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/room"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/theater"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// Service is an interface for Room Service
type Service interface {
	ListRooms(ctx context.Context, input model.ListRoomInput) ([]*ent.Room, int, error)
}

// impl is implement for Room Service
type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.Configurations
}

func (i impl) ListRooms(ctx context.Context, input model.ListRoomInput) ([]*ent.Room, int, error) {
	query := i.repository.Room().RoomQuery().WithTheater().WithShowTimes().WithSeats()
	if input.Filter != nil {
		theaterID, err := uuid.Parse(input.Filter.TheaterID)
		if err != nil {
			i.logger.Error(err.Error())
			return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
		}

		query.Where(room.HasTheaterWith(theater.ID(theaterID)))

		if input.Filter.StartAt != nil {
			query.Where(room.HasShowTimesWith(showtime.StartAtGTE(*input.Filter.StartAt)))
		}

		if input.Filter.EndAt != nil {
			query.Where(room.HasShowTimesWith(showtime.EndAtLTE(*input.Filter.EndAt)))
		}

		if input.Filter.ShowTimeID != nil {
			showTimeID, err := uuid.Parse(*input.Filter.ShowTimeID)
			if err != nil {
				i.logger.Error(err.Error())
				return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
			}

			query.Where(room.HasShowTimesWith(showtime.ID(showTimeID)))
		}
	}

	count, err := i.repository.Room().CountRooms(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	rooms, err := i.repository.Room().GetAllRooms(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return rooms, *count, nil
}

// New is function init Service Rooms
func New(repository repository.Registry, logger *zap.Logger, appConfig config.Configurations) Service {
	return &impl{
		repository: repository,
		logger:     logger,
		appConfig:  appConfig,
	}
}
