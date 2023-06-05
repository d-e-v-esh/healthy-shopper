// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"healthyshopper/ent/orderline"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/product"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/shoppingcartitem"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductItemQuery is the builder for querying ProductItem entities.
type ProductItemQuery struct {
	config
	ctx                       *QueryContext
	order                     []productitem.OrderOption
	inters                    []Interceptor
	predicates                []predicate.ProductItem
	withProduct               *ProductQuery
	withOrderLine             *OrderLineQuery
	withShoppingCartItem      *ShoppingCartItemQuery
	modifiers                 []func(*sql.Selector)
	loadTotal                 []func(context.Context, []*ProductItem) error
	withNamedOrderLine        map[string]*OrderLineQuery
	withNamedShoppingCartItem map[string]*ShoppingCartItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductItemQuery builder.
func (piq *ProductItemQuery) Where(ps ...predicate.ProductItem) *ProductItemQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit the number of records to be returned by this query.
func (piq *ProductItemQuery) Limit(limit int) *ProductItemQuery {
	piq.ctx.Limit = &limit
	return piq
}

// Offset to start from.
func (piq *ProductItemQuery) Offset(offset int) *ProductItemQuery {
	piq.ctx.Offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *ProductItemQuery) Unique(unique bool) *ProductItemQuery {
	piq.ctx.Unique = &unique
	return piq
}

// Order specifies how the records should be ordered.
func (piq *ProductItemQuery) Order(o ...productitem.OrderOption) *ProductItemQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// QueryProduct chains the current query on the "product" edge.
func (piq *ProductItemQuery) QueryProduct() *ProductQuery {
	query := (&ProductClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productitem.Table, productitem.FieldID, selector),
			sqlgraph.To(product.Table, product.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, productitem.ProductTable, productitem.ProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderLine chains the current query on the "order_line" edge.
func (piq *ProductItemQuery) QueryOrderLine() *OrderLineQuery {
	query := (&OrderLineClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productitem.Table, productitem.FieldID, selector),
			sqlgraph.To(orderline.Table, orderline.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, productitem.OrderLineTable, productitem.OrderLineColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryShoppingCartItem chains the current query on the "shopping_cart_item" edge.
func (piq *ProductItemQuery) QueryShoppingCartItem() *ShoppingCartItemQuery {
	query := (&ShoppingCartItemClient{config: piq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := piq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(productitem.Table, productitem.FieldID, selector),
			sqlgraph.To(shoppingcartitem.Table, shoppingcartitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, productitem.ShoppingCartItemTable, productitem.ShoppingCartItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(piq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ProductItem entity from the query.
// Returns a *NotFoundError when no ProductItem was found.
func (piq *ProductItemQuery) First(ctx context.Context) (*ProductItem, error) {
	nodes, err := piq.Limit(1).All(setContextOp(ctx, piq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{productitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *ProductItemQuery) FirstX(ctx context.Context) *ProductItem {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProductItem ID from the query.
// Returns a *NotFoundError when no ProductItem ID was found.
func (piq *ProductItemQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(1).IDs(setContextOp(ctx, piq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{productitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *ProductItemQuery) FirstIDX(ctx context.Context) int {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProductItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProductItem entity is found.
// Returns a *NotFoundError when no ProductItem entities are found.
func (piq *ProductItemQuery) Only(ctx context.Context) (*ProductItem, error) {
	nodes, err := piq.Limit(2).All(setContextOp(ctx, piq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{productitem.Label}
	default:
		return nil, &NotSingularError{productitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *ProductItemQuery) OnlyX(ctx context.Context) *ProductItem {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProductItem ID in the query.
// Returns a *NotSingularError when more than one ProductItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *ProductItemQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(2).IDs(setContextOp(ctx, piq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{productitem.Label}
	default:
		err = &NotSingularError{productitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *ProductItemQuery) OnlyIDX(ctx context.Context) int {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProductItems.
func (piq *ProductItemQuery) All(ctx context.Context) ([]*ProductItem, error) {
	ctx = setContextOp(ctx, piq.ctx, "All")
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ProductItem, *ProductItemQuery]()
	return withInterceptors[[]*ProductItem](ctx, piq, qr, piq.inters)
}

// AllX is like All, but panics if an error occurs.
func (piq *ProductItemQuery) AllX(ctx context.Context) []*ProductItem {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProductItem IDs.
func (piq *ProductItemQuery) IDs(ctx context.Context) (ids []int, err error) {
	if piq.ctx.Unique == nil && piq.path != nil {
		piq.Unique(true)
	}
	ctx = setContextOp(ctx, piq.ctx, "IDs")
	if err = piq.Select(productitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *ProductItemQuery) IDsX(ctx context.Context) []int {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *ProductItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, piq.ctx, "Count")
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, piq, querierCount[*ProductItemQuery](), piq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (piq *ProductItemQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *ProductItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, piq.ctx, "Exist")
	switch _, err := piq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *ProductItemQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *ProductItemQuery) Clone() *ProductItemQuery {
	if piq == nil {
		return nil
	}
	return &ProductItemQuery{
		config:               piq.config,
		ctx:                  piq.ctx.Clone(),
		order:                append([]productitem.OrderOption{}, piq.order...),
		inters:               append([]Interceptor{}, piq.inters...),
		predicates:           append([]predicate.ProductItem{}, piq.predicates...),
		withProduct:          piq.withProduct.Clone(),
		withOrderLine:        piq.withOrderLine.Clone(),
		withShoppingCartItem: piq.withShoppingCartItem.Clone(),
		// clone intermediate query.
		sql:  piq.sql.Clone(),
		path: piq.path,
	}
}

// WithProduct tells the query-builder to eager-load the nodes that are connected to
// the "product" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProductItemQuery) WithProduct(opts ...func(*ProductQuery)) *ProductItemQuery {
	query := (&ProductClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withProduct = query
	return piq
}

// WithOrderLine tells the query-builder to eager-load the nodes that are connected to
// the "order_line" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProductItemQuery) WithOrderLine(opts ...func(*OrderLineQuery)) *ProductItemQuery {
	query := (&OrderLineClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withOrderLine = query
	return piq
}

// WithShoppingCartItem tells the query-builder to eager-load the nodes that are connected to
// the "shopping_cart_item" edge. The optional arguments are used to configure the query builder of the edge.
func (piq *ProductItemQuery) WithShoppingCartItem(opts ...func(*ShoppingCartItemQuery)) *ProductItemQuery {
	query := (&ShoppingCartItemClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	piq.withShoppingCartItem = query
	return piq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProductID int `json:"product_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProductItem.Query().
//		GroupBy(productitem.FieldProductID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (piq *ProductItemQuery) GroupBy(field string, fields ...string) *ProductItemGroupBy {
	piq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductItemGroupBy{build: piq}
	grbuild.flds = &piq.ctx.Fields
	grbuild.label = productitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProductID int `json:"product_id,omitempty"`
//	}
//
//	client.ProductItem.Query().
//		Select(productitem.FieldProductID).
//		Scan(ctx, &v)
func (piq *ProductItemQuery) Select(fields ...string) *ProductItemSelect {
	piq.ctx.Fields = append(piq.ctx.Fields, fields...)
	sbuild := &ProductItemSelect{ProductItemQuery: piq}
	sbuild.label = productitem.Label
	sbuild.flds, sbuild.scan = &piq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductItemSelect configured with the given aggregations.
func (piq *ProductItemQuery) Aggregate(fns ...AggregateFunc) *ProductItemSelect {
	return piq.Select().Aggregate(fns...)
}

func (piq *ProductItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range piq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, piq); err != nil {
				return err
			}
		}
	}
	for _, f := range piq.ctx.Fields {
		if !productitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *ProductItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProductItem, error) {
	var (
		nodes       = []*ProductItem{}
		_spec       = piq.querySpec()
		loadedTypes = [3]bool{
			piq.withProduct != nil,
			piq.withOrderLine != nil,
			piq.withShoppingCartItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ProductItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ProductItem{config: piq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(piq.modifiers) > 0 {
		_spec.Modifiers = piq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := piq.withProduct; query != nil {
		if err := piq.loadProduct(ctx, query, nodes, nil,
			func(n *ProductItem, e *Product) { n.Edges.Product = e }); err != nil {
			return nil, err
		}
	}
	if query := piq.withOrderLine; query != nil {
		if err := piq.loadOrderLine(ctx, query, nodes,
			func(n *ProductItem) { n.Edges.OrderLine = []*OrderLine{} },
			func(n *ProductItem, e *OrderLine) { n.Edges.OrderLine = append(n.Edges.OrderLine, e) }); err != nil {
			return nil, err
		}
	}
	if query := piq.withShoppingCartItem; query != nil {
		if err := piq.loadShoppingCartItem(ctx, query, nodes,
			func(n *ProductItem) { n.Edges.ShoppingCartItem = []*ShoppingCartItem{} },
			func(n *ProductItem, e *ShoppingCartItem) {
				n.Edges.ShoppingCartItem = append(n.Edges.ShoppingCartItem, e)
			}); err != nil {
			return nil, err
		}
	}
	for name, query := range piq.withNamedOrderLine {
		if err := piq.loadOrderLine(ctx, query, nodes,
			func(n *ProductItem) { n.appendNamedOrderLine(name) },
			func(n *ProductItem, e *OrderLine) { n.appendNamedOrderLine(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range piq.withNamedShoppingCartItem {
		if err := piq.loadShoppingCartItem(ctx, query, nodes,
			func(n *ProductItem) { n.appendNamedShoppingCartItem(name) },
			func(n *ProductItem, e *ShoppingCartItem) { n.appendNamedShoppingCartItem(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range piq.loadTotal {
		if err := piq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (piq *ProductItemQuery) loadProduct(ctx context.Context, query *ProductQuery, nodes []*ProductItem, init func(*ProductItem), assign func(*ProductItem, *Product)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ProductItem)
	for i := range nodes {
		fk := nodes[i].ProductID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(product.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "product_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (piq *ProductItemQuery) loadOrderLine(ctx context.Context, query *OrderLineQuery, nodes []*ProductItem, init func(*ProductItem), assign func(*ProductItem, *OrderLine)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ProductItem)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(orderline.FieldProductItemID)
	}
	query.Where(predicate.OrderLine(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(productitem.OrderLineColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProductItemID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "product_item_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (piq *ProductItemQuery) loadShoppingCartItem(ctx context.Context, query *ShoppingCartItemQuery, nodes []*ProductItem, init func(*ProductItem), assign func(*ProductItem, *ShoppingCartItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ProductItem)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(shoppingcartitem.FieldProductItemID)
	}
	query.Where(predicate.ShoppingCartItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(productitem.ShoppingCartItemColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProductItemID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "product_item_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (piq *ProductItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	if len(piq.modifiers) > 0 {
		_spec.Modifiers = piq.modifiers
	}
	_spec.Node.Columns = piq.ctx.Fields
	if len(piq.ctx.Fields) > 0 {
		_spec.Unique = piq.ctx.Unique != nil && *piq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *ProductItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(productitem.Table, productitem.Columns, sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeInt))
	_spec.From = piq.sql
	if unique := piq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if piq.path != nil {
		_spec.Unique = true
	}
	if fields := piq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productitem.FieldID)
		for i := range fields {
			if fields[i] != productitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if piq.withProduct != nil {
			_spec.Node.AddColumnOnce(productitem.FieldProductID)
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *ProductItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(productitem.Table)
	columns := piq.ctx.Fields
	if len(columns) == 0 {
		columns = productitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.ctx.Unique != nil && *piq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedOrderLine tells the query-builder to eager-load the nodes that are connected to the "order_line"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (piq *ProductItemQuery) WithNamedOrderLine(name string, opts ...func(*OrderLineQuery)) *ProductItemQuery {
	query := (&OrderLineClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if piq.withNamedOrderLine == nil {
		piq.withNamedOrderLine = make(map[string]*OrderLineQuery)
	}
	piq.withNamedOrderLine[name] = query
	return piq
}

// WithNamedShoppingCartItem tells the query-builder to eager-load the nodes that are connected to the "shopping_cart_item"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (piq *ProductItemQuery) WithNamedShoppingCartItem(name string, opts ...func(*ShoppingCartItemQuery)) *ProductItemQuery {
	query := (&ShoppingCartItemClient{config: piq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if piq.withNamedShoppingCartItem == nil {
		piq.withNamedShoppingCartItem = make(map[string]*ShoppingCartItemQuery)
	}
	piq.withNamedShoppingCartItem[name] = query
	return piq
}

// ProductItemGroupBy is the group-by builder for ProductItem entities.
type ProductItemGroupBy struct {
	selector
	build *ProductItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *ProductItemGroupBy) Aggregate(fns ...AggregateFunc) *ProductItemGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the selector query and scans the result into the given value.
func (pigb *ProductItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pigb.build.ctx, "GroupBy")
	if err := pigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductItemQuery, *ProductItemGroupBy](ctx, pigb.build, pigb, pigb.build.inters, v)
}

func (pigb *ProductItemGroupBy) sqlScan(ctx context.Context, root *ProductItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pigb.flds)+len(pigb.fns))
		for _, f := range *pigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductItemSelect is the builder for selecting fields of ProductItem entities.
type ProductItemSelect struct {
	*ProductItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pis *ProductItemSelect) Aggregate(fns ...AggregateFunc) *ProductItemSelect {
	pis.fns = append(pis.fns, fns...)
	return pis
}

// Scan applies the selector query and scans the result into the given value.
func (pis *ProductItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pis.ctx, "Select")
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductItemQuery, *ProductItemSelect](ctx, pis.ProductItemQuery, pis, pis.inters, v)
}

func (pis *ProductItemSelect) sqlScan(ctx context.Context, root *ProductItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pis.fns))
	for _, fn := range pis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
