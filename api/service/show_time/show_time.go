package show_time

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/movie"
	"PopcornMovie/ent/room"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/ticket"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Service interface {
	ListShowTimes(ctx context.Context, input model.ListShowTimeInput) ([]*ent.ShowTime, int, error)
	CreateShowTime(ctx context.Context, input model.CreateShowTimeInput) (*ent.ShowTime, error)
	UpdateShowTime(ctx context.Context, input model.UpdateShowTimeInput) (*ent.ShowTime, error)
	DeleteShowTime(ctx context.Context, id string) (string, error)
}

type impl struct {
	logger     *zap.Logger
	repository repository.Registry
	config     config.Configurations
}

func (i impl) CreateShowTime(ctx context.Context, input model.CreateShowTimeInput) (*ent.ShowTime, error) {
	movieID, err := uuid.Parse(input.MovieID)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	exist, _ := i.repository.Movie().MovieQuery().Where(movie.ID(movieID)).Exist(ctx)
	if !exist {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "Movie"), utils.ErrorCodeNotFound)
	}

	roomID, err := uuid.Parse(input.RoomID)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	exist, _ = i.repository.Room().RoomQuery().Where(room.ID(roomID)).Exist(ctx)
	if !exist {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "Room"), utils.ErrorCodeNotFound)
	}

	// validate time format
	exist, _ = i.repository.ShowTime().ShowTimeQuery().Where(showtime.StartAtGTE(input.StartAt), showtime.EndAtLTE(input.EndAt)).Exist(ctx)
	if exist {
		return nil, utils.WrapGQLError(ctx, "Time is not available", utils.ErrorCodeBadRequest)
	}

	showtimeRecord, err := i.repository.ShowTime().ShowTimeCreate().SetMovieID(movieID).SetRoomID(roomID).SetStartAt(input.StartAt).SetEndAt(input.EndAt).Save(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return showtimeRecord, nil
}

func (i impl) UpdateShowTime(ctx context.Context, input model.UpdateShowTimeInput) (*ent.ShowTime, error) {
	showtimeID, err := uuid.Parse(input.ID)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	showtimeRecord, _ := i.repository.ShowTime().ShowTimeQuery().Where(showtime.IDEQ(showtimeID)).First(ctx)
	if showtimeRecord == nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "Showtime"), utils.ErrorCodeNotFound)
	}

	showtimeQuery := showtimeRecord.Update()
	if input.MovieID != nil {
		movieID, err := uuid.Parse(*input.MovieID)
		if err != nil {
			i.logger.Error(err.Error())
			return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
		}

		exist, _ := i.repository.Movie().MovieQuery().Where(movie.ID(movieID)).Exist(ctx)
		if !exist {
			i.logger.Error(err.Error())
			return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "Movie"), utils.ErrorCodeNotFound)
		}
		showtimeQuery.SetMovieID(movieID)
	}

	if input.RoomID != nil {
		roomID, err := uuid.Parse(*input.RoomID)
		if err != nil {
			i.logger.Error(err.Error())
			return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
		}

		exist, _ := i.repository.Room().RoomQuery().Where(room.ID(roomID)).Exist(ctx)
		if !exist {
			i.logger.Error(err.Error())
			return nil, utils.WrapGQLError(ctx, fmt.Sprintf(string(utils.ErrorNotFound), "Room"), utils.ErrorCodeNotFound)
		}
		showtimeQuery.SetRoomID(roomID)
	}

	if input.StartAt != nil {
		showtimeQuery.SetStartAt(*input.StartAt)
	}

	if input.EndAt != nil {
		showtimeQuery.SetEndAt(*input.EndAt)
	}

	updatedRecord, err := showtimeQuery.Save(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return updatedRecord, nil
}

func (i impl) DeleteShowTime(ctx context.Context, id string) (string, error) {
	showtimeID, err := uuid.Parse(id)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	exist, _ := i.repository.Ticket().TicketQuery().Where(ticket.ShowTimeID(showtimeID), ticket.IsBooked(true)).Exist(ctx)
	if exist {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, "Cannot delete showtime with booked ticket", utils.ErrorCodeBadRequest)
	}

	_, err = i.repository.ShowTime().ShowTimeDelete().Where(showtime.ID(showtimeID)).Exec(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return "", utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return "Showtime deleted successfully", nil
}

func New(repository repository.Registry, logger *zap.Logger, appConfig config.Configurations) Service {
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

		if input.Filter.Date != nil {
			startDay := *input.Filter.Date
			endDay := startDay.AddDate(0, 0, 1)
			query.Where(showtime.StartAtGTE(startDay), showtime.EndAtLTE(endDay))
		}

		if input.Filter.MovieID != nil {
			movieID, err := uuid.Parse(*input.Filter.MovieID)
			if err != nil {
				i.logger.Error(err.Error())
				return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
			}

			query.Where(showtime.MovieID(movieID))
		}

		if input.Filter.RoomID != nil {
			roomID, err := uuid.Parse(*input.Filter.RoomID)
			if err != nil {
				i.logger.Error(err.Error())
				return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
			}

			query.Where(showtime.RoomID(roomID))
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
