// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/shoppingcart"
	"healthyshopper/ent/shoppingcartitem"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShoppingCartItemQuery is the builder for querying ShoppingCartItem entities.
type ShoppingCartItemQuery struct {
	config
	ctx              *QueryContext
	order            []shoppingcartitem.OrderOption
	inters           []Interceptor
	predicates       []predicate.ShoppingCartItem
	withShoppingCart *ShoppingCartQuery
	withProductItem  *ProductItemQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*ShoppingCartItem) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShoppingCartItemQuery builder.
func (sciq *ShoppingCartItemQuery) Where(ps ...predicate.ShoppingCartItem) *ShoppingCartItemQuery {
	sciq.predicates = append(sciq.predicates, ps...)
	return sciq
}

// Limit the number of records to be returned by this query.
func (sciq *ShoppingCartItemQuery) Limit(limit int) *ShoppingCartItemQuery {
	sciq.ctx.Limit = &limit
	return sciq
}

// Offset to start from.
func (sciq *ShoppingCartItemQuery) Offset(offset int) *ShoppingCartItemQuery {
	sciq.ctx.Offset = &offset
	return sciq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sciq *ShoppingCartItemQuery) Unique(unique bool) *ShoppingCartItemQuery {
	sciq.ctx.Unique = &unique
	return sciq
}

// Order specifies how the records should be ordered.
func (sciq *ShoppingCartItemQuery) Order(o ...shoppingcartitem.OrderOption) *ShoppingCartItemQuery {
	sciq.order = append(sciq.order, o...)
	return sciq
}

// QueryShoppingCart chains the current query on the "shopping_cart" edge.
func (sciq *ShoppingCartItemQuery) QueryShoppingCart() *ShoppingCartQuery {
	query := (&ShoppingCartClient{config: sciq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shoppingcartitem.Table, shoppingcartitem.FieldID, selector),
			sqlgraph.To(shoppingcart.Table, shoppingcart.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, shoppingcartitem.ShoppingCartTable, shoppingcartitem.ShoppingCartColumn),
		)
		fromU = sqlgraph.SetNeighbors(sciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryProductItem chains the current query on the "product_item" edge.
func (sciq *ShoppingCartItemQuery) QueryProductItem() *ProductItemQuery {
	query := (&ProductItemClient{config: sciq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sciq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sciq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(shoppingcartitem.Table, shoppingcartitem.FieldID, selector),
			sqlgraph.To(productitem.Table, productitem.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, shoppingcartitem.ProductItemTable, shoppingcartitem.ProductItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(sciq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShoppingCartItem entity from the query.
// Returns a *NotFoundError when no ShoppingCartItem was found.
func (sciq *ShoppingCartItemQuery) First(ctx context.Context) (*ShoppingCartItem, error) {
	nodes, err := sciq.Limit(1).All(setContextOp(ctx, sciq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{shoppingcartitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) FirstX(ctx context.Context) *ShoppingCartItem {
	node, err := sciq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShoppingCartItem ID from the query.
// Returns a *NotFoundError when no ShoppingCartItem ID was found.
func (sciq *ShoppingCartItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sciq.Limit(1).IDs(setContextOp(ctx, sciq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shoppingcartitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) FirstIDX(ctx context.Context) int {
	id, err := sciq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShoppingCartItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShoppingCartItem entity is found.
// Returns a *NotFoundError when no ShoppingCartItem entities are found.
func (sciq *ShoppingCartItemQuery) Only(ctx context.Context) (*ShoppingCartItem, error) {
	nodes, err := sciq.Limit(2).All(setContextOp(ctx, sciq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{shoppingcartitem.Label}
	default:
		return nil, &NotSingularError{shoppingcartitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) OnlyX(ctx context.Context) *ShoppingCartItem {
	node, err := sciq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShoppingCartItem ID in the query.
// Returns a *NotSingularError when more than one ShoppingCartItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (sciq *ShoppingCartItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sciq.Limit(2).IDs(setContextOp(ctx, sciq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shoppingcartitem.Label}
	default:
		err = &NotSingularError{shoppingcartitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := sciq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShoppingCartItems.
func (sciq *ShoppingCartItemQuery) All(ctx context.Context) ([]*ShoppingCartItem, error) {
	ctx = setContextOp(ctx, sciq.ctx, "All")
	if err := sciq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ShoppingCartItem, *ShoppingCartItemQuery]()
	return withInterceptors[[]*ShoppingCartItem](ctx, sciq, qr, sciq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) AllX(ctx context.Context) []*ShoppingCartItem {
	nodes, err := sciq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShoppingCartItem IDs.
func (sciq *ShoppingCartItemQuery) IDs(ctx context.Context) (ids []int, err error) {
	if sciq.ctx.Unique == nil && sciq.path != nil {
		sciq.Unique(true)
	}
	ctx = setContextOp(ctx, sciq.ctx, "IDs")
	if err = sciq.Select(shoppingcartitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) IDsX(ctx context.Context) []int {
	ids, err := sciq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sciq *ShoppingCartItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sciq.ctx, "Count")
	if err := sciq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sciq, querierCount[*ShoppingCartItemQuery](), sciq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) CountX(ctx context.Context) int {
	count, err := sciq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sciq *ShoppingCartItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sciq.ctx, "Exist")
	switch _, err := sciq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sciq *ShoppingCartItemQuery) ExistX(ctx context.Context) bool {
	exist, err := sciq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShoppingCartItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sciq *ShoppingCartItemQuery) Clone() *ShoppingCartItemQuery {
	if sciq == nil {
		return nil
	}
	return &ShoppingCartItemQuery{
		config:           sciq.config,
		ctx:              sciq.ctx.Clone(),
		order:            append([]shoppingcartitem.OrderOption{}, sciq.order...),
		inters:           append([]Interceptor{}, sciq.inters...),
		predicates:       append([]predicate.ShoppingCartItem{}, sciq.predicates...),
		withShoppingCart: sciq.withShoppingCart.Clone(),
		withProductItem:  sciq.withProductItem.Clone(),
		// clone intermediate query.
		sql:  sciq.sql.Clone(),
		path: sciq.path,
	}
}

// WithShoppingCart tells the query-builder to eager-load the nodes that are connected to
// the "shopping_cart" edge. The optional arguments are used to configure the query builder of the edge.
func (sciq *ShoppingCartItemQuery) WithShoppingCart(opts ...func(*ShoppingCartQuery)) *ShoppingCartItemQuery {
	query := (&ShoppingCartClient{config: sciq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sciq.withShoppingCart = query
	return sciq
}

// WithProductItem tells the query-builder to eager-load the nodes that are connected to
// the "product_item" edge. The optional arguments are used to configure the query builder of the edge.
func (sciq *ShoppingCartItemQuery) WithProductItem(opts ...func(*ProductItemQuery)) *ShoppingCartItemQuery {
	query := (&ProductItemClient{config: sciq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sciq.withProductItem = query
	return sciq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ShoppingCartID int `json:"shopping_cart_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ShoppingCartItem.Query().
//		GroupBy(shoppingcartitem.FieldShoppingCartID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sciq *ShoppingCartItemQuery) GroupBy(field string, fields ...string) *ShoppingCartItemGroupBy {
	sciq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShoppingCartItemGroupBy{build: sciq}
	grbuild.flds = &sciq.ctx.Fields
	grbuild.label = shoppingcartitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ShoppingCartID int `json:"shopping_cart_id,omitempty"`
//	}
//
//	client.ShoppingCartItem.Query().
//		Select(shoppingcartitem.FieldShoppingCartID).
//		Scan(ctx, &v)
func (sciq *ShoppingCartItemQuery) Select(fields ...string) *ShoppingCartItemSelect {
	sciq.ctx.Fields = append(sciq.ctx.Fields, fields...)
	sbuild := &ShoppingCartItemSelect{ShoppingCartItemQuery: sciq}
	sbuild.label = shoppingcartitem.Label
	sbuild.flds, sbuild.scan = &sciq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShoppingCartItemSelect configured with the given aggregations.
func (sciq *ShoppingCartItemQuery) Aggregate(fns ...AggregateFunc) *ShoppingCartItemSelect {
	return sciq.Select().Aggregate(fns...)
}

func (sciq *ShoppingCartItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sciq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sciq); err != nil {
				return err
			}
		}
	}
	for _, f := range sciq.ctx.Fields {
		if !shoppingcartitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sciq.path != nil {
		prev, err := sciq.path(ctx)
		if err != nil {
			return err
		}
		sciq.sql = prev
	}
	return nil
}

func (sciq *ShoppingCartItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShoppingCartItem, error) {
	var (
		nodes       = []*ShoppingCartItem{}
		_spec       = sciq.querySpec()
		loadedTypes = [2]bool{
			sciq.withShoppingCart != nil,
			sciq.withProductItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ShoppingCartItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ShoppingCartItem{config: sciq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(sciq.modifiers) > 0 {
		_spec.Modifiers = sciq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sciq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sciq.withShoppingCart; query != nil {
		if err := sciq.loadShoppingCart(ctx, query, nodes, nil,
			func(n *ShoppingCartItem, e *ShoppingCart) { n.Edges.ShoppingCart = e }); err != nil {
			return nil, err
		}
	}
	if query := sciq.withProductItem; query != nil {
		if err := sciq.loadProductItem(ctx, query, nodes, nil,
			func(n *ShoppingCartItem, e *ProductItem) { n.Edges.ProductItem = e }); err != nil {
			return nil, err
		}
	}
	for i := range sciq.loadTotal {
		if err := sciq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sciq *ShoppingCartItemQuery) loadShoppingCart(ctx context.Context, query *ShoppingCartQuery, nodes []*ShoppingCartItem, init func(*ShoppingCartItem), assign func(*ShoppingCartItem, *ShoppingCart)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ShoppingCartItem)
	for i := range nodes {
		fk := nodes[i].ShoppingCartID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(shoppingcart.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "shopping_cart_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sciq *ShoppingCartItemQuery) loadProductItem(ctx context.Context, query *ProductItemQuery, nodes []*ShoppingCartItem, init func(*ShoppingCartItem), assign func(*ShoppingCartItem, *ProductItem)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ShoppingCartItem)
	for i := range nodes {
		fk := nodes[i].ProductItemID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(productitem.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_item_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sciq *ShoppingCartItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sciq.querySpec()
	if len(sciq.modifiers) > 0 {
		_spec.Modifiers = sciq.modifiers
	}
	_spec.Node.Columns = sciq.ctx.Fields
	if len(sciq.ctx.Fields) > 0 {
		_spec.Unique = sciq.ctx.Unique != nil && *sciq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sciq.driver, _spec)
}

func (sciq *ShoppingCartItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(shoppingcartitem.Table, shoppingcartitem.Columns, sqlgraph.NewFieldSpec(shoppingcartitem.FieldID, field.TypeInt))
	_spec.From = sciq.sql
	if unique := sciq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if sciq.path != nil {
		_spec.Unique = true
	}
	if fields := sciq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shoppingcartitem.FieldID)
		for i := range fields {
			if fields[i] != shoppingcartitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if sciq.withShoppingCart != nil {
			_spec.Node.AddColumnOnce(shoppingcartitem.FieldShoppingCartID)
		}
		if sciq.withProductItem != nil {
			_spec.Node.AddColumnOnce(shoppingcartitem.FieldProductItemID)
		}
	}
	if ps := sciq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sciq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sciq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sciq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sciq *ShoppingCartItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sciq.driver.Dialect())
	t1 := builder.Table(shoppingcartitem.Table)
	columns := sciq.ctx.Fields
	if len(columns) == 0 {
		columns = shoppingcartitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sciq.sql != nil {
		selector = sciq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sciq.ctx.Unique != nil && *sciq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sciq.predicates {
		p(selector)
	}
	for _, p := range sciq.order {
		p(selector)
	}
	if offset := sciq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sciq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShoppingCartItemGroupBy is the group-by builder for ShoppingCartItem entities.
type ShoppingCartItemGroupBy struct {
	selector
	build *ShoppingCartItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scigb *ShoppingCartItemGroupBy) Aggregate(fns ...AggregateFunc) *ShoppingCartItemGroupBy {
	scigb.fns = append(scigb.fns, fns...)
	return scigb
}

// Scan applies the selector query and scans the result into the given value.
func (scigb *ShoppingCartItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scigb.build.ctx, "GroupBy")
	if err := scigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShoppingCartItemQuery, *ShoppingCartItemGroupBy](ctx, scigb.build, scigb, scigb.build.inters, v)
}

func (scigb *ShoppingCartItemGroupBy) sqlScan(ctx context.Context, root *ShoppingCartItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scigb.fns))
	for _, fn := range scigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scigb.flds)+len(scigb.fns))
		for _, f := range *scigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShoppingCartItemSelect is the builder for selecting fields of ShoppingCartItem entities.
type ShoppingCartItemSelect struct {
	*ShoppingCartItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scis *ShoppingCartItemSelect) Aggregate(fns ...AggregateFunc) *ShoppingCartItemSelect {
	scis.fns = append(scis.fns, fns...)
	return scis
}

// Scan applies the selector query and scans the result into the given value.
func (scis *ShoppingCartItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scis.ctx, "Select")
	if err := scis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShoppingCartItemQuery, *ShoppingCartItemSelect](ctx, scis.ShoppingCartItemQuery, scis, scis.inters, v)
}

func (scis *ShoppingCartItemSelect) sqlScan(ctx context.Context, root *ShoppingCartItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scis.fns))
	for _, fn := range scis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
