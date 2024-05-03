// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/food"
	"PopcornMovie/ent/foodorderline"
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/transaction"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FoodOrderLineUpdate is the builder for updating FoodOrderLine entities.
type FoodOrderLineUpdate struct {
	config
	hooks    []Hook
	mutation *FoodOrderLineMutation
}

// Where appends a list predicates to the FoodOrderLineUpdate builder.
func (folu *FoodOrderLineUpdate) Where(ps ...predicate.FoodOrderLine) *FoodOrderLineUpdate {
	folu.mutation.Where(ps...)
	return folu
}

// SetQuantity sets the "quantity" field.
func (folu *FoodOrderLineUpdate) SetQuantity(i int) *FoodOrderLineUpdate {
	folu.mutation.ResetQuantity()
	folu.mutation.SetQuantity(i)
	return folu
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (folu *FoodOrderLineUpdate) SetNillableQuantity(i *int) *FoodOrderLineUpdate {
	if i != nil {
		folu.SetQuantity(*i)
	}
	return folu
}

// AddQuantity adds i to the "quantity" field.
func (folu *FoodOrderLineUpdate) AddQuantity(i int) *FoodOrderLineUpdate {
	folu.mutation.AddQuantity(i)
	return folu
}

// SetFoodID sets the "food_id" field.
func (folu *FoodOrderLineUpdate) SetFoodID(u uuid.UUID) *FoodOrderLineUpdate {
	folu.mutation.SetFoodID(u)
	return folu
}

// SetNillableFoodID sets the "food_id" field if the given value is not nil.
func (folu *FoodOrderLineUpdate) SetNillableFoodID(u *uuid.UUID) *FoodOrderLineUpdate {
	if u != nil {
		folu.SetFoodID(*u)
	}
	return folu
}

// SetTransactionID sets the "transaction_id" field.
func (folu *FoodOrderLineUpdate) SetTransactionID(u uuid.UUID) *FoodOrderLineUpdate {
	folu.mutation.SetTransactionID(u)
	return folu
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (folu *FoodOrderLineUpdate) SetNillableTransactionID(u *uuid.UUID) *FoodOrderLineUpdate {
	if u != nil {
		folu.SetTransactionID(*u)
	}
	return folu
}

// SetUpdatedAt sets the "updated_at" field.
func (folu *FoodOrderLineUpdate) SetUpdatedAt(t time.Time) *FoodOrderLineUpdate {
	folu.mutation.SetUpdatedAt(t)
	return folu
}

// SetFood sets the "food" edge to the Food entity.
func (folu *FoodOrderLineUpdate) SetFood(f *Food) *FoodOrderLineUpdate {
	return folu.SetFoodID(f.ID)
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (folu *FoodOrderLineUpdate) SetTransaction(t *Transaction) *FoodOrderLineUpdate {
	return folu.SetTransactionID(t.ID)
}

// Mutation returns the FoodOrderLineMutation object of the builder.
func (folu *FoodOrderLineUpdate) Mutation() *FoodOrderLineMutation {
	return folu.mutation
}

// ClearFood clears the "food" edge to the Food entity.
func (folu *FoodOrderLineUpdate) ClearFood() *FoodOrderLineUpdate {
	folu.mutation.ClearFood()
	return folu
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (folu *FoodOrderLineUpdate) ClearTransaction() *FoodOrderLineUpdate {
	folu.mutation.ClearTransaction()
	return folu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (folu *FoodOrderLineUpdate) Save(ctx context.Context) (int, error) {
	folu.defaults()
	return withHooks(ctx, folu.sqlSave, folu.mutation, folu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (folu *FoodOrderLineUpdate) SaveX(ctx context.Context) int {
	affected, err := folu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (folu *FoodOrderLineUpdate) Exec(ctx context.Context) error {
	_, err := folu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (folu *FoodOrderLineUpdate) ExecX(ctx context.Context) {
	if err := folu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (folu *FoodOrderLineUpdate) defaults() {
	if _, ok := folu.mutation.UpdatedAt(); !ok {
		v := foodorderline.UpdateDefaultUpdatedAt()
		folu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (folu *FoodOrderLineUpdate) check() error {
	if _, ok := folu.mutation.FoodID(); folu.mutation.FoodCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FoodOrderLine.food"`)
	}
	if _, ok := folu.mutation.TransactionID(); folu.mutation.TransactionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FoodOrderLine.transaction"`)
	}
	return nil
}

func (folu *FoodOrderLineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := folu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(foodorderline.Table, foodorderline.Columns, sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID))
	if ps := folu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := folu.mutation.Quantity(); ok {
		_spec.SetField(foodorderline.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := folu.mutation.AddedQuantity(); ok {
		_spec.AddField(foodorderline.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := folu.mutation.UpdatedAt(); ok {
		_spec.SetField(foodorderline.FieldUpdatedAt, field.TypeTime, value)
	}
	if folu.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.FoodTable,
			Columns: []string{foodorderline.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := folu.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.FoodTable,
			Columns: []string{foodorderline.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if folu.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.TransactionTable,
			Columns: []string{foodorderline.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := folu.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.TransactionTable,
			Columns: []string{foodorderline.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, folu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{foodorderline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	folu.mutation.done = true
	return n, nil
}

// FoodOrderLineUpdateOne is the builder for updating a single FoodOrderLine entity.
type FoodOrderLineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FoodOrderLineMutation
}

// SetQuantity sets the "quantity" field.
func (foluo *FoodOrderLineUpdateOne) SetQuantity(i int) *FoodOrderLineUpdateOne {
	foluo.mutation.ResetQuantity()
	foluo.mutation.SetQuantity(i)
	return foluo
}

// SetNillableQuantity sets the "quantity" field if the given value is not nil.
func (foluo *FoodOrderLineUpdateOne) SetNillableQuantity(i *int) *FoodOrderLineUpdateOne {
	if i != nil {
		foluo.SetQuantity(*i)
	}
	return foluo
}

// AddQuantity adds i to the "quantity" field.
func (foluo *FoodOrderLineUpdateOne) AddQuantity(i int) *FoodOrderLineUpdateOne {
	foluo.mutation.AddQuantity(i)
	return foluo
}

// SetFoodID sets the "food_id" field.
func (foluo *FoodOrderLineUpdateOne) SetFoodID(u uuid.UUID) *FoodOrderLineUpdateOne {
	foluo.mutation.SetFoodID(u)
	return foluo
}

// SetNillableFoodID sets the "food_id" field if the given value is not nil.
func (foluo *FoodOrderLineUpdateOne) SetNillableFoodID(u *uuid.UUID) *FoodOrderLineUpdateOne {
	if u != nil {
		foluo.SetFoodID(*u)
	}
	return foluo
}

// SetTransactionID sets the "transaction_id" field.
func (foluo *FoodOrderLineUpdateOne) SetTransactionID(u uuid.UUID) *FoodOrderLineUpdateOne {
	foluo.mutation.SetTransactionID(u)
	return foluo
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (foluo *FoodOrderLineUpdateOne) SetNillableTransactionID(u *uuid.UUID) *FoodOrderLineUpdateOne {
	if u != nil {
		foluo.SetTransactionID(*u)
	}
	return foluo
}

// SetUpdatedAt sets the "updated_at" field.
func (foluo *FoodOrderLineUpdateOne) SetUpdatedAt(t time.Time) *FoodOrderLineUpdateOne {
	foluo.mutation.SetUpdatedAt(t)
	return foluo
}

// SetFood sets the "food" edge to the Food entity.
func (foluo *FoodOrderLineUpdateOne) SetFood(f *Food) *FoodOrderLineUpdateOne {
	return foluo.SetFoodID(f.ID)
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (foluo *FoodOrderLineUpdateOne) SetTransaction(t *Transaction) *FoodOrderLineUpdateOne {
	return foluo.SetTransactionID(t.ID)
}

// Mutation returns the FoodOrderLineMutation object of the builder.
func (foluo *FoodOrderLineUpdateOne) Mutation() *FoodOrderLineMutation {
	return foluo.mutation
}

// ClearFood clears the "food" edge to the Food entity.
func (foluo *FoodOrderLineUpdateOne) ClearFood() *FoodOrderLineUpdateOne {
	foluo.mutation.ClearFood()
	return foluo
}

// ClearTransaction clears the "transaction" edge to the Transaction entity.
func (foluo *FoodOrderLineUpdateOne) ClearTransaction() *FoodOrderLineUpdateOne {
	foluo.mutation.ClearTransaction()
	return foluo
}

// Where appends a list predicates to the FoodOrderLineUpdate builder.
func (foluo *FoodOrderLineUpdateOne) Where(ps ...predicate.FoodOrderLine) *FoodOrderLineUpdateOne {
	foluo.mutation.Where(ps...)
	return foluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (foluo *FoodOrderLineUpdateOne) Select(field string, fields ...string) *FoodOrderLineUpdateOne {
	foluo.fields = append([]string{field}, fields...)
	return foluo
}

// Save executes the query and returns the updated FoodOrderLine entity.
func (foluo *FoodOrderLineUpdateOne) Save(ctx context.Context) (*FoodOrderLine, error) {
	foluo.defaults()
	return withHooks(ctx, foluo.sqlSave, foluo.mutation, foluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (foluo *FoodOrderLineUpdateOne) SaveX(ctx context.Context) *FoodOrderLine {
	node, err := foluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (foluo *FoodOrderLineUpdateOne) Exec(ctx context.Context) error {
	_, err := foluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (foluo *FoodOrderLineUpdateOne) ExecX(ctx context.Context) {
	if err := foluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (foluo *FoodOrderLineUpdateOne) defaults() {
	if _, ok := foluo.mutation.UpdatedAt(); !ok {
		v := foodorderline.UpdateDefaultUpdatedAt()
		foluo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (foluo *FoodOrderLineUpdateOne) check() error {
	if _, ok := foluo.mutation.FoodID(); foluo.mutation.FoodCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FoodOrderLine.food"`)
	}
	if _, ok := foluo.mutation.TransactionID(); foluo.mutation.TransactionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FoodOrderLine.transaction"`)
	}
	return nil
}

func (foluo *FoodOrderLineUpdateOne) sqlSave(ctx context.Context) (_node *FoodOrderLine, err error) {
	if err := foluo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(foodorderline.Table, foodorderline.Columns, sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID))
	id, ok := foluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FoodOrderLine.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := foluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, foodorderline.FieldID)
		for _, f := range fields {
			if !foodorderline.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != foodorderline.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := foluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := foluo.mutation.Quantity(); ok {
		_spec.SetField(foodorderline.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := foluo.mutation.AddedQuantity(); ok {
		_spec.AddField(foodorderline.FieldQuantity, field.TypeInt, value)
	}
	if value, ok := foluo.mutation.UpdatedAt(); ok {
		_spec.SetField(foodorderline.FieldUpdatedAt, field.TypeTime, value)
	}
	if foluo.mutation.FoodCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.FoodTable,
			Columns: []string{foodorderline.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := foluo.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.FoodTable,
			Columns: []string{foodorderline.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if foluo.mutation.TransactionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.TransactionTable,
			Columns: []string{foodorderline.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := foluo.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.TransactionTable,
			Columns: []string{foodorderline.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FoodOrderLine{config: foluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, foluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{foodorderline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	foluo.mutation.done = true
	return _node, nil
}