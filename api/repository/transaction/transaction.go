package transaction

import (
	"PopcornMovie/ent"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repository interface {
	TransactionQuery() *ent.TransactionQuery
	CreateTransaction(ctx context.Context, userId uuid.UUID) (*ent.Transaction, error)
}

type impl struct {
	client *ent.Client
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
	//TODO implement me
	panic("implement me")
}

func New(client *ent.Client) Repository {
	return &impl{
		client: client,
	}
}
