// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/movie"
	"PopcornMovie/ent/room"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/ticket"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ShowTimeCreate is the builder for creating a ShowTime entity.
type ShowTimeCreate struct {
	config
	mutation *ShowTimeMutation
	hooks    []Hook
}

// SetStartAt sets the "start_at" field.
func (stc *ShowTimeCreate) SetStartAt(t time.Time) *ShowTimeCreate {
	stc.mutation.SetStartAt(t)
	return stc
}

// SetEndAt sets the "end_at" field.
func (stc *ShowTimeCreate) SetEndAt(t time.Time) *ShowTimeCreate {
	stc.mutation.SetEndAt(t)
	return stc
}

// SetMovieID sets the "movie_id" field.
func (stc *ShowTimeCreate) SetMovieID(u uuid.UUID) *ShowTimeCreate {
	stc.mutation.SetMovieID(u)
	return stc
}

// SetRoomID sets the "room_id" field.
func (stc *ShowTimeCreate) SetRoomID(u uuid.UUID) *ShowTimeCreate {
	stc.mutation.SetRoomID(u)
	return stc
}

// SetCreatedAt sets the "created_at" field.
func (stc *ShowTimeCreate) SetCreatedAt(t time.Time) *ShowTimeCreate {
	stc.mutation.SetCreatedAt(t)
	return stc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (stc *ShowTimeCreate) SetNillableCreatedAt(t *time.Time) *ShowTimeCreate {
	if t != nil {
		stc.SetCreatedAt(*t)
	}
	return stc
}

// SetUpdatedAt sets the "updated_at" field.
func (stc *ShowTimeCreate) SetUpdatedAt(t time.Time) *ShowTimeCreate {
	stc.mutation.SetUpdatedAt(t)
	return stc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (stc *ShowTimeCreate) SetNillableUpdatedAt(t *time.Time) *ShowTimeCreate {
	if t != nil {
		stc.SetUpdatedAt(*t)
	}
	return stc
}

// SetID sets the "id" field.
func (stc *ShowTimeCreate) SetID(u uuid.UUID) *ShowTimeCreate {
	stc.mutation.SetID(u)
	return stc
}

// SetRoom sets the "room" edge to the Room entity.
func (stc *ShowTimeCreate) SetRoom(r *Room) *ShowTimeCreate {
	return stc.SetRoomID(r.ID)
}

// SetMovie sets the "movie" edge to the Movie entity.
func (stc *ShowTimeCreate) SetMovie(m *Movie) *ShowTimeCreate {
	return stc.SetMovieID(m.ID)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (stc *ShowTimeCreate) AddTicketIDs(ids ...uuid.UUID) *ShowTimeCreate {
	stc.mutation.AddTicketIDs(ids...)
	return stc
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (stc *ShowTimeCreate) AddTickets(t ...*Ticket) *ShowTimeCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stc.AddTicketIDs(ids...)
}

// Mutation returns the ShowTimeMutation object of the builder.
func (stc *ShowTimeCreate) Mutation() *ShowTimeMutation {
	return stc.mutation
}

// Save creates the ShowTime in the database.
func (stc *ShowTimeCreate) Save(ctx context.Context) (*ShowTime, error) {
	stc.defaults()
	return withHooks(ctx, stc.sqlSave, stc.mutation, stc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (stc *ShowTimeCreate) SaveX(ctx context.Context) *ShowTime {
	v, err := stc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stc *ShowTimeCreate) Exec(ctx context.Context) error {
	_, err := stc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stc *ShowTimeCreate) ExecX(ctx context.Context) {
	if err := stc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stc *ShowTimeCreate) defaults() {
	if _, ok := stc.mutation.CreatedAt(); !ok {
		v := showtime.DefaultCreatedAt()
		stc.mutation.SetCreatedAt(v)
	}
	if _, ok := stc.mutation.UpdatedAt(); !ok {
		v := showtime.DefaultUpdatedAt()
		stc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stc *ShowTimeCreate) check() error {
	if _, ok := stc.mutation.StartAt(); !ok {
		return &ValidationError{Name: "start_at", err: errors.New(`ent: missing required field "ShowTime.start_at"`)}
	}
	if _, ok := stc.mutation.EndAt(); !ok {
		return &ValidationError{Name: "end_at", err: errors.New(`ent: missing required field "ShowTime.end_at"`)}
	}
	if _, ok := stc.mutation.MovieID(); !ok {
		return &ValidationError{Name: "movie_id", err: errors.New(`ent: missing required field "ShowTime.movie_id"`)}
	}
	if _, ok := stc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "room_id", err: errors.New(`ent: missing required field "ShowTime.room_id"`)}
	}
	if _, ok := stc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ShowTime.created_at"`)}
	}
	if _, ok := stc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ShowTime.updated_at"`)}
	}
	if _, ok := stc.mutation.RoomID(); !ok {
		return &ValidationError{Name: "room", err: errors.New(`ent: missing required edge "ShowTime.room"`)}
	}
	if _, ok := stc.mutation.MovieID(); !ok {
		return &ValidationError{Name: "movie", err: errors.New(`ent: missing required edge "ShowTime.movie"`)}
	}
	return nil
}

func (stc *ShowTimeCreate) sqlSave(ctx context.Context) (*ShowTime, error) {
	if err := stc.check(); err != nil {
		return nil, err
	}
	_node, _spec := stc.createSpec()
	if err := sqlgraph.CreateNode(ctx, stc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	stc.mutation.id = &_node.ID
	stc.mutation.done = true
	return _node, nil
}

func (stc *ShowTimeCreate) createSpec() (*ShowTime, *sqlgraph.CreateSpec) {
	var (
		_node = &ShowTime{config: stc.config}
		_spec = sqlgraph.NewCreateSpec(showtime.Table, sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID))
	)
	if id, ok := stc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := stc.mutation.StartAt(); ok {
		_spec.SetField(showtime.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := stc.mutation.EndAt(); ok {
		_spec.SetField(showtime.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := stc.mutation.CreatedAt(); ok {
		_spec.SetField(showtime.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := stc.mutation.UpdatedAt(); ok {
		_spec.SetField(showtime.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := stc.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.RoomTable,
			Columns: []string{showtime.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.RoomID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := stc.mutation.MovieIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.MovieTable,
			Columns: []string{showtime.MovieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MovieID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := stc.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   showtime.TicketsTable,
			Columns: []string{showtime.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ShowTimeCreateBulk is the builder for creating many ShowTime entities in bulk.
type ShowTimeCreateBulk struct {
	config
	err      error
	builders []*ShowTimeCreate
}

// Save creates the ShowTime entities in the database.
func (stcb *ShowTimeCreateBulk) Save(ctx context.Context) ([]*ShowTime, error) {
	if stcb.err != nil {
		return nil, stcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(stcb.builders))
	nodes := make([]*ShowTime, len(stcb.builders))
	mutators := make([]Mutator, len(stcb.builders))
	for i := range stcb.builders {
		func(i int, root context.Context) {
			builder := stcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShowTimeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, stcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, stcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, stcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (stcb *ShowTimeCreateBulk) SaveX(ctx context.Context) []*ShowTime {
	v, err := stcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (stcb *ShowTimeCreateBulk) Exec(ctx context.Context) error {
	_, err := stcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stcb *ShowTimeCreateBulk) ExecX(ctx context.Context) {
	if err := stcb.Exec(ctx); err != nil {
		panic(err)
	}
}