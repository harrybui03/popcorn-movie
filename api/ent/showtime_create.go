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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ShowTimeCreate is the builder for creating a ShowTime entity.
type ShowTimeCreate struct {
	config
	mutation *ShowTimeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
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
	_spec.OnConflict = stc.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ShowTime.Create().
//		SetStartAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShowTimeUpsert) {
//			SetStartAt(v+v).
//		}).
//		Exec(ctx)
func (stc *ShowTimeCreate) OnConflict(opts ...sql.ConflictOption) *ShowTimeUpsertOne {
	stc.conflict = opts
	return &ShowTimeUpsertOne{
		create: stc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ShowTime.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (stc *ShowTimeCreate) OnConflictColumns(columns ...string) *ShowTimeUpsertOne {
	stc.conflict = append(stc.conflict, sql.ConflictColumns(columns...))
	return &ShowTimeUpsertOne{
		create: stc,
	}
}

type (
	// ShowTimeUpsertOne is the builder for "upsert"-ing
	//  one ShowTime node.
	ShowTimeUpsertOne struct {
		create *ShowTimeCreate
	}

	// ShowTimeUpsert is the "OnConflict" setter.
	ShowTimeUpsert struct {
		*sql.UpdateSet
	}
)

// SetStartAt sets the "start_at" field.
func (u *ShowTimeUpsert) SetStartAt(v time.Time) *ShowTimeUpsert {
	u.Set(showtime.FieldStartAt, v)
	return u
}

// UpdateStartAt sets the "start_at" field to the value that was provided on create.
func (u *ShowTimeUpsert) UpdateStartAt() *ShowTimeUpsert {
	u.SetExcluded(showtime.FieldStartAt)
	return u
}

// SetEndAt sets the "end_at" field.
func (u *ShowTimeUpsert) SetEndAt(v time.Time) *ShowTimeUpsert {
	u.Set(showtime.FieldEndAt, v)
	return u
}

// UpdateEndAt sets the "end_at" field to the value that was provided on create.
func (u *ShowTimeUpsert) UpdateEndAt() *ShowTimeUpsert {
	u.SetExcluded(showtime.FieldEndAt)
	return u
}

// SetMovieID sets the "movie_id" field.
func (u *ShowTimeUpsert) SetMovieID(v uuid.UUID) *ShowTimeUpsert {
	u.Set(showtime.FieldMovieID, v)
	return u
}

// UpdateMovieID sets the "movie_id" field to the value that was provided on create.
func (u *ShowTimeUpsert) UpdateMovieID() *ShowTimeUpsert {
	u.SetExcluded(showtime.FieldMovieID)
	return u
}

// SetRoomID sets the "room_id" field.
func (u *ShowTimeUpsert) SetRoomID(v uuid.UUID) *ShowTimeUpsert {
	u.Set(showtime.FieldRoomID, v)
	return u
}

// UpdateRoomID sets the "room_id" field to the value that was provided on create.
func (u *ShowTimeUpsert) UpdateRoomID() *ShowTimeUpsert {
	u.SetExcluded(showtime.FieldRoomID)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ShowTimeUpsert) SetUpdatedAt(v time.Time) *ShowTimeUpsert {
	u.Set(showtime.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ShowTimeUpsert) UpdateUpdatedAt() *ShowTimeUpsert {
	u.SetExcluded(showtime.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ShowTime.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(showtime.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ShowTimeUpsertOne) UpdateNewValues() *ShowTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(showtime.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(showtime.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ShowTime.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ShowTimeUpsertOne) Ignore() *ShowTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShowTimeUpsertOne) DoNothing() *ShowTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShowTimeCreate.OnConflict
// documentation for more info.
func (u *ShowTimeUpsertOne) Update(set func(*ShowTimeUpsert)) *ShowTimeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShowTimeUpsert{UpdateSet: update})
	}))
	return u
}

// SetStartAt sets the "start_at" field.
func (u *ShowTimeUpsertOne) SetStartAt(v time.Time) *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetStartAt(v)
	})
}

// UpdateStartAt sets the "start_at" field to the value that was provided on create.
func (u *ShowTimeUpsertOne) UpdateStartAt() *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateStartAt()
	})
}

// SetEndAt sets the "end_at" field.
func (u *ShowTimeUpsertOne) SetEndAt(v time.Time) *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetEndAt(v)
	})
}

// UpdateEndAt sets the "end_at" field to the value that was provided on create.
func (u *ShowTimeUpsertOne) UpdateEndAt() *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateEndAt()
	})
}

// SetMovieID sets the "movie_id" field.
func (u *ShowTimeUpsertOne) SetMovieID(v uuid.UUID) *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetMovieID(v)
	})
}

// UpdateMovieID sets the "movie_id" field to the value that was provided on create.
func (u *ShowTimeUpsertOne) UpdateMovieID() *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateMovieID()
	})
}

// SetRoomID sets the "room_id" field.
func (u *ShowTimeUpsertOne) SetRoomID(v uuid.UUID) *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetRoomID(v)
	})
}

// UpdateRoomID sets the "room_id" field to the value that was provided on create.
func (u *ShowTimeUpsertOne) UpdateRoomID() *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateRoomID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ShowTimeUpsertOne) SetUpdatedAt(v time.Time) *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ShowTimeUpsertOne) UpdateUpdatedAt() *ShowTimeUpsertOne {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ShowTimeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShowTimeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShowTimeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ShowTimeUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ShowTimeUpsertOne.ID is not supported by MySQL driver. Use ShowTimeUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ShowTimeUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ShowTimeCreateBulk is the builder for creating many ShowTime entities in bulk.
type ShowTimeCreateBulk struct {
	config
	err      error
	builders []*ShowTimeCreate
	conflict []sql.ConflictOption
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
					spec.OnConflict = stcb.conflict
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ShowTime.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShowTimeUpsert) {
//			SetStartAt(v+v).
//		}).
//		Exec(ctx)
func (stcb *ShowTimeCreateBulk) OnConflict(opts ...sql.ConflictOption) *ShowTimeUpsertBulk {
	stcb.conflict = opts
	return &ShowTimeUpsertBulk{
		create: stcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ShowTime.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (stcb *ShowTimeCreateBulk) OnConflictColumns(columns ...string) *ShowTimeUpsertBulk {
	stcb.conflict = append(stcb.conflict, sql.ConflictColumns(columns...))
	return &ShowTimeUpsertBulk{
		create: stcb,
	}
}

// ShowTimeUpsertBulk is the builder for "upsert"-ing
// a bulk of ShowTime nodes.
type ShowTimeUpsertBulk struct {
	create *ShowTimeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ShowTime.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(showtime.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ShowTimeUpsertBulk) UpdateNewValues() *ShowTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(showtime.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(showtime.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ShowTime.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ShowTimeUpsertBulk) Ignore() *ShowTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShowTimeUpsertBulk) DoNothing() *ShowTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShowTimeCreateBulk.OnConflict
// documentation for more info.
func (u *ShowTimeUpsertBulk) Update(set func(*ShowTimeUpsert)) *ShowTimeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShowTimeUpsert{UpdateSet: update})
	}))
	return u
}

// SetStartAt sets the "start_at" field.
func (u *ShowTimeUpsertBulk) SetStartAt(v time.Time) *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetStartAt(v)
	})
}

// UpdateStartAt sets the "start_at" field to the value that was provided on create.
func (u *ShowTimeUpsertBulk) UpdateStartAt() *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateStartAt()
	})
}

// SetEndAt sets the "end_at" field.
func (u *ShowTimeUpsertBulk) SetEndAt(v time.Time) *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetEndAt(v)
	})
}

// UpdateEndAt sets the "end_at" field to the value that was provided on create.
func (u *ShowTimeUpsertBulk) UpdateEndAt() *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateEndAt()
	})
}

// SetMovieID sets the "movie_id" field.
func (u *ShowTimeUpsertBulk) SetMovieID(v uuid.UUID) *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetMovieID(v)
	})
}

// UpdateMovieID sets the "movie_id" field to the value that was provided on create.
func (u *ShowTimeUpsertBulk) UpdateMovieID() *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateMovieID()
	})
}

// SetRoomID sets the "room_id" field.
func (u *ShowTimeUpsertBulk) SetRoomID(v uuid.UUID) *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetRoomID(v)
	})
}

// UpdateRoomID sets the "room_id" field to the value that was provided on create.
func (u *ShowTimeUpsertBulk) UpdateRoomID() *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateRoomID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ShowTimeUpsertBulk) SetUpdatedAt(v time.Time) *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ShowTimeUpsertBulk) UpdateUpdatedAt() *ShowTimeUpsertBulk {
	return u.Update(func(s *ShowTimeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ShowTimeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ShowTimeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShowTimeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShowTimeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
