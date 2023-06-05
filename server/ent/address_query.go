// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"healthyshopper/ent/address"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/useraddress"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AddressQuery is the builder for querying Address entities.
type AddressQuery struct {
	config
	ctx                  *QueryContext
	order                []address.OrderOption
	inters               []Interceptor
	predicates           []predicate.Address
	withUserAddress      *UserAddressQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Address) error
	withNamedUserAddress map[string]*UserAddressQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AddressQuery builder.
func (aq *AddressQuery) Where(ps ...predicate.Address) *AddressQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AddressQuery) Limit(limit int) *AddressQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AddressQuery) Offset(offset int) *AddressQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AddressQuery) Unique(unique bool) *AddressQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AddressQuery) Order(o ...address.OrderOption) *AddressQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryUserAddress chains the current query on the "user_address" edge.
func (aq *AddressQuery) QueryUserAddress() *UserAddressQuery {
	query := (&UserAddressClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(address.Table, address.FieldID, selector),
			sqlgraph.To(useraddress.Table, useraddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, address.UserAddressTable, address.UserAddressColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Address entity from the query.
// Returns a *NotFoundError when no Address was found.
func (aq *AddressQuery) First(ctx context.Context) (*Address, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{address.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AddressQuery) FirstX(ctx context.Context) *Address {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Address ID from the query.
// Returns a *NotFoundError when no Address ID was found.
func (aq *AddressQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{address.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AddressQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Address entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Address entity is found.
// Returns a *NotFoundError when no Address entities are found.
func (aq *AddressQuery) Only(ctx context.Context) (*Address, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{address.Label}
	default:
		return nil, &NotSingularError{address.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AddressQuery) OnlyX(ctx context.Context) *Address {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Address ID in the query.
// Returns a *NotSingularError when more than one Address ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AddressQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{address.Label}
	default:
		err = &NotSingularError{address.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AddressQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Addresses.
func (aq *AddressQuery) All(ctx context.Context) ([]*Address, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Address, *AddressQuery]()
	return withInterceptors[[]*Address](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AddressQuery) AllX(ctx context.Context) []*Address {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Address IDs.
func (aq *AddressQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(address.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AddressQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AddressQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AddressQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AddressQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AddressQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AddressQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AddressQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AddressQuery) Clone() *AddressQuery {
	if aq == nil {
		return nil
	}
	return &AddressQuery{
		config:          aq.config,
		ctx:             aq.ctx.Clone(),
		order:           append([]address.OrderOption{}, aq.order...),
		inters:          append([]Interceptor{}, aq.inters...),
		predicates:      append([]predicate.Address{}, aq.predicates...),
		withUserAddress: aq.withUserAddress.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithUserAddress tells the query-builder to eager-load the nodes that are connected to
// the "user_address" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AddressQuery) WithUserAddress(opts ...func(*UserAddressQuery)) *AddressQuery {
	query := (&UserAddressClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withUserAddress = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		PhoneNumber string `json:"phone_number,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Address.Query().
//		GroupBy(address.FieldPhoneNumber).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AddressQuery) GroupBy(field string, fields ...string) *AddressGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AddressGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = address.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		PhoneNumber string `json:"phone_number,omitempty"`
//	}
//
//	client.Address.Query().
//		Select(address.FieldPhoneNumber).
//		Scan(ctx, &v)
func (aq *AddressQuery) Select(fields ...string) *AddressSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AddressSelect{AddressQuery: aq}
	sbuild.label = address.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AddressSelect configured with the given aggregations.
func (aq *AddressQuery) Aggregate(fns ...AggregateFunc) *AddressSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AddressQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !address.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AddressQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Address, error) {
	var (
		nodes       = []*Address{}
		_spec       = aq.querySpec()
		loadedTypes = [1]bool{
			aq.withUserAddress != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Address).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Address{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withUserAddress; query != nil {
		if err := aq.loadUserAddress(ctx, query, nodes,
			func(n *Address) { n.Edges.UserAddress = []*UserAddress{} },
			func(n *Address, e *UserAddress) { n.Edges.UserAddress = append(n.Edges.UserAddress, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range aq.withNamedUserAddress {
		if err := aq.loadUserAddress(ctx, query, nodes,
			func(n *Address) { n.appendNamedUserAddress(name) },
			func(n *Address, e *UserAddress) { n.appendNamedUserAddress(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range aq.loadTotal {
		if err := aq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AddressQuery) loadUserAddress(ctx context.Context, query *UserAddressQuery, nodes []*Address, init func(*Address), assign func(*Address, *UserAddress)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Address)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(useraddress.FieldAddressID)
	}
	query.Where(predicate.UserAddress(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(address.UserAddressColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AddressID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "address_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (aq *AddressQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	if len(aq.modifiers) > 0 {
		_spec.Modifiers = aq.modifiers
	}
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AddressQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(address.Table, address.Columns, sqlgraph.NewFieldSpec(address.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, address.FieldID)
		for i := range fields {
			if fields[i] != address.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AddressQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(address.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = address.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedUserAddress tells the query-builder to eager-load the nodes that are connected to the "user_address"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (aq *AddressQuery) WithNamedUserAddress(name string, opts ...func(*UserAddressQuery)) *AddressQuery {
	query := (&UserAddressClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if aq.withNamedUserAddress == nil {
		aq.withNamedUserAddress = make(map[string]*UserAddressQuery)
	}
	aq.withNamedUserAddress[name] = query
	return aq
}

// AddressGroupBy is the group-by builder for Address entities.
type AddressGroupBy struct {
	selector
	build *AddressQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AddressGroupBy) Aggregate(fns ...AggregateFunc) *AddressGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AddressGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AddressQuery, *AddressGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AddressGroupBy) sqlScan(ctx context.Context, root *AddressQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AddressSelect is the builder for selecting fields of Address entities.
type AddressSelect struct {
	*AddressQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AddressSelect) Aggregate(fns ...AggregateFunc) *AddressSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AddressSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AddressQuery, *AddressSelect](ctx, as.AddressQuery, as, as.inters, v)
}

func (as *AddressSelect) sqlScan(ctx context.Context, root *AddressQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}