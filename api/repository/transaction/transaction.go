package transaction

import (
	"PopcornMovie/ent"
	"PopcornMovie/model"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repository interface {
	TransactionQuery() *ent.TransactionQuery
	CreateTransaction(ctx context.Context, userId uuid.UUID) (*ent.Transaction, error)
	GetAllTransactions(ctx context.Context, query *ent.TransactionQuery) ([]*ent.Transaction, error)
	CountTransactions(ctx context.Context, query *ent.TransactionQuery) (int, error)
	GetRevenue(ctx context.Context, input model.RevenueInput) (*model.MonthlyRevenueOutput, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) GetRevenue(ctx context.Context, input model.RevenueInput) (*model.MonthlyRevenueOutput, error) {
	//i.client.
	return nil, nil
}

func (i impl) CountTransactions(ctx context.Context, query *ent.TransactionQuery) (int, error) {
	return query.Count(ctx)
}

func (i impl) GetAllTransactions(ctx context.Context, query *ent.TransactionQuery) ([]*ent.Transaction, error) {
	return query.All(ctx)
}

func (i impl) CreateTransaction(ctx context.Context, userId uuid.UUID) (*ent.Transaction, error) {
	transaction, err := i.client.Transaction.Create().
		SetUserID(userId).
		SetTotal(0).
		Save(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return transaction, nil
}

func (i impl) TransactionQuery() *ent.TransactionQuery {
	return i.client.Transaction.Query()
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
