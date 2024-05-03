package ticket

import (
	"PopcornMovie/ent"
	"PopcornMovie/model"
	"context"
	"github.com/pkg/errors"
)

type Repository interface {
	TicketQuery() *ent.TicketQuery
	CreateTicket(ctx context.Context, input model.CreateTicket) (*ent.Ticket, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) CreateTicket(ctx context.Context, input model.CreateTicket) (*ent.Ticket, error) {
	ticket, err := i.client.Ticket.Create().SetPrice(input.Price).SetSeatID(input.SeatID).SetTransactionID(input.TransactionID).SetShowTimeID(input.ShowTimeID).SetIsBooked(input.IsBooked).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ticket, nil
}

func (i impl) TicketQuery() *ent.TicketQuery {
	//TODO implement me
	panic("implement me")
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
