package theater

import (
	"PopcornMovie/ent"
	"context"
	"github.com/pkg/errors"
)

type Repository interface {
	TheaterQuery() *ent.TheaterQuery
	GetAllTheaters(ctx context.Context, query *ent.TheaterQuery) ([]*ent.Theater, error)
}

func (i impl) TheaterQuery() *ent.TheaterQuery {
	return i.client.Theater.Query()
}

func (i impl) GetAllTheaters(ctx context.Context, query *ent.TheaterQuery) ([]*ent.Theater, error) {
	theaters, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return theaters, nil
}

type impl struct {
	client *ent.Client
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
