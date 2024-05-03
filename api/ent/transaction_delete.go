// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/transaction"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionDelete is the builder for deleting a Transaction entity.
type TransactionDelete struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionDelete builder.
func (td *TransactionDelete) Where(ps ...predicate.Transaction) *TransactionDelete {
	td.mutation.Where(ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TransactionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, td.sqlExec, td.mutation, td.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TransactionDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TransactionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(transaction.Table, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID))
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, td.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	td.mutation.done = true
	return affected, err
}

// TransactionDeleteOne is the builder for deleting a single Transaction entity.
type TransactionDeleteOne struct {
	td *TransactionDelete
}

// Where appends a list predicates to the TransactionDelete builder.
func (tdo *TransactionDeleteOne) Where(ps ...predicate.Transaction) *TransactionDeleteOne {
	tdo.td.mutation.Where(ps...)
	return tdo
}

// Exec executes the deletion query.
func (tdo *TransactionDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{transaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TransactionDeleteOne) ExecX(ctx context.Context) {
	if err := tdo.Exec(ctx); err != nil {
		panic(err)
	}
}