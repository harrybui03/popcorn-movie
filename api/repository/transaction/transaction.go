package transaction

import (
	"PopcornMovie/ent"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Repository interface {
	TransactionQuery() *ent.TransactionQuery
	CreateTransaction(ctx context.Context, userId uuid.UUID) (*ent.Transaction, error)
	GetAllTransactions(ctx context.Context, query *ent.TransactionQuery) ([]*ent.Transaction, error)
	CountTransactions(ctx context.Context, query *ent.TransactionQuery) (int, error)
	GetRevenue(ctx context.Context, year, month int) (float64, error)
}

type impl struct {
	client *ent.Client
}

func (i impl) GetRevenue(ctx context.Context, year, month int) (float64, error) {
	var total float64
	query := fmt.Sprintf(`SELECT COALESCE(SUM(total), 0) AS total_sum FROM public.transactions WHERE EXTRACT(YEAR FROM created_at) = %d AND EXTRACT(MONTH FROM created_at) = %d;`, year, month)
	rows, err := i.client.QueryContext(ctx, query)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	for rows.Next() {
		if err = rows.Scan(&total); err != nil {
			return 0, errors.WithStack(err)
		}
	}

	return total, err
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
