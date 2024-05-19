// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/foodorderline"
	"PopcornMovie/ent/ticket"
	"PopcornMovie/ent/transaction"
	"PopcornMovie/ent/user"
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

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTotal sets the "total" field.
func (tc *TransactionCreate) SetTotal(f float64) *TransactionCreate {
	tc.mutation.SetTotal(f)
	return tc
}

// SetUserID sets the "user_id" field.
func (tc *TransactionCreate) SetUserID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetUserID(u)
	return tc
}

// SetCode sets the "code" field.
func (tc *TransactionCreate) SetCode(i int) *TransactionCreate {
	tc.mutation.SetCode(i)
	return tc
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCode(i *int) *TransactionCreate {
	if i != nil {
		tc.SetCode(*i)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TransactionCreate) SetStatus(t transaction.Status) *TransactionCreate {
	tc.mutation.SetStatus(t)
	return tc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableStatus(t *transaction.Status) *TransactionCreate {
	if t != nil {
		tc.SetStatus(*t)
	}
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TransactionCreate) SetCreatedAt(t time.Time) *TransactionCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableCreatedAt(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TransactionCreate) SetUpdatedAt(t time.Time) *TransactionCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableUpdatedAt(t *time.Time) *TransactionCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(u uuid.UUID) *TransactionCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetUser sets the "user" edge to the User entity.
func (tc *TransactionCreate) SetUser(u *User) *TransactionCreate {
	return tc.SetUserID(u.ID)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (tc *TransactionCreate) AddTicketIDs(ids ...uuid.UUID) *TransactionCreate {
	tc.mutation.AddTicketIDs(ids...)
	return tc
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (tc *TransactionCreate) AddTickets(t ...*Ticket) *TransactionCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tc.AddTicketIDs(ids...)
}

// AddFoodOrderLineIDs adds the "food_order_lines" edge to the FoodOrderLine entity by IDs.
func (tc *TransactionCreate) AddFoodOrderLineIDs(ids ...uuid.UUID) *TransactionCreate {
	tc.mutation.AddFoodOrderLineIDs(ids...)
	return tc
}

// AddFoodOrderLines adds the "food_order_lines" edges to the FoodOrderLine entity.
func (tc *TransactionCreate) AddFoodOrderLines(f ...*FoodOrderLine) *TransactionCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return tc.AddFoodOrderLineIDs(ids...)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TransactionCreate) defaults() {
	if _, ok := tc.mutation.Status(); !ok {
		v := transaction.DefaultStatus
		tc.mutation.SetStatus(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := transaction.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := transaction.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.Total(); !ok {
		return &ValidationError{Name: "total", err: errors.New(`ent: missing required field "Transaction.total"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Transaction.user_id"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Transaction.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := transaction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Transaction.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Transaction.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Transaction.updated_at"`)}
	}
	if _, ok := tc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Transaction.user"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(transaction.Table, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.Total(); ok {
		_spec.SetField(transaction.FieldTotal, field.TypeFloat64, value)
		_node.Total = value
	}
	if value, ok := tc.mutation.Code(); ok {
		_spec.SetField(transaction.FieldCode, field.TypeInt, value)
		_node.Code = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(transaction.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(transaction.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(transaction.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.UserTable,
			Columns: []string{transaction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transaction.TicketsTable,
			Columns: []string{transaction.TicketsColumn},
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
	if nodes := tc.mutation.FoodOrderLinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transaction.FoodOrderLinesTable,
			Columns: []string{transaction.FoodOrderLinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID),
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
//	client.Transaction.Create().
//		SetTotal(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TransactionUpsert) {
//			SetTotal(v+v).
//		}).
//		Exec(ctx)
func (tc *TransactionCreate) OnConflict(opts ...sql.ConflictOption) *TransactionUpsertOne {
	tc.conflict = opts
	return &TransactionUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TransactionCreate) OnConflictColumns(columns ...string) *TransactionUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TransactionUpsertOne{
		create: tc,
	}
}

type (
	// TransactionUpsertOne is the builder for "upsert"-ing
	//  one Transaction node.
	TransactionUpsertOne struct {
		create *TransactionCreate
	}

	// TransactionUpsert is the "OnConflict" setter.
	TransactionUpsert struct {
		*sql.UpdateSet
	}
)

// SetTotal sets the "total" field.
func (u *TransactionUpsert) SetTotal(v float64) *TransactionUpsert {
	u.Set(transaction.FieldTotal, v)
	return u
}

// UpdateTotal sets the "total" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateTotal() *TransactionUpsert {
	u.SetExcluded(transaction.FieldTotal)
	return u
}

// AddTotal adds v to the "total" field.
func (u *TransactionUpsert) AddTotal(v float64) *TransactionUpsert {
	u.Add(transaction.FieldTotal, v)
	return u
}

// SetUserID sets the "user_id" field.
func (u *TransactionUpsert) SetUserID(v uuid.UUID) *TransactionUpsert {
	u.Set(transaction.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateUserID() *TransactionUpsert {
	u.SetExcluded(transaction.FieldUserID)
	return u
}

// SetCode sets the "code" field.
func (u *TransactionUpsert) SetCode(v int) *TransactionUpsert {
	u.Set(transaction.FieldCode, v)
	return u
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateCode() *TransactionUpsert {
	u.SetExcluded(transaction.FieldCode)
	return u
}

// AddCode adds v to the "code" field.
func (u *TransactionUpsert) AddCode(v int) *TransactionUpsert {
	u.Add(transaction.FieldCode, v)
	return u
}

// ClearCode clears the value of the "code" field.
func (u *TransactionUpsert) ClearCode() *TransactionUpsert {
	u.SetNull(transaction.FieldCode)
	return u
}

// SetStatus sets the "status" field.
func (u *TransactionUpsert) SetStatus(v transaction.Status) *TransactionUpsert {
	u.Set(transaction.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateStatus() *TransactionUpsert {
	u.SetExcluded(transaction.FieldStatus)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TransactionUpsert) SetUpdatedAt(v time.Time) *TransactionUpsert {
	u.Set(transaction.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateUpdatedAt() *TransactionUpsert {
	u.SetExcluded(transaction.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transaction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TransactionUpsertOne) UpdateNewValues() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(transaction.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(transaction.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Transaction.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TransactionUpsertOne) Ignore() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TransactionUpsertOne) DoNothing() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TransactionCreate.OnConflict
// documentation for more info.
func (u *TransactionUpsertOne) Update(set func(*TransactionUpsert)) *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TransactionUpsert{UpdateSet: update})
	}))
	return u
}

// SetTotal sets the "total" field.
func (u *TransactionUpsertOne) SetTotal(v float64) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTotal(v)
	})
}

// AddTotal adds v to the "total" field.
func (u *TransactionUpsertOne) AddTotal(v float64) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.AddTotal(v)
	})
}

// UpdateTotal sets the "total" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateTotal() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTotal()
	})
}

// SetUserID sets the "user_id" field.
func (u *TransactionUpsertOne) SetUserID(v uuid.UUID) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateUserID() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateUserID()
	})
}

// SetCode sets the "code" field.
func (u *TransactionUpsertOne) SetCode(v int) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCode(v)
	})
}

// AddCode adds v to the "code" field.
func (u *TransactionUpsertOne) AddCode(v int) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.AddCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateCode() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCode()
	})
}

// ClearCode clears the value of the "code" field.
func (u *TransactionUpsertOne) ClearCode() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearCode()
	})
}

// SetStatus sets the "status" field.
func (u *TransactionUpsertOne) SetStatus(v transaction.Status) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateStatus() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TransactionUpsertOne) SetUpdatedAt(v time.Time) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateUpdatedAt() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TransactionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TransactionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TransactionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TransactionUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TransactionUpsertOne.ID is not supported by MySQL driver. Use TransactionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TransactionUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	err      error
	builders []*TransactionCreate
	conflict []sql.ConflictOption
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Transaction.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TransactionUpsert) {
//			SetTotal(v+v).
//		}).
//		Exec(ctx)
func (tcb *TransactionCreateBulk) OnConflict(opts ...sql.ConflictOption) *TransactionUpsertBulk {
	tcb.conflict = opts
	return &TransactionUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TransactionCreateBulk) OnConflictColumns(columns ...string) *TransactionUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TransactionUpsertBulk{
		create: tcb,
	}
}

// TransactionUpsertBulk is the builder for "upsert"-ing
// a bulk of Transaction nodes.
type TransactionUpsertBulk struct {
	create *TransactionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transaction.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TransactionUpsertBulk) UpdateNewValues() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(transaction.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(transaction.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TransactionUpsertBulk) Ignore() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TransactionUpsertBulk) DoNothing() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TransactionCreateBulk.OnConflict
// documentation for more info.
func (u *TransactionUpsertBulk) Update(set func(*TransactionUpsert)) *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TransactionUpsert{UpdateSet: update})
	}))
	return u
}

// SetTotal sets the "total" field.
func (u *TransactionUpsertBulk) SetTotal(v float64) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTotal(v)
	})
}

// AddTotal adds v to the "total" field.
func (u *TransactionUpsertBulk) AddTotal(v float64) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.AddTotal(v)
	})
}

// UpdateTotal sets the "total" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateTotal() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTotal()
	})
}

// SetUserID sets the "user_id" field.
func (u *TransactionUpsertBulk) SetUserID(v uuid.UUID) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateUserID() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateUserID()
	})
}

// SetCode sets the "code" field.
func (u *TransactionUpsertBulk) SetCode(v int) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCode(v)
	})
}

// AddCode adds v to the "code" field.
func (u *TransactionUpsertBulk) AddCode(v int) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.AddCode(v)
	})
}

// UpdateCode sets the "code" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateCode() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCode()
	})
}

// ClearCode clears the value of the "code" field.
func (u *TransactionUpsertBulk) ClearCode() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearCode()
	})
}

// SetStatus sets the "status" field.
func (u *TransactionUpsertBulk) SetStatus(v transaction.Status) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateStatus() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateStatus()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TransactionUpsertBulk) SetUpdatedAt(v time.Time) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateUpdatedAt() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TransactionUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TransactionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TransactionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TransactionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
