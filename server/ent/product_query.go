// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"healthyshopper/ent/nutritionalinformation"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/product"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/promotion"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductQuery is the builder for querying Product entities.
type ProductQuery struct {
	config
	ctx                        *QueryContext
	order                      []product.OrderOption
	inters                     []Interceptor
	predicates                 []predicate.Product
	withProductItem            *ProductItemQuery
	withPromotion              *PromotionQuery
	withNutritionalInformation *NutritionalInformationQuery
	modifiers                  []func(*sql.Selector)
	loadTotal                  []func(context.Context, []*Product) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProductQuery builder.
func (pq *ProductQuery) Where(ps ...predicate.Product) *ProductQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProductQuery) Limit(limit int) *ProductQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProductQuery) Offset(offset int) *ProductQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProductQuery) Unique(unique bool) *ProductQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProductQuery) Order(o ...product.OrderOption) *ProductQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryProductItem chains the current query on the "product_item" edge.
func (pq *ProductQuery) QueryProductItem() *ProductItemQuery {
	query := (&ProductItemClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(productitem.Table, productitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, product.ProductItemTable, product.ProductItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPromotion chains the current query on the "promotion" edge.
func (pq *ProductQuery) QueryPromotion() *PromotionQuery {
	query := (&PromotionClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(promotion.Table, promotion.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, product.PromotionTable, product.PromotionColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryNutritionalInformation chains the current query on the "nutritional_information" edge.
func (pq *ProductQuery) QueryNutritionalInformation() *NutritionalInformationQuery {
	query := (&NutritionalInformationClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(product.Table, product.FieldID, selector),
			sqlgraph.To(nutritionalinformation.Table, nutritionalinformation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, product.NutritionalInformationTable, product.NutritionalInformationColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Product entity from the query.
// Returns a *NotFoundError when no Product was found.
func (pq *ProductQuery) First(ctx context.Context) (*Product, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{product.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProductQuery) FirstX(ctx context.Context) *Product {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Product ID from the query.
// Returns a *NotFoundError when no Product ID was found.
func (pq *ProductQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{product.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProductQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Product entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Product entity is found.
// Returns a *NotFoundError when no Product entities are found.
func (pq *ProductQuery) Only(ctx context.Context) (*Product, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{product.Label}
	default:
		return nil, &NotSingularError{product.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProductQuery) OnlyX(ctx context.Context) *Product {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Product ID in the query.
// Returns a *NotSingularError when more than one Product ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProductQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{product.Label}
	default:
		err = &NotSingularError{product.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProductQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Products.
func (pq *ProductQuery) All(ctx context.Context) ([]*Product, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Product, *ProductQuery]()
	return withInterceptors[[]*Product](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProductQuery) AllX(ctx context.Context) []*Product {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Product IDs.
func (pq *ProductQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(product.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProductQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProductQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProductQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProductQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProductQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProductQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProductQuery) Clone() *ProductQuery {
	if pq == nil {
		return nil
	}
	return &ProductQuery{
		config:                     pq.config,
		ctx:                        pq.ctx.Clone(),
		order:                      append([]product.OrderOption{}, pq.order...),
		inters:                     append([]Interceptor{}, pq.inters...),
		predicates:                 append([]predicate.Product{}, pq.predicates...),
		withProductItem:            pq.withProductItem.Clone(),
		withPromotion:              pq.withPromotion.Clone(),
		withNutritionalInformation: pq.withNutritionalInformation.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithProductItem tells the query-builder to eager-load the nodes that are connected to
// the "product_item" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithProductItem(opts ...func(*ProductItemQuery)) *ProductQuery {
	query := (&ProductItemClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withProductItem = query
	return pq
}

// WithPromotion tells the query-builder to eager-load the nodes that are connected to
// the "promotion" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithPromotion(opts ...func(*PromotionQuery)) *ProductQuery {
	query := (&PromotionClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withPromotion = query
	return pq
}

// WithNutritionalInformation tells the query-builder to eager-load the nodes that are connected to
// the "nutritional_information" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProductQuery) WithNutritionalInformation(opts ...func(*NutritionalInformationQuery)) *ProductQuery {
	query := (&NutritionalInformationClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withNutritionalInformation = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Product.Query().
//		GroupBy(product.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProductQuery) GroupBy(field string, fields ...string) *ProductGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProductGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = product.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Product.Query().
//		Select(product.FieldName).
//		Scan(ctx, &v)
func (pq *ProductQuery) Select(fields ...string) *ProductSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProductSelect{ProductQuery: pq}
	sbuild.label = product.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProductSelect configured with the given aggregations.
func (pq *ProductQuery) Aggregate(fns ...AggregateFunc) *ProductSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProductQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !product.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProductQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Product, error) {
	var (
		nodes       = []*Product{}
		_spec       = pq.querySpec()
		loadedTypes = [3]bool{
			pq.withProductItem != nil,
			pq.withPromotion != nil,
			pq.withNutritionalInformation != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Product).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Product{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withProductItem; query != nil {
		if err := pq.loadProductItem(ctx, query, nodes, nil,
			func(n *Product, e *ProductItem) { n.Edges.ProductItem = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withPromotion; query != nil {
		if err := pq.loadPromotion(ctx, query, nodes, nil,
			func(n *Product, e *Promotion) { n.Edges.Promotion = e }); err != nil {
			return nil, err
		}
	}
	if query := pq.withNutritionalInformation; query != nil {
		if err := pq.loadNutritionalInformation(ctx, query, nodes, nil,
			func(n *Product, e *NutritionalInformation) { n.Edges.NutritionalInformation = e }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProductQuery) loadProductItem(ctx context.Context, query *ProductItemQuery, nodes []*Product, init func(*Product), assign func(*Product, *ProductItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Product)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(productitem.FieldProductID)
	}
	query.Where(predicate.ProductItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(product.ProductItemColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *ProductQuery) loadPromotion(ctx context.Context, query *PromotionQuery, nodes []*Product, init func(*Product), assign func(*Product, *Promotion)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Product)
	for i := range nodes {
		fk := nodes[i].PromotionID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(promotion.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "promotion_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (pq *ProductQuery) loadNutritionalInformation(ctx context.Context, query *NutritionalInformationQuery, nodes []*Product, init func(*Product), assign func(*Product, *NutritionalInformation)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Product)
	for i := range nodes {
		fk := nodes[i].NutritionalInformationID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(nutritionalinformation.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "nutritional_information_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *ProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for i := range fields {
			if fields[i] != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if pq.withPromotion != nil {
			_spec.Node.AddColumnOnce(product.FieldPromotionID)
		}
		if pq.withNutritionalInformation != nil {
			_spec.Node.AddColumnOnce(product.FieldNutritionalInformationID)
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(product.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = product.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProductGroupBy is the group-by builder for Product entities.
type ProductGroupBy struct {
	selector
	build *ProductQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProductGroupBy) Aggregate(fns ...AggregateFunc) *ProductGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProductGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductQuery, *ProductGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProductGroupBy) sqlScan(ctx context.Context, root *ProductQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProductSelect is the builder for selecting fields of Product entities.
type ProductSelect struct {
	*ProductQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProductSelect) Aggregate(fns ...AggregateFunc) *ProductSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProductSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProductQuery, *ProductSelect](ctx, ps.ProductQuery, ps, ps.inters, v)
}

func (ps *ProductSelect) sqlScan(ctx context.Context, root *ProductQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
