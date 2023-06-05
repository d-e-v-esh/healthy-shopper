// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/orderline"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/product"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/shoppingcartitem"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductItemUpdate is the builder for updating ProductItem entities.
type ProductItemUpdate struct {
	config
	hooks    []Hook
	mutation *ProductItemMutation
}

// Where appends a list predicates to the ProductItemUpdate builder.
func (piu *ProductItemUpdate) Where(ps ...predicate.ProductItem) *ProductItemUpdate {
	piu.mutation.Where(ps...)
	return piu
}

// SetProductID sets the "product_id" field.
func (piu *ProductItemUpdate) SetProductID(i int) *ProductItemUpdate {
	piu.mutation.SetProductID(i)
	return piu
}

// SetStockKeepingUnit sets the "stock_keeping_unit" field.
func (piu *ProductItemUpdate) SetStockKeepingUnit(s string) *ProductItemUpdate {
	piu.mutation.SetStockKeepingUnit(s)
	return piu
}

// SetQuantityInStock sets the "quantity_in_stock" field.
func (piu *ProductItemUpdate) SetQuantityInStock(i int) *ProductItemUpdate {
	piu.mutation.ResetQuantityInStock()
	piu.mutation.SetQuantityInStock(i)
	return piu
}

// AddQuantityInStock adds i to the "quantity_in_stock" field.
func (piu *ProductItemUpdate) AddQuantityInStock(i int) *ProductItemUpdate {
	piu.mutation.AddQuantityInStock(i)
	return piu
}

// SetProductImage sets the "product_image" field.
func (piu *ProductItemUpdate) SetProductImage(s string) *ProductItemUpdate {
	piu.mutation.SetProductImage(s)
	return piu
}

// SetPrice sets the "price" field.
func (piu *ProductItemUpdate) SetPrice(f float32) *ProductItemUpdate {
	piu.mutation.ResetPrice()
	piu.mutation.SetPrice(f)
	return piu
}

// AddPrice adds f to the "price" field.
func (piu *ProductItemUpdate) AddPrice(f float32) *ProductItemUpdate {
	piu.mutation.AddPrice(f)
	return piu
}

// SetCreatedAt sets the "created_at" field.
func (piu *ProductItemUpdate) SetCreatedAt(t time.Time) *ProductItemUpdate {
	piu.mutation.SetCreatedAt(t)
	return piu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piu *ProductItemUpdate) SetNillableCreatedAt(t *time.Time) *ProductItemUpdate {
	if t != nil {
		piu.SetCreatedAt(*t)
	}
	return piu
}

// SetUpdatedAt sets the "updated_at" field.
func (piu *ProductItemUpdate) SetUpdatedAt(t time.Time) *ProductItemUpdate {
	piu.mutation.SetUpdatedAt(t)
	return piu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piu *ProductItemUpdate) SetNillableUpdatedAt(t *time.Time) *ProductItemUpdate {
	if t != nil {
		piu.SetUpdatedAt(*t)
	}
	return piu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (piu *ProductItemUpdate) ClearUpdatedAt() *ProductItemUpdate {
	piu.mutation.ClearUpdatedAt()
	return piu
}

// SetProduct sets the "product" edge to the Product entity.
func (piu *ProductItemUpdate) SetProduct(p *Product) *ProductItemUpdate {
	return piu.SetProductID(p.ID)
}

// AddOrderLineIDs adds the "order_line" edge to the OrderLine entity by IDs.
func (piu *ProductItemUpdate) AddOrderLineIDs(ids ...int) *ProductItemUpdate {
	piu.mutation.AddOrderLineIDs(ids...)
	return piu
}

// AddOrderLine adds the "order_line" edges to the OrderLine entity.
func (piu *ProductItemUpdate) AddOrderLine(o ...*OrderLine) *ProductItemUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return piu.AddOrderLineIDs(ids...)
}

// AddShoppingCartItemIDs adds the "shopping_cart_item" edge to the ShoppingCartItem entity by IDs.
func (piu *ProductItemUpdate) AddShoppingCartItemIDs(ids ...int) *ProductItemUpdate {
	piu.mutation.AddShoppingCartItemIDs(ids...)
	return piu
}

// AddShoppingCartItem adds the "shopping_cart_item" edges to the ShoppingCartItem entity.
func (piu *ProductItemUpdate) AddShoppingCartItem(s ...*ShoppingCartItem) *ProductItemUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return piu.AddShoppingCartItemIDs(ids...)
}

// Mutation returns the ProductItemMutation object of the builder.
func (piu *ProductItemUpdate) Mutation() *ProductItemMutation {
	return piu.mutation
}

// ClearProduct clears the "product" edge to the Product entity.
func (piu *ProductItemUpdate) ClearProduct() *ProductItemUpdate {
	piu.mutation.ClearProduct()
	return piu
}

// ClearOrderLine clears all "order_line" edges to the OrderLine entity.
func (piu *ProductItemUpdate) ClearOrderLine() *ProductItemUpdate {
	piu.mutation.ClearOrderLine()
	return piu
}

// RemoveOrderLineIDs removes the "order_line" edge to OrderLine entities by IDs.
func (piu *ProductItemUpdate) RemoveOrderLineIDs(ids ...int) *ProductItemUpdate {
	piu.mutation.RemoveOrderLineIDs(ids...)
	return piu
}

// RemoveOrderLine removes "order_line" edges to OrderLine entities.
func (piu *ProductItemUpdate) RemoveOrderLine(o ...*OrderLine) *ProductItemUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return piu.RemoveOrderLineIDs(ids...)
}

// ClearShoppingCartItem clears all "shopping_cart_item" edges to the ShoppingCartItem entity.
func (piu *ProductItemUpdate) ClearShoppingCartItem() *ProductItemUpdate {
	piu.mutation.ClearShoppingCartItem()
	return piu
}

// RemoveShoppingCartItemIDs removes the "shopping_cart_item" edge to ShoppingCartItem entities by IDs.
func (piu *ProductItemUpdate) RemoveShoppingCartItemIDs(ids ...int) *ProductItemUpdate {
	piu.mutation.RemoveShoppingCartItemIDs(ids...)
	return piu
}

// RemoveShoppingCartItem removes "shopping_cart_item" edges to ShoppingCartItem entities.
func (piu *ProductItemUpdate) RemoveShoppingCartItem(s ...*ShoppingCartItem) *ProductItemUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return piu.RemoveShoppingCartItemIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (piu *ProductItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, piu.sqlSave, piu.mutation, piu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piu *ProductItemUpdate) SaveX(ctx context.Context) int {
	affected, err := piu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (piu *ProductItemUpdate) Exec(ctx context.Context) error {
	_, err := piu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piu *ProductItemUpdate) ExecX(ctx context.Context) {
	if err := piu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piu *ProductItemUpdate) check() error {
	if v, ok := piu.mutation.StockKeepingUnit(); ok {
		if err := productitem.StockKeepingUnitValidator(v); err != nil {
			return &ValidationError{Name: "stock_keeping_unit", err: fmt.Errorf(`ent: validator failed for field "ProductItem.stock_keeping_unit": %w`, err)}
		}
	}
	if v, ok := piu.mutation.QuantityInStock(); ok {
		if err := productitem.QuantityInStockValidator(v); err != nil {
			return &ValidationError{Name: "quantity_in_stock", err: fmt.Errorf(`ent: validator failed for field "ProductItem.quantity_in_stock": %w`, err)}
		}
	}
	if v, ok := piu.mutation.ProductImage(); ok {
		if err := productitem.ProductImageValidator(v); err != nil {
			return &ValidationError{Name: "product_image", err: fmt.Errorf(`ent: validator failed for field "ProductItem.product_image": %w`, err)}
		}
	}
	if v, ok := piu.mutation.Price(); ok {
		if err := productitem.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "ProductItem.price": %w`, err)}
		}
	}
	if _, ok := piu.mutation.ProductID(); piu.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductItem.product"`)
	}
	return nil
}

func (piu *ProductItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := piu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(productitem.Table, productitem.Columns, sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeInt))
	if ps := piu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piu.mutation.StockKeepingUnit(); ok {
		_spec.SetField(productitem.FieldStockKeepingUnit, field.TypeString, value)
	}
	if value, ok := piu.mutation.QuantityInStock(); ok {
		_spec.SetField(productitem.FieldQuantityInStock, field.TypeInt, value)
	}
	if value, ok := piu.mutation.AddedQuantityInStock(); ok {
		_spec.AddField(productitem.FieldQuantityInStock, field.TypeInt, value)
	}
	if value, ok := piu.mutation.ProductImage(); ok {
		_spec.SetField(productitem.FieldProductImage, field.TypeString, value)
	}
	if value, ok := piu.mutation.Price(); ok {
		_spec.SetField(productitem.FieldPrice, field.TypeFloat32, value)
	}
	if value, ok := piu.mutation.AddedPrice(); ok {
		_spec.AddField(productitem.FieldPrice, field.TypeFloat32, value)
	}
	if value, ok := piu.mutation.CreatedAt(); ok {
		_spec.SetField(productitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := piu.mutation.UpdatedAt(); ok {
		_spec.SetField(productitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if piu.mutation.UpdatedAtCleared() {
		_spec.ClearField(productitem.FieldUpdatedAt, field.TypeTime)
	}
	if piu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   productitem.ProductTable,
			Columns: []string{productitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   productitem.ProductTable,
			Columns: []string{productitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piu.mutation.OrderLineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.OrderLineTable,
			Columns: []string{productitem.OrderLineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderline.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.RemovedOrderLineIDs(); len(nodes) > 0 && !piu.mutation.OrderLineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.OrderLineTable,
			Columns: []string{productitem.OrderLineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderline.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.OrderLineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.OrderLineTable,
			Columns: []string{productitem.OrderLineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderline.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piu.mutation.ShoppingCartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.ShoppingCartItemTable,
			Columns: []string{productitem.ShoppingCartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shoppingcartitem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.RemovedShoppingCartItemIDs(); len(nodes) > 0 && !piu.mutation.ShoppingCartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.ShoppingCartItemTable,
			Columns: []string{productitem.ShoppingCartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shoppingcartitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piu.mutation.ShoppingCartItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.ShoppingCartItemTable,
			Columns: []string{productitem.ShoppingCartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shoppingcartitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, piu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	piu.mutation.done = true
	return n, nil
}

// ProductItemUpdateOne is the builder for updating a single ProductItem entity.
type ProductItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductItemMutation
}

// SetProductID sets the "product_id" field.
func (piuo *ProductItemUpdateOne) SetProductID(i int) *ProductItemUpdateOne {
	piuo.mutation.SetProductID(i)
	return piuo
}

// SetStockKeepingUnit sets the "stock_keeping_unit" field.
func (piuo *ProductItemUpdateOne) SetStockKeepingUnit(s string) *ProductItemUpdateOne {
	piuo.mutation.SetStockKeepingUnit(s)
	return piuo
}

// SetQuantityInStock sets the "quantity_in_stock" field.
func (piuo *ProductItemUpdateOne) SetQuantityInStock(i int) *ProductItemUpdateOne {
	piuo.mutation.ResetQuantityInStock()
	piuo.mutation.SetQuantityInStock(i)
	return piuo
}

// AddQuantityInStock adds i to the "quantity_in_stock" field.
func (piuo *ProductItemUpdateOne) AddQuantityInStock(i int) *ProductItemUpdateOne {
	piuo.mutation.AddQuantityInStock(i)
	return piuo
}

// SetProductImage sets the "product_image" field.
func (piuo *ProductItemUpdateOne) SetProductImage(s string) *ProductItemUpdateOne {
	piuo.mutation.SetProductImage(s)
	return piuo
}

// SetPrice sets the "price" field.
func (piuo *ProductItemUpdateOne) SetPrice(f float32) *ProductItemUpdateOne {
	piuo.mutation.ResetPrice()
	piuo.mutation.SetPrice(f)
	return piuo
}

// AddPrice adds f to the "price" field.
func (piuo *ProductItemUpdateOne) AddPrice(f float32) *ProductItemUpdateOne {
	piuo.mutation.AddPrice(f)
	return piuo
}

// SetCreatedAt sets the "created_at" field.
func (piuo *ProductItemUpdateOne) SetCreatedAt(t time.Time) *ProductItemUpdateOne {
	piuo.mutation.SetCreatedAt(t)
	return piuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (piuo *ProductItemUpdateOne) SetNillableCreatedAt(t *time.Time) *ProductItemUpdateOne {
	if t != nil {
		piuo.SetCreatedAt(*t)
	}
	return piuo
}

// SetUpdatedAt sets the "updated_at" field.
func (piuo *ProductItemUpdateOne) SetUpdatedAt(t time.Time) *ProductItemUpdateOne {
	piuo.mutation.SetUpdatedAt(t)
	return piuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (piuo *ProductItemUpdateOne) SetNillableUpdatedAt(t *time.Time) *ProductItemUpdateOne {
	if t != nil {
		piuo.SetUpdatedAt(*t)
	}
	return piuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (piuo *ProductItemUpdateOne) ClearUpdatedAt() *ProductItemUpdateOne {
	piuo.mutation.ClearUpdatedAt()
	return piuo
}

// SetProduct sets the "product" edge to the Product entity.
func (piuo *ProductItemUpdateOne) SetProduct(p *Product) *ProductItemUpdateOne {
	return piuo.SetProductID(p.ID)
}

// AddOrderLineIDs adds the "order_line" edge to the OrderLine entity by IDs.
func (piuo *ProductItemUpdateOne) AddOrderLineIDs(ids ...int) *ProductItemUpdateOne {
	piuo.mutation.AddOrderLineIDs(ids...)
	return piuo
}

// AddOrderLine adds the "order_line" edges to the OrderLine entity.
func (piuo *ProductItemUpdateOne) AddOrderLine(o ...*OrderLine) *ProductItemUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return piuo.AddOrderLineIDs(ids...)
}

// AddShoppingCartItemIDs adds the "shopping_cart_item" edge to the ShoppingCartItem entity by IDs.
func (piuo *ProductItemUpdateOne) AddShoppingCartItemIDs(ids ...int) *ProductItemUpdateOne {
	piuo.mutation.AddShoppingCartItemIDs(ids...)
	return piuo
}

// AddShoppingCartItem adds the "shopping_cart_item" edges to the ShoppingCartItem entity.
func (piuo *ProductItemUpdateOne) AddShoppingCartItem(s ...*ShoppingCartItem) *ProductItemUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return piuo.AddShoppingCartItemIDs(ids...)
}

// Mutation returns the ProductItemMutation object of the builder.
func (piuo *ProductItemUpdateOne) Mutation() *ProductItemMutation {
	return piuo.mutation
}

// ClearProduct clears the "product" edge to the Product entity.
func (piuo *ProductItemUpdateOne) ClearProduct() *ProductItemUpdateOne {
	piuo.mutation.ClearProduct()
	return piuo
}

// ClearOrderLine clears all "order_line" edges to the OrderLine entity.
func (piuo *ProductItemUpdateOne) ClearOrderLine() *ProductItemUpdateOne {
	piuo.mutation.ClearOrderLine()
	return piuo
}

// RemoveOrderLineIDs removes the "order_line" edge to OrderLine entities by IDs.
func (piuo *ProductItemUpdateOne) RemoveOrderLineIDs(ids ...int) *ProductItemUpdateOne {
	piuo.mutation.RemoveOrderLineIDs(ids...)
	return piuo
}

// RemoveOrderLine removes "order_line" edges to OrderLine entities.
func (piuo *ProductItemUpdateOne) RemoveOrderLine(o ...*OrderLine) *ProductItemUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return piuo.RemoveOrderLineIDs(ids...)
}

// ClearShoppingCartItem clears all "shopping_cart_item" edges to the ShoppingCartItem entity.
func (piuo *ProductItemUpdateOne) ClearShoppingCartItem() *ProductItemUpdateOne {
	piuo.mutation.ClearShoppingCartItem()
	return piuo
}

// RemoveShoppingCartItemIDs removes the "shopping_cart_item" edge to ShoppingCartItem entities by IDs.
func (piuo *ProductItemUpdateOne) RemoveShoppingCartItemIDs(ids ...int) *ProductItemUpdateOne {
	piuo.mutation.RemoveShoppingCartItemIDs(ids...)
	return piuo
}

// RemoveShoppingCartItem removes "shopping_cart_item" edges to ShoppingCartItem entities.
func (piuo *ProductItemUpdateOne) RemoveShoppingCartItem(s ...*ShoppingCartItem) *ProductItemUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return piuo.RemoveShoppingCartItemIDs(ids...)
}

// Where appends a list predicates to the ProductItemUpdate builder.
func (piuo *ProductItemUpdateOne) Where(ps ...predicate.ProductItem) *ProductItemUpdateOne {
	piuo.mutation.Where(ps...)
	return piuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (piuo *ProductItemUpdateOne) Select(field string, fields ...string) *ProductItemUpdateOne {
	piuo.fields = append([]string{field}, fields...)
	return piuo
}

// Save executes the query and returns the updated ProductItem entity.
func (piuo *ProductItemUpdateOne) Save(ctx context.Context) (*ProductItem, error) {
	return withHooks(ctx, piuo.sqlSave, piuo.mutation, piuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (piuo *ProductItemUpdateOne) SaveX(ctx context.Context) *ProductItem {
	node, err := piuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (piuo *ProductItemUpdateOne) Exec(ctx context.Context) error {
	_, err := piuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (piuo *ProductItemUpdateOne) ExecX(ctx context.Context) {
	if err := piuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (piuo *ProductItemUpdateOne) check() error {
	if v, ok := piuo.mutation.StockKeepingUnit(); ok {
		if err := productitem.StockKeepingUnitValidator(v); err != nil {
			return &ValidationError{Name: "stock_keeping_unit", err: fmt.Errorf(`ent: validator failed for field "ProductItem.stock_keeping_unit": %w`, err)}
		}
	}
	if v, ok := piuo.mutation.QuantityInStock(); ok {
		if err := productitem.QuantityInStockValidator(v); err != nil {
			return &ValidationError{Name: "quantity_in_stock", err: fmt.Errorf(`ent: validator failed for field "ProductItem.quantity_in_stock": %w`, err)}
		}
	}
	if v, ok := piuo.mutation.ProductImage(); ok {
		if err := productitem.ProductImageValidator(v); err != nil {
			return &ValidationError{Name: "product_image", err: fmt.Errorf(`ent: validator failed for field "ProductItem.product_image": %w`, err)}
		}
	}
	if v, ok := piuo.mutation.Price(); ok {
		if err := productitem.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "ProductItem.price": %w`, err)}
		}
	}
	if _, ok := piuo.mutation.ProductID(); piuo.mutation.ProductCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ProductItem.product"`)
	}
	return nil
}

func (piuo *ProductItemUpdateOne) sqlSave(ctx context.Context) (_node *ProductItem, err error) {
	if err := piuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(productitem.Table, productitem.Columns, sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeInt))
	id, ok := piuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := piuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productitem.FieldID)
		for _, f := range fields {
			if !productitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := piuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := piuo.mutation.StockKeepingUnit(); ok {
		_spec.SetField(productitem.FieldStockKeepingUnit, field.TypeString, value)
	}
	if value, ok := piuo.mutation.QuantityInStock(); ok {
		_spec.SetField(productitem.FieldQuantityInStock, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.AddedQuantityInStock(); ok {
		_spec.AddField(productitem.FieldQuantityInStock, field.TypeInt, value)
	}
	if value, ok := piuo.mutation.ProductImage(); ok {
		_spec.SetField(productitem.FieldProductImage, field.TypeString, value)
	}
	if value, ok := piuo.mutation.Price(); ok {
		_spec.SetField(productitem.FieldPrice, field.TypeFloat32, value)
	}
	if value, ok := piuo.mutation.AddedPrice(); ok {
		_spec.AddField(productitem.FieldPrice, field.TypeFloat32, value)
	}
	if value, ok := piuo.mutation.CreatedAt(); ok {
		_spec.SetField(productitem.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := piuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if piuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(productitem.FieldUpdatedAt, field.TypeTime)
	}
	if piuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   productitem.ProductTable,
			Columns: []string{productitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   productitem.ProductTable,
			Columns: []string{productitem.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piuo.mutation.OrderLineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.OrderLineTable,
			Columns: []string{productitem.OrderLineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderline.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.RemovedOrderLineIDs(); len(nodes) > 0 && !piuo.mutation.OrderLineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.OrderLineTable,
			Columns: []string{productitem.OrderLineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderline.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.OrderLineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.OrderLineTable,
			Columns: []string{productitem.OrderLineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(orderline.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if piuo.mutation.ShoppingCartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.ShoppingCartItemTable,
			Columns: []string{productitem.ShoppingCartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shoppingcartitem.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.RemovedShoppingCartItemIDs(); len(nodes) > 0 && !piuo.mutation.ShoppingCartItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.ShoppingCartItemTable,
			Columns: []string{productitem.ShoppingCartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shoppingcartitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := piuo.mutation.ShoppingCartItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   productitem.ShoppingCartItemTable,
			Columns: []string{productitem.ShoppingCartItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shoppingcartitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProductItem{config: piuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, piuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	piuo.mutation.done = true
	return _node, nil
}
