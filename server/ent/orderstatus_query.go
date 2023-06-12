// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"healthyshopper/ent/orderstatus"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/shoporder"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderStatusQuery is the builder for querying OrderStatus entities.
type OrderStatusQuery struct {
	config
	ctx                *QueryContext
	order              []orderstatus.OrderOption
	inters             []Interceptor
	predicates         []predicate.OrderStatus
	withShopOrder      *ShopOrderQuery
	modifiers          []func(*sql.Selector)
	loadTotal          []func(context.Context, []*OrderStatus) error
	withNamedShopOrder map[string]*ShopOrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderStatusQuery builder.
func (osq *OrderStatusQuery) Where(ps ...predicate.OrderStatus) *OrderStatusQuery {
	osq.predicates = append(osq.predicates, ps...)
	return osq
}

// Limit the number of records to be returned by this query.
func (osq *OrderStatusQuery) Limit(limit int) *OrderStatusQuery {
	osq.ctx.Limit = &limit
	return osq
}

// Offset to start from.
func (osq *OrderStatusQuery) Offset(offset int) *OrderStatusQuery {
	osq.ctx.Offset = &offset
	return osq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (osq *OrderStatusQuery) Unique(unique bool) *OrderStatusQuery {
	osq.ctx.Unique = &unique
	return osq
}

// Order specifies how the records should be ordered.
func (osq *OrderStatusQuery) Order(o ...orderstatus.OrderOption) *OrderStatusQuery {
	osq.order = append(osq.order, o...)
	return osq
}

// QueryShopOrder chains the current query on the "shop_order" edge.
func (osq *OrderStatusQuery) QueryShopOrder() *ShopOrderQuery {
	query := (&ShopOrderClient{config: osq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := osq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := osq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderstatus.Table, orderstatus.FieldID, selector),
			sqlgraph.To(shoporder.Table, shoporder.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, orderstatus.ShopOrderTable, orderstatus.ShopOrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(osq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderStatus entity from the query.
// Returns a *NotFoundError when no OrderStatus was found.
func (osq *OrderStatusQuery) First(ctx context.Context) (*OrderStatus, error) {
	nodes, err := osq.Limit(1).All(setContextOp(ctx, osq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderstatus.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (osq *OrderStatusQuery) FirstX(ctx context.Context) *OrderStatus {
	node, err := osq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderStatus ID from the query.
// Returns a *NotFoundError when no OrderStatus ID was found.
func (osq *OrderStatusQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(1).IDs(setContextOp(ctx, osq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderstatus.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (osq *OrderStatusQuery) FirstIDX(ctx context.Context) int {
	id, err := osq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderStatus entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderStatus entity is found.
// Returns a *NotFoundError when no OrderStatus entities are found.
func (osq *OrderStatusQuery) Only(ctx context.Context) (*OrderStatus, error) {
	nodes, err := osq.Limit(2).All(setContextOp(ctx, osq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderstatus.Label}
	default:
		return nil, &NotSingularError{orderstatus.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (osq *OrderStatusQuery) OnlyX(ctx context.Context) *OrderStatus {
	node, err := osq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderStatus ID in the query.
// Returns a *NotSingularError when more than one OrderStatus ID is found.
// Returns a *NotFoundError when no entities are found.
func (osq *OrderStatusQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = osq.Limit(2).IDs(setContextOp(ctx, osq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderstatus.Label}
	default:
		err = &NotSingularError{orderstatus.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (osq *OrderStatusQuery) OnlyIDX(ctx context.Context) int {
	id, err := osq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderStatusSlice.
func (osq *OrderStatusQuery) All(ctx context.Context) ([]*OrderStatus, error) {
	ctx = setContextOp(ctx, osq.ctx, "All")
	if err := osq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderStatus, *OrderStatusQuery]()
	return withInterceptors[[]*OrderStatus](ctx, osq, qr, osq.inters)
}

// AllX is like All, but panics if an error occurs.
func (osq *OrderStatusQuery) AllX(ctx context.Context) []*OrderStatus {
	nodes, err := osq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderStatus IDs.
func (osq *OrderStatusQuery) IDs(ctx context.Context) (ids []int, err error) {
	if osq.ctx.Unique == nil && osq.path != nil {
		osq.Unique(true)
	}
	ctx = setContextOp(ctx, osq.ctx, "IDs")
	if err = osq.Select(orderstatus.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (osq *OrderStatusQuery) IDsX(ctx context.Context) []int {
	ids, err := osq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (osq *OrderStatusQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, osq.ctx, "Count")
	if err := osq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, osq, querierCount[*OrderStatusQuery](), osq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (osq *OrderStatusQuery) CountX(ctx context.Context) int {
	count, err := osq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (osq *OrderStatusQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, osq.ctx, "Exist")
	switch _, err := osq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (osq *OrderStatusQuery) ExistX(ctx context.Context) bool {
	exist, err := osq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderStatusQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (osq *OrderStatusQuery) Clone() *OrderStatusQuery {
	if osq == nil {
		return nil
	}
	return &OrderStatusQuery{
		config:        osq.config,
		ctx:           osq.ctx.Clone(),
		order:         append([]orderstatus.OrderOption{}, osq.order...),
		inters:        append([]Interceptor{}, osq.inters...),
		predicates:    append([]predicate.OrderStatus{}, osq.predicates...),
		withShopOrder: osq.withShopOrder.Clone(),
		// clone intermediate query.
		sql:  osq.sql.Clone(),
		path: osq.path,
	}
}

// WithShopOrder tells the query-builder to eager-load the nodes that are connected to
// the "shop_order" edge. The optional arguments are used to configure the query builder of the edge.
func (osq *OrderStatusQuery) WithShopOrder(opts ...func(*ShopOrderQuery)) *OrderStatusQuery {
	query := (&ShopOrderClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	osq.withShopOrder = query
	return osq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderStatus.Query().
//		GroupBy(orderstatus.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (osq *OrderStatusQuery) GroupBy(field string, fields ...string) *OrderStatusGroupBy {
	osq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderStatusGroupBy{build: osq}
	grbuild.flds = &osq.ctx.Fields
	grbuild.label = orderstatus.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status string `json:"status,omitempty"`
//	}
//
//	client.OrderStatus.Query().
//		Select(orderstatus.FieldStatus).
//		Scan(ctx, &v)
func (osq *OrderStatusQuery) Select(fields ...string) *OrderStatusSelect {
	osq.ctx.Fields = append(osq.ctx.Fields, fields...)
	sbuild := &OrderStatusSelect{OrderStatusQuery: osq}
	sbuild.label = orderstatus.Label
	sbuild.flds, sbuild.scan = &osq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderStatusSelect configured with the given aggregations.
func (osq *OrderStatusQuery) Aggregate(fns ...AggregateFunc) *OrderStatusSelect {
	return osq.Select().Aggregate(fns...)
}

func (osq *OrderStatusQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range osq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, osq); err != nil {
				return err
			}
		}
	}
	for _, f := range osq.ctx.Fields {
		if !orderstatus.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if osq.path != nil {
		prev, err := osq.path(ctx)
		if err != nil {
			return err
		}
		osq.sql = prev
	}
	return nil
}

func (osq *OrderStatusQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderStatus, error) {
	var (
		nodes       = []*OrderStatus{}
		_spec       = osq.querySpec()
		loadedTypes = [1]bool{
			osq.withShopOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderStatus).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderStatus{config: osq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, osq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := osq.withShopOrder; query != nil {
		if err := osq.loadShopOrder(ctx, query, nodes,
			func(n *OrderStatus) { n.Edges.ShopOrder = []*ShopOrder{} },
			func(n *OrderStatus, e *ShopOrder) { n.Edges.ShopOrder = append(n.Edges.ShopOrder, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range osq.withNamedShopOrder {
		if err := osq.loadShopOrder(ctx, query, nodes,
			func(n *OrderStatus) { n.appendNamedShopOrder(name) },
			func(n *OrderStatus, e *ShopOrder) { n.appendNamedShopOrder(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range osq.loadTotal {
		if err := osq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (osq *OrderStatusQuery) loadShopOrder(ctx context.Context, query *ShopOrderQuery, nodes []*OrderStatus, init func(*OrderStatus), assign func(*OrderStatus, *ShopOrder)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*OrderStatus)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(shoporder.FieldOrderStatusID)
	}
	query.Where(predicate.ShopOrder(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(orderstatus.ShopOrderColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.OrderStatusID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "order_status_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (osq *OrderStatusQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := osq.querySpec()
	if len(osq.modifiers) > 0 {
		_spec.Modifiers = osq.modifiers
	}
	_spec.Node.Columns = osq.ctx.Fields
	if len(osq.ctx.Fields) > 0 {
		_spec.Unique = osq.ctx.Unique != nil && *osq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, osq.driver, _spec)
}

func (osq *OrderStatusQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderstatus.Table, orderstatus.Columns, sqlgraph.NewFieldSpec(orderstatus.FieldID, field.TypeInt))
	_spec.From = osq.sql
	if unique := osq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if osq.path != nil {
		_spec.Unique = true
	}
	if fields := osq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderstatus.FieldID)
		for i := range fields {
			if fields[i] != orderstatus.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := osq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := osq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := osq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := osq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (osq *OrderStatusQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(osq.driver.Dialect())
	t1 := builder.Table(orderstatus.Table)
	columns := osq.ctx.Fields
	if len(columns) == 0 {
		columns = orderstatus.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if osq.sql != nil {
		selector = osq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if osq.ctx.Unique != nil && *osq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range osq.predicates {
		p(selector)
	}
	for _, p := range osq.order {
		p(selector)
	}
	if offset := osq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := osq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedShopOrder tells the query-builder to eager-load the nodes that are connected to the "shop_order"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (osq *OrderStatusQuery) WithNamedShopOrder(name string, opts ...func(*ShopOrderQuery)) *OrderStatusQuery {
	query := (&ShopOrderClient{config: osq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if osq.withNamedShopOrder == nil {
		osq.withNamedShopOrder = make(map[string]*ShopOrderQuery)
	}
	osq.withNamedShopOrder[name] = query
	return osq
}

// OrderStatusGroupBy is the group-by builder for OrderStatus entities.
type OrderStatusGroupBy struct {
	selector
	build *OrderStatusQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (osgb *OrderStatusGroupBy) Aggregate(fns ...AggregateFunc) *OrderStatusGroupBy {
	osgb.fns = append(osgb.fns, fns...)
	return osgb
}

// Scan applies the selector query and scans the result into the given value.
func (osgb *OrderStatusGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, osgb.build.ctx, "GroupBy")
	if err := osgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderStatusQuery, *OrderStatusGroupBy](ctx, osgb.build, osgb, osgb.build.inters, v)
}

func (osgb *OrderStatusGroupBy) sqlScan(ctx context.Context, root *OrderStatusQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(osgb.fns))
	for _, fn := range osgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*osgb.flds)+len(osgb.fns))
		for _, f := range *osgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*osgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := osgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderStatusSelect is the builder for selecting fields of OrderStatus entities.
type OrderStatusSelect struct {
	*OrderStatusQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oss *OrderStatusSelect) Aggregate(fns ...AggregateFunc) *OrderStatusSelect {
	oss.fns = append(oss.fns, fns...)
	return oss
}

// Scan applies the selector query and scans the result into the given value.
func (oss *OrderStatusSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oss.ctx, "Select")
	if err := oss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderStatusQuery, *OrderStatusSelect](ctx, oss.OrderStatusQuery, oss, oss.inters, v)
}

func (oss *OrderStatusSelect) sqlScan(ctx context.Context, root *OrderStatusQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oss.fns))
	for _, fn := range oss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
