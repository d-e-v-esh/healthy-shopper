// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/user"
	"healthyshopper/ent/userreview"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserReviewQuery is the builder for querying UserReview entities.
type UserReviewQuery struct {
	config
	ctx        *QueryContext
	order      []userreview.OrderOption
	inters     []Interceptor
	predicates []predicate.UserReview
	withUser   *UserQuery
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*UserReview) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserReviewQuery builder.
func (urq *UserReviewQuery) Where(ps ...predicate.UserReview) *UserReviewQuery {
	urq.predicates = append(urq.predicates, ps...)
	return urq
}

// Limit the number of records to be returned by this query.
func (urq *UserReviewQuery) Limit(limit int) *UserReviewQuery {
	urq.ctx.Limit = &limit
	return urq
}

// Offset to start from.
func (urq *UserReviewQuery) Offset(offset int) *UserReviewQuery {
	urq.ctx.Offset = &offset
	return urq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (urq *UserReviewQuery) Unique(unique bool) *UserReviewQuery {
	urq.ctx.Unique = &unique
	return urq
}

// Order specifies how the records should be ordered.
func (urq *UserReviewQuery) Order(o ...userreview.OrderOption) *UserReviewQuery {
	urq.order = append(urq.order, o...)
	return urq
}

// QueryUser chains the current query on the "user" edge.
func (urq *UserReviewQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: urq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := urq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := urq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userreview.Table, userreview.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userreview.UserTable, userreview.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(urq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserReview entity from the query.
// Returns a *NotFoundError when no UserReview was found.
func (urq *UserReviewQuery) First(ctx context.Context) (*UserReview, error) {
	nodes, err := urq.Limit(1).All(setContextOp(ctx, urq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userreview.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (urq *UserReviewQuery) FirstX(ctx context.Context) *UserReview {
	node, err := urq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserReview ID from the query.
// Returns a *NotFoundError when no UserReview ID was found.
func (urq *UserReviewQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = urq.Limit(1).IDs(setContextOp(ctx, urq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userreview.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (urq *UserReviewQuery) FirstIDX(ctx context.Context) int {
	id, err := urq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserReview entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserReview entity is found.
// Returns a *NotFoundError when no UserReview entities are found.
func (urq *UserReviewQuery) Only(ctx context.Context) (*UserReview, error) {
	nodes, err := urq.Limit(2).All(setContextOp(ctx, urq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userreview.Label}
	default:
		return nil, &NotSingularError{userreview.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (urq *UserReviewQuery) OnlyX(ctx context.Context) *UserReview {
	node, err := urq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserReview ID in the query.
// Returns a *NotSingularError when more than one UserReview ID is found.
// Returns a *NotFoundError when no entities are found.
func (urq *UserReviewQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = urq.Limit(2).IDs(setContextOp(ctx, urq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userreview.Label}
	default:
		err = &NotSingularError{userreview.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (urq *UserReviewQuery) OnlyIDX(ctx context.Context) int {
	id, err := urq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserReviews.
func (urq *UserReviewQuery) All(ctx context.Context) ([]*UserReview, error) {
	ctx = setContextOp(ctx, urq.ctx, "All")
	if err := urq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserReview, *UserReviewQuery]()
	return withInterceptors[[]*UserReview](ctx, urq, qr, urq.inters)
}

// AllX is like All, but panics if an error occurs.
func (urq *UserReviewQuery) AllX(ctx context.Context) []*UserReview {
	nodes, err := urq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserReview IDs.
func (urq *UserReviewQuery) IDs(ctx context.Context) (ids []int, err error) {
	if urq.ctx.Unique == nil && urq.path != nil {
		urq.Unique(true)
	}
	ctx = setContextOp(ctx, urq.ctx, "IDs")
	if err = urq.Select(userreview.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (urq *UserReviewQuery) IDsX(ctx context.Context) []int {
	ids, err := urq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (urq *UserReviewQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, urq.ctx, "Count")
	if err := urq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, urq, querierCount[*UserReviewQuery](), urq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (urq *UserReviewQuery) CountX(ctx context.Context) int {
	count, err := urq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (urq *UserReviewQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, urq.ctx, "Exist")
	switch _, err := urq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (urq *UserReviewQuery) ExistX(ctx context.Context) bool {
	exist, err := urq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserReviewQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (urq *UserReviewQuery) Clone() *UserReviewQuery {
	if urq == nil {
		return nil
	}
	return &UserReviewQuery{
		config:     urq.config,
		ctx:        urq.ctx.Clone(),
		order:      append([]userreview.OrderOption{}, urq.order...),
		inters:     append([]Interceptor{}, urq.inters...),
		predicates: append([]predicate.UserReview{}, urq.predicates...),
		withUser:   urq.withUser.Clone(),
		// clone intermediate query.
		sql:  urq.sql.Clone(),
		path: urq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (urq *UserReviewQuery) WithUser(opts ...func(*UserQuery)) *UserReviewQuery {
	query := (&UserClient{config: urq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	urq.withUser = query
	return urq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserReview.Query().
//		GroupBy(userreview.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (urq *UserReviewQuery) GroupBy(field string, fields ...string) *UserReviewGroupBy {
	urq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserReviewGroupBy{build: urq}
	grbuild.flds = &urq.ctx.Fields
	grbuild.label = userreview.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.UserReview.Query().
//		Select(userreview.FieldUserID).
//		Scan(ctx, &v)
func (urq *UserReviewQuery) Select(fields ...string) *UserReviewSelect {
	urq.ctx.Fields = append(urq.ctx.Fields, fields...)
	sbuild := &UserReviewSelect{UserReviewQuery: urq}
	sbuild.label = userreview.Label
	sbuild.flds, sbuild.scan = &urq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserReviewSelect configured with the given aggregations.
func (urq *UserReviewQuery) Aggregate(fns ...AggregateFunc) *UserReviewSelect {
	return urq.Select().Aggregate(fns...)
}

func (urq *UserReviewQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range urq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, urq); err != nil {
				return err
			}
		}
	}
	for _, f := range urq.ctx.Fields {
		if !userreview.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if urq.path != nil {
		prev, err := urq.path(ctx)
		if err != nil {
			return err
		}
		urq.sql = prev
	}
	return nil
}

func (urq *UserReviewQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserReview, error) {
	var (
		nodes       = []*UserReview{}
		_spec       = urq.querySpec()
		loadedTypes = [1]bool{
			urq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserReview).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserReview{config: urq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, urq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := urq.withUser; query != nil {
		if err := urq.loadUser(ctx, query, nodes, nil,
			func(n *UserReview, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	for i := range urq.loadTotal {
		if err := urq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (urq *UserReviewQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*UserReview, init func(*UserReview), assign func(*UserReview, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UserReview)
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

func (urq *UserReviewQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := urq.querySpec()
	if len(urq.modifiers) > 0 {
		_spec.Modifiers = urq.modifiers
	}
	_spec.Node.Columns = urq.ctx.Fields
	if len(urq.ctx.Fields) > 0 {
		_spec.Unique = urq.ctx.Unique != nil && *urq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, urq.driver, _spec)
}

func (urq *UserReviewQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(userreview.Table, userreview.Columns, sqlgraph.NewFieldSpec(userreview.FieldID, field.TypeInt))
	_spec.From = urq.sql
	if unique := urq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if urq.path != nil {
		_spec.Unique = true
	}
	if fields := urq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userreview.FieldID)
		for i := range fields {
			if fields[i] != userreview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if urq.withUser != nil {
			_spec.Node.AddColumnOnce(userreview.FieldUserID)
		}
	}
	if ps := urq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := urq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := urq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := urq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (urq *UserReviewQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(urq.driver.Dialect())
	t1 := builder.Table(userreview.Table)
	columns := urq.ctx.Fields
	if len(columns) == 0 {
		columns = userreview.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if urq.sql != nil {
		selector = urq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if urq.ctx.Unique != nil && *urq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range urq.predicates {
		p(selector)
	}
	for _, p := range urq.order {
		p(selector)
	}
	if offset := urq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := urq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserReviewGroupBy is the group-by builder for UserReview entities.
type UserReviewGroupBy struct {
	selector
	build *UserReviewQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (urgb *UserReviewGroupBy) Aggregate(fns ...AggregateFunc) *UserReviewGroupBy {
	urgb.fns = append(urgb.fns, fns...)
	return urgb
}

// Scan applies the selector query and scans the result into the given value.
func (urgb *UserReviewGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urgb.build.ctx, "GroupBy")
	if err := urgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserReviewQuery, *UserReviewGroupBy](ctx, urgb.build, urgb, urgb.build.inters, v)
}

func (urgb *UserReviewGroupBy) sqlScan(ctx context.Context, root *UserReviewQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(urgb.fns))
	for _, fn := range urgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*urgb.flds)+len(urgb.fns))
		for _, f := range *urgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*urgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserReviewSelect is the builder for selecting fields of UserReview entities.
type UserReviewSelect struct {
	*UserReviewQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (urs *UserReviewSelect) Aggregate(fns ...AggregateFunc) *UserReviewSelect {
	urs.fns = append(urs.fns, fns...)
	return urs
}

// Scan applies the selector query and scans the result into the given value.
func (urs *UserReviewSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, urs.ctx, "Select")
	if err := urs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserReviewQuery, *UserReviewSelect](ctx, urs.UserReviewQuery, urs, urs.inters, v)
}

func (urs *UserReviewSelect) sqlScan(ctx context.Context, root *UserReviewQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(urs.fns))
	for _, fn := range urs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*urs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := urs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
