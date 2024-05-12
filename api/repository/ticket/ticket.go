package ticket

import (
	"PopcornMovie/ent"
	"PopcornMovie/model"
	"context"
	"github.com/pkg/errors"
)

type Repository interface {
	TicketQuery() *ent.TicketQuery
	TicketCreate() *ent.TicketCreate
	CreateTicket(ctx context.Context, input model.CreateTicket) (*ent.Ticket, error)
	CreateBulkTicket(ctx context.Context, input []*ent.TicketCreate) ([]*ent.Ticket, error)
	CountTickets(ctx context.Context, query *ent.TicketQuery) (int, error)
	GetAllTickets(ctx context.Context, query *ent.TicketQuery) ([]*ent.Ticket, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) TicketCreate() *ent.TicketCreate {
	return i.client.Ticket.Create()
}

func (i impl) CreateBulkTicket(ctx context.Context, input []*ent.TicketCreate) ([]*ent.Ticket, error) {
	tickets, err := i.client.Ticket.CreateBulk(input...).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return tickets, nil
}

func (i impl) CountTickets(ctx context.Context, query *ent.TicketQuery) (int, error) {
	return i.client.Ticket.Query().Count(ctx)
}

func (i impl) GetAllTickets(ctx context.Context, query *ent.TicketQuery) ([]*ent.Ticket, error) {
	return query.All(ctx)
}

func (i impl) CreateTicket(ctx context.Context, input model.CreateTicket) (*ent.Ticket, error) {
	ticket, err := i.client.Ticket.Create().SetPrice(input.Price).SetSeatID(input.SeatID).SetTransactionID(input.TransactionID).SetShowTimeID(input.ShowTimeID).SetIsBooked(input.IsBooked).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return ticket, nil
}

func (i impl) TicketQuery() *ent.TicketQuery {
	return i.client.Ticket.Query()
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
