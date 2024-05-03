package food

import (
	"PopcornMovie/ent"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repository interface {
	FoodQuery() *ent.FoodQuery
	GetAllFoods(ctx context.Context, query *ent.FoodQuery) ([]*ent.Food, error)
	CountFoods(ctx context.Context, query *ent.FoodQuery) (*int, error)
	GetFoodByID(ctx context.Context, id uuid.UUID) (*ent.Food, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) GetFoodByID(ctx context.Context, id uuid.UUID) (*ent.Food, error) {
	food, err := i.client.Food.Get(ctx, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return food, nil
}

func (i impl) FoodQuery() *ent.FoodQuery {
	return i.client.Food.Query()
}

func (i impl) GetAllFoods(ctx context.Context, query *ent.FoodQuery) ([]*ent.Food, error) {
	foods, err := query.All(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return foods, nil
}

func (i impl) CountFoods(ctx context.Context, query *ent.FoodQuery) (*int, error) {
	cnt, err := query.Count(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &cnt, nil
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
