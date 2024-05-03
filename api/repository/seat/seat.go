package seat

import (
	"PopcornMovie/ent"
	"context"
	"github.com/pkg/errors"
)

type Repository interface {
	SeatQuery() *ent.SeatQuery
	GetAllSeats(ctx context.Context, query *ent.SeatQuery) ([]*ent.Seat, error)
	CountSeats(ctx context.Context, query *ent.SeatQuery) (*int, error)
}

func (i impl) SeatQuery() *ent.SeatQuery {
	return i.client.Seat.Query()
}

func (i impl) GetAllSeats(ctx context.Context, query *ent.SeatQuery) ([]*ent.Seat, error) {
	seats, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return seats, nil
}

func (i impl) CountSeats(ctx context.Context, query *ent.SeatQuery) (*int, error) {
	count, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &count, nil
}

type impl struct {
	client *ent.Client
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
