package show_time

import (
	"PopcornMovie/ent"
	"context"
)

type Repository interface {
	ShowTimeQuery() *ent.ShowTimeQuery
	GetAllShowTime(ctx context.Context, query *ent.ShowTimeQuery) ([]*ent.ShowTime, error)
	CountShowTime(ctx context.Context, query *ent.ShowTimeQuery) (*int, error)
	ShowTimeCreate() *ent.ShowTimeCreate
	ShowTimeDelete() *ent.ShowTimeDelete
}

func (i impl) ShowTimeQuery() *ent.ShowTimeQuery {
	return i.client.ShowTime.Query()
}

func (i impl) GetAllShowTime(ctx context.Context, query *ent.ShowTimeQuery) ([]*ent.ShowTime, error) {
	showTimes, err := query.All(ctx)
	if err != nil {
		return nil, err
	}

	return showTimes, nil
}

func (i impl) CountShowTime(ctx context.Context, query *ent.ShowTimeQuery) (*int, error) {
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}

type impl struct {
	client *ent.Client
}

func (i impl) ShowTimeCreate() *ent.ShowTimeCreate {
	return i.client.ShowTime.Create()
}

func (i impl) ShowTimeDelete() *ent.ShowTimeDelete {
	return i.client.ShowTime.Delete()
}
