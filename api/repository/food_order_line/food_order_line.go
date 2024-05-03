package food_order_line

import (
	"PopcornMovie/ent"
	"PopcornMovie/model"
	"context"
	"github.com/pkg/errors"
)

type Repository interface {
	FoodOrderLineQuery() *ent.FoodOrderLineQuery
	CreateFoodOrder(ctx context.Context, input model.CreateFoodOrderLine) (*ent.FoodOrderLine, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) CreateFoodOrder(ctx context.Context, input model.CreateFoodOrderLine) (*ent.FoodOrderLine, error) {
	foodOrder, err := i.client.FoodOrderLine.Create().SetFoodID(input.FoodID).SetQuantity(input.Quantity).SetTransactionID(input.TransactionID).Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return foodOrder, nil
}

func (i impl) FoodOrderLineQuery() *ent.FoodOrderLineQuery {
	//TODO implement me
	panic("implement me")
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
