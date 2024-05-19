// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/resetpassword"
	"PopcornMovie/ent/user"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ResetPasswordQuery is the builder for querying ResetPassword entities.
type ResetPasswordQuery struct {
	config
	ctx        *QueryContext
	order      []resetpassword.OrderOption
	inters     []Interceptor
	predicates []predicate.ResetPassword
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ResetPasswordQuery builder.
func (rpq *ResetPasswordQuery) Where(ps ...predicate.ResetPassword) *ResetPasswordQuery {
	rpq.predicates = append(rpq.predicates, ps...)
	return rpq
}

// Limit the number of records to be returned by this query.
func (rpq *ResetPasswordQuery) Limit(limit int) *ResetPasswordQuery {
	rpq.ctx.Limit = &limit
	return rpq
}

// Offset to start from.
func (rpq *ResetPasswordQuery) Offset(offset int) *ResetPasswordQuery {
	rpq.ctx.Offset = &offset
	return rpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rpq *ResetPasswordQuery) Unique(unique bool) *ResetPasswordQuery {
	rpq.ctx.Unique = &unique
	return rpq
}

// Order specifies how the records should be ordered.
func (rpq *ResetPasswordQuery) Order(o ...resetpassword.OrderOption) *ResetPasswordQuery {
	rpq.order = append(rpq.order, o...)
	return rpq
}

// QueryUser chains the current query on the "user" edge.
func (rpq *ResetPasswordQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: rpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(resetpassword.Table, resetpassword.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, resetpassword.UserTable, resetpassword.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(rpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ResetPassword entity from the query.
// Returns a *NotFoundError when no ResetPassword was found.
func (rpq *ResetPasswordQuery) First(ctx context.Context) (*ResetPassword, error) {
	nodes, err := rpq.Limit(1).All(setContextOp(ctx, rpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{resetpassword.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rpq *ResetPasswordQuery) FirstX(ctx context.Context) *ResetPassword {
	node, err := rpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ResetPassword ID from the query.
// Returns a *NotFoundError when no ResetPassword ID was found.
func (rpq *ResetPasswordQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rpq.Limit(1).IDs(setContextOp(ctx, rpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{resetpassword.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rpq *ResetPasswordQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ResetPassword entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ResetPassword entity is found.
// Returns a *NotFoundError when no ResetPassword entities are found.
func (rpq *ResetPasswordQuery) Only(ctx context.Context) (*ResetPassword, error) {
	nodes, err := rpq.Limit(2).All(setContextOp(ctx, rpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{resetpassword.Label}
	default:
		return nil, &NotSingularError{resetpassword.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rpq *ResetPasswordQuery) OnlyX(ctx context.Context) *ResetPassword {
	node, err := rpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ResetPassword ID in the query.
// Returns a *NotSingularError when more than one ResetPassword ID is found.
// Returns a *NotFoundError when no entities are found.
func (rpq *ResetPasswordQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rpq.Limit(2).IDs(setContextOp(ctx, rpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{resetpassword.Label}
	default:
		err = &NotSingularError{resetpassword.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rpq *ResetPasswordQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ResetPasswords.
func (rpq *ResetPasswordQuery) All(ctx context.Context) ([]*ResetPassword, error) {
	ctx = setContextOp(ctx, rpq.ctx, "All")
	if err := rpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ResetPassword, *ResetPasswordQuery]()
	return withInterceptors[[]*ResetPassword](ctx, rpq, qr, rpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rpq *ResetPasswordQuery) AllX(ctx context.Context) []*ResetPassword {
	nodes, err := rpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ResetPassword IDs.
func (rpq *ResetPasswordQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if rpq.ctx.Unique == nil && rpq.path != nil {
		rpq.Unique(true)
	}
	ctx = setContextOp(ctx, rpq.ctx, "IDs")
	if err = rpq.Select(resetpassword.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rpq *ResetPasswordQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rpq *ResetPasswordQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Count")
	if err := rpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rpq, querierCount[*ResetPasswordQuery](), rpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rpq *ResetPasswordQuery) CountX(ctx context.Context) int {
	count, err := rpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rpq *ResetPasswordQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rpq.ctx, "Exist")
	switch _, err := rpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rpq *ResetPasswordQuery) ExistX(ctx context.Context) bool {
	exist, err := rpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ResetPasswordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rpq *ResetPasswordQuery) Clone() *ResetPasswordQuery {
	if rpq == nil {
		return nil
	}
	return &ResetPasswordQuery{
		config:     rpq.config,
		ctx:        rpq.ctx.Clone(),
		order:      append([]resetpassword.OrderOption{}, rpq.order...),
		inters:     append([]Interceptor{}, rpq.inters...),
		predicates: append([]predicate.ResetPassword{}, rpq.predicates...),
		withUser:   rpq.withUser.Clone(),
		// clone intermediate query.
		sql:  rpq.sql.Clone(),
		path: rpq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (rpq *ResetPasswordQuery) WithUser(opts ...func(*UserQuery)) *ResetPasswordQuery {
	query := (&UserClient{config: rpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rpq.withUser = query
	return rpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ResetPassword.Query().
//		GroupBy(resetpassword.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rpq *ResetPasswordQuery) GroupBy(field string, fields ...string) *ResetPasswordGroupBy {
	rpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ResetPasswordGroupBy{build: rpq}
	grbuild.flds = &rpq.ctx.Fields
	grbuild.label = resetpassword.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID uuid.UUID `json:"user_id,omitempty"`
//	}
//
//	client.ResetPassword.Query().
//		Select(resetpassword.FieldUserID).
//		Scan(ctx, &v)
func (rpq *ResetPasswordQuery) Select(fields ...string) *ResetPasswordSelect {
	rpq.ctx.Fields = append(rpq.ctx.Fields, fields...)
	sbuild := &ResetPasswordSelect{ResetPasswordQuery: rpq}
	sbuild.label = resetpassword.Label
	sbuild.flds, sbuild.scan = &rpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ResetPasswordSelect configured with the given aggregations.
func (rpq *ResetPasswordQuery) Aggregate(fns ...AggregateFunc) *ResetPasswordSelect {
	return rpq.Select().Aggregate(fns...)
}

func (rpq *ResetPasswordQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rpq); err != nil {
				return err
			}
		}
	}
	for _, f := range rpq.ctx.Fields {
		if !resetpassword.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rpq.path != nil {
		prev, err := rpq.path(ctx)
		if err != nil {
			return err
		}
		rpq.sql = prev
	}
	return nil
}

func (rpq *ResetPasswordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ResetPassword, error) {
	var (
		nodes       = []*ResetPassword{}
		_spec       = rpq.querySpec()
		loadedTypes = [1]bool{
			rpq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ResetPassword).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ResetPassword{config: rpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rpq.modifiers) > 0 {
		_spec.Modifiers = rpq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rpq.withUser; query != nil {
		if err := rpq.loadUser(ctx, query, nodes, nil,
			func(n *ResetPassword, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rpq *ResetPasswordQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ResetPassword, init func(*ResetPassword), assign func(*ResetPassword, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ResetPassword)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rpq *ResetPasswordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rpq.querySpec()
	if len(rpq.modifiers) > 0 {
		_spec.Modifiers = rpq.modifiers
	}
	_spec.Node.Columns = rpq.ctx.Fields
	if len(rpq.ctx.Fields) > 0 {
		_spec.Unique = rpq.ctx.Unique != nil && *rpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rpq.driver, _spec)
}

func (rpq *ResetPasswordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(resetpassword.Table, resetpassword.Columns, sqlgraph.NewFieldSpec(resetpassword.FieldID, field.TypeUUID))
	_spec.From = rpq.sql
	if unique := rpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rpq.path != nil {
		_spec.Unique = true
	}
	if fields := rpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, resetpassword.FieldID)
		for i := range fields {
			if fields[i] != resetpassword.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rpq.withUser != nil {
			_spec.Node.AddColumnOnce(resetpassword.FieldUserID)
		}
	}
	if ps := rpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rpq *ResetPasswordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rpq.driver.Dialect())
	t1 := builder.Table(resetpassword.Table)
	columns := rpq.ctx.Fields
	if len(columns) == 0 {
		columns = resetpassword.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rpq.sql != nil {
		selector = rpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rpq.ctx.Unique != nil && *rpq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range rpq.modifiers {
		m(selector)
	}
	for _, p := range rpq.predicates {
		p(selector)
	}
	for _, p := range rpq.order {
		p(selector)
	}
	if offset := rpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rpq *ResetPasswordQuery) ForUpdate(opts ...sql.LockOption) *ResetPasswordQuery {
	if rpq.driver.Dialect() == dialect.Postgres {
		rpq.Unique(false)
	}
	rpq.modifiers = append(rpq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rpq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rpq *ResetPasswordQuery) ForShare(opts ...sql.LockOption) *ResetPasswordQuery {
	if rpq.driver.Dialect() == dialect.Postgres {
		rpq.Unique(false)
	}
	rpq.modifiers = append(rpq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rpq
}

// ResetPasswordGroupBy is the group-by builder for ResetPassword entities.
type ResetPasswordGroupBy struct {
	selector
	build *ResetPasswordQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rpgb *ResetPasswordGroupBy) Aggregate(fns ...AggregateFunc) *ResetPasswordGroupBy {
	rpgb.fns = append(rpgb.fns, fns...)
	return rpgb
}

// Scan applies the selector query and scans the result into the given value.
func (rpgb *ResetPasswordGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rpgb.build.ctx, "GroupBy")
	if err := rpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResetPasswordQuery, *ResetPasswordGroupBy](ctx, rpgb.build, rpgb, rpgb.build.inters, v)
}

func (rpgb *ResetPasswordGroupBy) sqlScan(ctx context.Context, root *ResetPasswordQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rpgb.fns))
	for _, fn := range rpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rpgb.flds)+len(rpgb.fns))
		for _, f := range *rpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ResetPasswordSelect is the builder for selecting fields of ResetPassword entities.
type ResetPasswordSelect struct {
	*ResetPasswordQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rps *ResetPasswordSelect) Aggregate(fns ...AggregateFunc) *ResetPasswordSelect {
	rps.fns = append(rps.fns, fns...)
	return rps
}

// Scan applies the selector query and scans the result into the given value.
func (rps *ResetPasswordSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rps.ctx, "Select")
	if err := rps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ResetPasswordQuery, *ResetPasswordSelect](ctx, rps.ResetPasswordQuery, rps, rps.inters, v)
}

func (rps *ResetPasswordSelect) sqlScan(ctx context.Context, root *ResetPasswordQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rps.fns))
	for _, fn := range rps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
