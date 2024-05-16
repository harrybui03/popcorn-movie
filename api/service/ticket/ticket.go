package ticket

import (
	"PopcornMovie/config"
	"PopcornMovie/ent"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/ticket"
	"PopcornMovie/internal/utils"
	"PopcornMovie/model"
	"PopcornMovie/repository"
	"context"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"sort"
	"strconv"
)

type Service interface {
	GetAllTickets(ctx context.Context, input model.ListTicketInput) ([]*ent.Ticket, int, error)
	GenerateTickets(ctx context.Context, input model.GenerateTicketInput) ([]*ent.Ticket, error)
}

type impl struct {
	repository repository.Registry
	logger     *zap.Logger
	appConfig  config.Configurations
}

func (i impl) GenerateTickets(ctx context.Context, input model.GenerateTicketInput) ([]*ent.Ticket, error) {
	showTimeID, err := uuid.Parse(input.ShowTimeID)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeBadRequest)
	}

	showTimeRecord, err := i.repository.ShowTime().ShowTimeQuery().WithRoom().Where(showtime.IDEQ(showTimeID)).First(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	roomRecord, err := showTimeRecord.QueryRoom().WithSeats().First(ctx)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	// Create tickets
	tickets := make([]*ent.TicketCreate, 0)
	for _, seat := range roomRecord.Edges.Seats {
		price := input.Price
		if model.SeatCategory(seat.Category) == model.SeatCategoryDouble {
			price = price * 2
		}
		ticketCreate := i.repository.Ticket().TicketCreate()
		ticketCreate.SetShowTimeID(showTimeID)
		ticketCreate.SetSeatID(seat.ID)
		ticketCreate.SetPrice(price)
		ticketCreate.SetIsBooked(false)
		tickets = append(tickets, ticketCreate)
	}

	createdTickets, err := i.repository.Ticket().CreateBulkTicket(ctx, tickets)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	return createdTickets, nil
}

func (i impl) GetAllTickets(ctx context.Context, input model.ListTicketInput) ([]*ent.Ticket, int, error) {
	query := i.repository.Ticket().TicketQuery().WithShowTime().WithSeat()
	if input.Filter != nil {
		showtimeID, err := uuid.Parse(input.Filter.ShowTimeID)
		if err != nil {
			i.logger.Error(err.Error())
			return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeBadRequest)
		}
		query.Where(ticket.HasShowTimeWith(showtime.ID(showtimeID)))
	}

	if input.Pagination != nil {
		offset := utils.CalculateOffset(input.Pagination.Page, input.Pagination.Limit)
		query.Limit(input.Pagination.Limit).Offset(offset)
	}

	count, err := i.repository.Ticket().CountTickets(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	tickets, err := i.repository.Ticket().GetAllTickets(ctx, query)
	if err != nil {
		i.logger.Error(err.Error())
		return nil, 0, utils.WrapGQLError(ctx, string(utils.ErrorMessageInternal), utils.ErrorCodeInternal)
	}

	sort.Slice(tickets, func(i, j int) bool {
		// Extract the seat numbers
		seatNumberI := tickets[i].Edges.Seat.SeatNumber
		seatNumberJ := tickets[j].Edges.Seat.SeatNumber

		// Extract the row and column parts
		rowI := seatNumberI[0] // First character (e.g., 'A')
		rowJ := seatNumberJ[0] // First character (e.g., 'A')

		colI, _ := strconv.Atoi(seatNumberI[1:]) // Remaining part after the first character (e.g., '1')
		colJ, _ := strconv.Atoi(seatNumberJ[1:]) // Remaining part after the first character (e.g., '1')

		// Compare rows first
		if rowI != rowJ {
			return rowI < rowJ
		}

		// If rows are the same, compare columns
		return colI < colJ
	})

	return tickets, count, nil
}

func New(r repository.Registry, l *zap.Logger, c config.Configurations) Service {
	return &impl{
		repository: r,
		logger:     l,
		appConfig:  c,
	}
}
