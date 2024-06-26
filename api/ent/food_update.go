// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/food"
	"PopcornMovie/ent/foodorderline"
	"PopcornMovie/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FoodUpdate is the builder for updating Food entities.
type FoodUpdate struct {
	config
	hooks    []Hook
	mutation *FoodMutation
}

// Where appends a list predicates to the FoodUpdate builder.
func (fu *FoodUpdate) Where(ps ...predicate.Food) *FoodUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetName sets the "name" field.
func (fu *FoodUpdate) SetName(s string) *FoodUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fu *FoodUpdate) SetNillableName(s *string) *FoodUpdate {
	if s != nil {
		fu.SetName(*s)
	}
	return fu
}

// SetPrice sets the "price" field.
func (fu *FoodUpdate) SetPrice(f float64) *FoodUpdate {
	fu.mutation.ResetPrice()
	fu.mutation.SetPrice(f)
	return fu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (fu *FoodUpdate) SetNillablePrice(f *float64) *FoodUpdate {
	if f != nil {
		fu.SetPrice(*f)
	}
	return fu
}

// AddPrice adds f to the "price" field.
func (fu *FoodUpdate) AddPrice(f float64) *FoodUpdate {
	fu.mutation.AddPrice(f)
	return fu
}

// SetImage sets the "image" field.
func (fu *FoodUpdate) SetImage(s string) *FoodUpdate {
	fu.mutation.SetImage(s)
	return fu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (fu *FoodUpdate) SetNillableImage(s *string) *FoodUpdate {
	if s != nil {
		fu.SetImage(*s)
	}
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FoodUpdate) SetUpdatedAt(t time.Time) *FoodUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// AddFoodOrderLineIDs adds the "food_order_lines" edge to the FoodOrderLine entity by IDs.
func (fu *FoodUpdate) AddFoodOrderLineIDs(ids ...uuid.UUID) *FoodUpdate {
	fu.mutation.AddFoodOrderLineIDs(ids...)
	return fu
}

// AddFoodOrderLines adds the "food_order_lines" edges to the FoodOrderLine entity.
func (fu *FoodUpdate) AddFoodOrderLines(f ...*FoodOrderLine) *FoodUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddFoodOrderLineIDs(ids...)
}

// Mutation returns the FoodMutation object of the builder.
func (fu *FoodUpdate) Mutation() *FoodMutation {
	return fu.mutation
}

// ClearFoodOrderLines clears all "food_order_lines" edges to the FoodOrderLine entity.
func (fu *FoodUpdate) ClearFoodOrderLines() *FoodUpdate {
	fu.mutation.ClearFoodOrderLines()
	return fu
}

// RemoveFoodOrderLineIDs removes the "food_order_lines" edge to FoodOrderLine entities by IDs.
func (fu *FoodUpdate) RemoveFoodOrderLineIDs(ids ...uuid.UUID) *FoodUpdate {
	fu.mutation.RemoveFoodOrderLineIDs(ids...)
	return fu
}

// RemoveFoodOrderLines removes "food_order_lines" edges to FoodOrderLine entities.
func (fu *FoodUpdate) RemoveFoodOrderLines(f ...*FoodOrderLine) *FoodUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveFoodOrderLineIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FoodUpdate) Save(ctx context.Context) (int, error) {
	fu.defaults()
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FoodUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FoodUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FoodUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FoodUpdate) defaults() {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		v := food.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
}

func (fu *FoodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(food.Table, food.Columns, sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(food.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.Price(); ok {
		_spec.SetField(food.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := fu.mutation.AddedPrice(); ok {
		_spec.AddField(food.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := fu.mutation.Image(); ok {
		_spec.SetField(food.FieldImage, field.TypeString, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(food.FieldUpdatedAt, field.TypeTime, value)
	}
	if fu.mutation.FoodOrderLinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   food.FoodOrderLinesTable,
			Columns: []string{food.FoodOrderLinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedFoodOrderLinesIDs(); len(nodes) > 0 && !fu.mutation.FoodOrderLinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   food.FoodOrderLinesTable,
			Columns: []string{food.FoodOrderLinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FoodOrderLinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   food.FoodOrderLinesTable,
			Columns: []string{food.FoodOrderLinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{food.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FoodUpdateOne is the builder for updating a single Food entity.
type FoodUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FoodMutation
}

// SetName sets the "name" field.
func (fuo *FoodUpdateOne) SetName(s string) *FoodUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fuo *FoodUpdateOne) SetNillableName(s *string) *FoodUpdateOne {
	if s != nil {
		fuo.SetName(*s)
	}
	return fuo
}

// SetPrice sets the "price" field.
func (fuo *FoodUpdateOne) SetPrice(f float64) *FoodUpdateOne {
	fuo.mutation.ResetPrice()
	fuo.mutation.SetPrice(f)
	return fuo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (fuo *FoodUpdateOne) SetNillablePrice(f *float64) *FoodUpdateOne {
	if f != nil {
		fuo.SetPrice(*f)
	}
	return fuo
}

// AddPrice adds f to the "price" field.
func (fuo *FoodUpdateOne) AddPrice(f float64) *FoodUpdateOne {
	fuo.mutation.AddPrice(f)
	return fuo
}

// SetImage sets the "image" field.
func (fuo *FoodUpdateOne) SetImage(s string) *FoodUpdateOne {
	fuo.mutation.SetImage(s)
	return fuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (fuo *FoodUpdateOne) SetNillableImage(s *string) *FoodUpdateOne {
	if s != nil {
		fuo.SetImage(*s)
	}
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FoodUpdateOne) SetUpdatedAt(t time.Time) *FoodUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// AddFoodOrderLineIDs adds the "food_order_lines" edge to the FoodOrderLine entity by IDs.
func (fuo *FoodUpdateOne) AddFoodOrderLineIDs(ids ...uuid.UUID) *FoodUpdateOne {
	fuo.mutation.AddFoodOrderLineIDs(ids...)
	return fuo
}

// AddFoodOrderLines adds the "food_order_lines" edges to the FoodOrderLine entity.
func (fuo *FoodUpdateOne) AddFoodOrderLines(f ...*FoodOrderLine) *FoodUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddFoodOrderLineIDs(ids...)
}

// Mutation returns the FoodMutation object of the builder.
func (fuo *FoodUpdateOne) Mutation() *FoodMutation {
	return fuo.mutation
}

// ClearFoodOrderLines clears all "food_order_lines" edges to the FoodOrderLine entity.
func (fuo *FoodUpdateOne) ClearFoodOrderLines() *FoodUpdateOne {
	fuo.mutation.ClearFoodOrderLines()
	return fuo
}

// RemoveFoodOrderLineIDs removes the "food_order_lines" edge to FoodOrderLine entities by IDs.
func (fuo *FoodUpdateOne) RemoveFoodOrderLineIDs(ids ...uuid.UUID) *FoodUpdateOne {
	fuo.mutation.RemoveFoodOrderLineIDs(ids...)
	return fuo
}

// RemoveFoodOrderLines removes "food_order_lines" edges to FoodOrderLine entities.
func (fuo *FoodUpdateOne) RemoveFoodOrderLines(f ...*FoodOrderLine) *FoodUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveFoodOrderLineIDs(ids...)
}

// Where appends a list predicates to the FoodUpdate builder.
func (fuo *FoodUpdateOne) Where(ps ...predicate.Food) *FoodUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FoodUpdateOne) Select(field string, fields ...string) *FoodUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Food entity.
func (fuo *FoodUpdateOne) Save(ctx context.Context) (*Food, error) {
	fuo.defaults()
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FoodUpdateOne) SaveX(ctx context.Context) *Food {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FoodUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FoodUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FoodUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		v := food.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
}

func (fuo *FoodUpdateOne) sqlSave(ctx context.Context) (_node *Food, err error) {
	_spec := sqlgraph.NewUpdateSpec(food.Table, food.Columns, sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Food.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, food.FieldID)
		for _, f := range fields {
			if !food.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != food.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(food.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Price(); ok {
		_spec.SetField(food.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := fuo.mutation.AddedPrice(); ok {
		_spec.AddField(food.FieldPrice, field.TypeFloat64, value)
	}
	if value, ok := fuo.mutation.Image(); ok {
		_spec.SetField(food.FieldImage, field.TypeString, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(food.FieldUpdatedAt, field.TypeTime, value)
	}
	if fuo.mutation.FoodOrderLinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   food.FoodOrderLinesTable,
			Columns: []string{food.FoodOrderLinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedFoodOrderLinesIDs(); len(nodes) > 0 && !fuo.mutation.FoodOrderLinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   food.FoodOrderLinesTable,
			Columns: []string{food.FoodOrderLinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FoodOrderLinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   food.FoodOrderLinesTable,
			Columns: []string{food.FoodOrderLinesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Food{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{food.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
