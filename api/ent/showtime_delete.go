// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/showtime"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShowTimeDelete is the builder for deleting a ShowTime entity.
type ShowTimeDelete struct {
	config
	hooks    []Hook
	mutation *ShowTimeMutation
}

// Where appends a list predicates to the ShowTimeDelete builder.
func (std *ShowTimeDelete) Where(ps ...predicate.ShowTime) *ShowTimeDelete {
	std.mutation.Where(ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *ShowTimeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, std.sqlExec, std.mutation, std.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (std *ShowTimeDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *ShowTimeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(showtime.Table, sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID))
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, std.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	std.mutation.done = true
	return affected, err
}

// ShowTimeDeleteOne is the builder for deleting a single ShowTime entity.
type ShowTimeDeleteOne struct {
	std *ShowTimeDelete
}

// Where appends a list predicates to the ShowTimeDelete builder.
func (stdo *ShowTimeDeleteOne) Where(ps ...predicate.ShowTime) *ShowTimeDeleteOne {
	stdo.std.mutation.Where(ps...)
	return stdo
}

// Exec executes the deletion query.
func (stdo *ShowTimeDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{showtime.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *ShowTimeDeleteOne) ExecX(ctx context.Context) {
	if err := stdo.Exec(ctx); err != nil {
		panic(err)
	}
}
