package room

import (
	"PopcornMovie/ent"
	"context"
	"github.com/pkg/errors"
)

// Repository is room repo
type Repository interface {
	RoomQuery() *ent.RoomQuery
	GetAllRooms(ctx context.Context, query *ent.RoomQuery) ([]*ent.Room, error)
	CountRooms(ctx context.Context, query *ent.RoomQuery) (*int, error)
}

// impl is implement Repository service
type impl struct {
	client *ent.Client
}

func (i impl) RoomQuery() *ent.RoomQuery {
	return i.client.Room.Query()
}

func (i impl) GetAllRooms(ctx context.Context, query *ent.RoomQuery) ([]*ent.Room, error) {
	rooms, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return rooms, nil
}

func (i impl) CountRooms(ctx context.Context, query *ent.RoomQuery) (*int, error) {
	cnt, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &cnt, nil
}

// New is function init Room Repo
func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
