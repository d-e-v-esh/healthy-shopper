// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/shippingaddress"
	"healthyshopper/ent/shoporder"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShopOrderUpdate is the builder for updating ShopOrder entities.
type ShopOrderUpdate struct {
	config
	hooks    []Hook
	mutation *ShopOrderMutation
}

// Where appends a list predicates to the ShopOrderUpdate builder.
func (sou *ShopOrderUpdate) Where(ps ...predicate.ShopOrder) *ShopOrderUpdate {
	sou.mutation.Where(ps...)
	return sou
}

// SetOrderDateAndTime sets the "order_date_and_time" field.
func (sou *ShopOrderUpdate) SetOrderDateAndTime(t time.Time) *ShopOrderUpdate {
	sou.mutation.SetOrderDateAndTime(t)
	return sou
}

// SetNillableOrderDateAndTime sets the "order_date_and_time" field if the given value is not nil.
func (sou *ShopOrderUpdate) SetNillableOrderDateAndTime(t *time.Time) *ShopOrderUpdate {
	if t != nil {
		sou.SetOrderDateAndTime(*t)
	}
	return sou
}

// SetPaymentMethod sets the "payment_method" field.
func (sou *ShopOrderUpdate) SetPaymentMethod(s string) *ShopOrderUpdate {
	sou.mutation.SetPaymentMethod(s)
	return sou
}

// SetTotalPrice sets the "total_price" field.
func (sou *ShopOrderUpdate) SetTotalPrice(f float64) *ShopOrderUpdate {
	sou.mutation.ResetTotalPrice()
	sou.mutation.SetTotalPrice(f)
	return sou
}

// AddTotalPrice adds f to the "total_price" field.
func (sou *ShopOrderUpdate) AddTotalPrice(f float64) *ShopOrderUpdate {
	sou.mutation.AddTotalPrice(f)
	return sou
}

// SetUserID sets the "user_id" field.
func (sou *ShopOrderUpdate) SetUserID(i int) *ShopOrderUpdate {
	sou.mutation.ResetUserID()
	sou.mutation.SetUserID(i)
	return sou
}

// AddUserID adds i to the "user_id" field.
func (sou *ShopOrderUpdate) AddUserID(i int) *ShopOrderUpdate {
	sou.mutation.AddUserID(i)
	return sou
}

// SetShippingAddressID sets the "shipping_address_id" field.
func (sou *ShopOrderUpdate) SetShippingAddressID(i int) *ShopOrderUpdate {
	sou.mutation.SetShippingAddressID(i)
	return sou
}

// SetShippingMethodID sets the "shipping_method_id" field.
func (sou *ShopOrderUpdate) SetShippingMethodID(i int) *ShopOrderUpdate {
	sou.mutation.ResetShippingMethodID()
	sou.mutation.SetShippingMethodID(i)
	return sou
}

// AddShippingMethodID adds i to the "shipping_method_id" field.
func (sou *ShopOrderUpdate) AddShippingMethodID(i int) *ShopOrderUpdate {
	sou.mutation.AddShippingMethodID(i)
	return sou
}

// SetOrderStatusID sets the "order_status_id" field.
func (sou *ShopOrderUpdate) SetOrderStatusID(i int) *ShopOrderUpdate {
	sou.mutation.ResetOrderStatusID()
	sou.mutation.SetOrderStatusID(i)
	return sou
}

// AddOrderStatusID adds i to the "order_status_id" field.
func (sou *ShopOrderUpdate) AddOrderStatusID(i int) *ShopOrderUpdate {
	sou.mutation.AddOrderStatusID(i)
	return sou
}

// SetShippingAddress sets the "shipping_address" edge to the ShippingAddress entity.
func (sou *ShopOrderUpdate) SetShippingAddress(s *ShippingAddress) *ShopOrderUpdate {
	return sou.SetShippingAddressID(s.ID)
}

// Mutation returns the ShopOrderMutation object of the builder.
func (sou *ShopOrderUpdate) Mutation() *ShopOrderMutation {
	return sou.mutation
}

// ClearShippingAddress clears the "shipping_address" edge to the ShippingAddress entity.
func (sou *ShopOrderUpdate) ClearShippingAddress() *ShopOrderUpdate {
	sou.mutation.ClearShippingAddress()
	return sou
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sou *ShopOrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, sou.sqlSave, sou.mutation, sou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sou *ShopOrderUpdate) SaveX(ctx context.Context) int {
	affected, err := sou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sou *ShopOrderUpdate) Exec(ctx context.Context) error {
	_, err := sou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sou *ShopOrderUpdate) ExecX(ctx context.Context) {
	if err := sou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sou *ShopOrderUpdate) check() error {
	if v, ok := sou.mutation.PaymentMethod(); ok {
		if err := shoporder.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "ShopOrder.payment_method": %w`, err)}
		}
	}
	if v, ok := sou.mutation.TotalPrice(); ok {
		if err := shoporder.TotalPriceValidator(v); err != nil {
			return &ValidationError{Name: "total_price", err: fmt.Errorf(`ent: validator failed for field "ShopOrder.total_price": %w`, err)}
		}
	}
	if _, ok := sou.mutation.ShippingAddressID(); sou.mutation.ShippingAddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShopOrder.shipping_address"`)
	}
	return nil
}

func (sou *ShopOrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := sou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(shoporder.Table, shoporder.Columns, sqlgraph.NewFieldSpec(shoporder.FieldID, field.TypeInt))
	if ps := sou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sou.mutation.OrderDateAndTime(); ok {
		_spec.SetField(shoporder.FieldOrderDateAndTime, field.TypeTime, value)
	}
	if value, ok := sou.mutation.PaymentMethod(); ok {
		_spec.SetField(shoporder.FieldPaymentMethod, field.TypeString, value)
	}
	if value, ok := sou.mutation.TotalPrice(); ok {
		_spec.SetField(shoporder.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := sou.mutation.AddedTotalPrice(); ok {
		_spec.AddField(shoporder.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := sou.mutation.UserID(); ok {
		_spec.SetField(shoporder.FieldUserID, field.TypeInt, value)
	}
	if value, ok := sou.mutation.AddedUserID(); ok {
		_spec.AddField(shoporder.FieldUserID, field.TypeInt, value)
	}
	if value, ok := sou.mutation.ShippingMethodID(); ok {
		_spec.SetField(shoporder.FieldShippingMethodID, field.TypeInt, value)
	}
	if value, ok := sou.mutation.AddedShippingMethodID(); ok {
		_spec.AddField(shoporder.FieldShippingMethodID, field.TypeInt, value)
	}
	if value, ok := sou.mutation.OrderStatusID(); ok {
		_spec.SetField(shoporder.FieldOrderStatusID, field.TypeInt, value)
	}
	if value, ok := sou.mutation.AddedOrderStatusID(); ok {
		_spec.AddField(shoporder.FieldOrderStatusID, field.TypeInt, value)
	}
	if sou.mutation.ShippingAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shoporder.ShippingAddressTable,
			Columns: []string{shoporder.ShippingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shippingaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sou.mutation.ShippingAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shoporder.ShippingAddressTable,
			Columns: []string{shoporder.ShippingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shippingaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shoporder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sou.mutation.done = true
	return n, nil
}

// ShopOrderUpdateOne is the builder for updating a single ShopOrder entity.
type ShopOrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShopOrderMutation
}

// SetOrderDateAndTime sets the "order_date_and_time" field.
func (souo *ShopOrderUpdateOne) SetOrderDateAndTime(t time.Time) *ShopOrderUpdateOne {
	souo.mutation.SetOrderDateAndTime(t)
	return souo
}

// SetNillableOrderDateAndTime sets the "order_date_and_time" field if the given value is not nil.
func (souo *ShopOrderUpdateOne) SetNillableOrderDateAndTime(t *time.Time) *ShopOrderUpdateOne {
	if t != nil {
		souo.SetOrderDateAndTime(*t)
	}
	return souo
}

// SetPaymentMethod sets the "payment_method" field.
func (souo *ShopOrderUpdateOne) SetPaymentMethod(s string) *ShopOrderUpdateOne {
	souo.mutation.SetPaymentMethod(s)
	return souo
}

// SetTotalPrice sets the "total_price" field.
func (souo *ShopOrderUpdateOne) SetTotalPrice(f float64) *ShopOrderUpdateOne {
	souo.mutation.ResetTotalPrice()
	souo.mutation.SetTotalPrice(f)
	return souo
}

// AddTotalPrice adds f to the "total_price" field.
func (souo *ShopOrderUpdateOne) AddTotalPrice(f float64) *ShopOrderUpdateOne {
	souo.mutation.AddTotalPrice(f)
	return souo
}

// SetUserID sets the "user_id" field.
func (souo *ShopOrderUpdateOne) SetUserID(i int) *ShopOrderUpdateOne {
	souo.mutation.ResetUserID()
	souo.mutation.SetUserID(i)
	return souo
}

// AddUserID adds i to the "user_id" field.
func (souo *ShopOrderUpdateOne) AddUserID(i int) *ShopOrderUpdateOne {
	souo.mutation.AddUserID(i)
	return souo
}

// SetShippingAddressID sets the "shipping_address_id" field.
func (souo *ShopOrderUpdateOne) SetShippingAddressID(i int) *ShopOrderUpdateOne {
	souo.mutation.SetShippingAddressID(i)
	return souo
}

// SetShippingMethodID sets the "shipping_method_id" field.
func (souo *ShopOrderUpdateOne) SetShippingMethodID(i int) *ShopOrderUpdateOne {
	souo.mutation.ResetShippingMethodID()
	souo.mutation.SetShippingMethodID(i)
	return souo
}

// AddShippingMethodID adds i to the "shipping_method_id" field.
func (souo *ShopOrderUpdateOne) AddShippingMethodID(i int) *ShopOrderUpdateOne {
	souo.mutation.AddShippingMethodID(i)
	return souo
}

// SetOrderStatusID sets the "order_status_id" field.
func (souo *ShopOrderUpdateOne) SetOrderStatusID(i int) *ShopOrderUpdateOne {
	souo.mutation.ResetOrderStatusID()
	souo.mutation.SetOrderStatusID(i)
	return souo
}

// AddOrderStatusID adds i to the "order_status_id" field.
func (souo *ShopOrderUpdateOne) AddOrderStatusID(i int) *ShopOrderUpdateOne {
	souo.mutation.AddOrderStatusID(i)
	return souo
}

// SetShippingAddress sets the "shipping_address" edge to the ShippingAddress entity.
func (souo *ShopOrderUpdateOne) SetShippingAddress(s *ShippingAddress) *ShopOrderUpdateOne {
	return souo.SetShippingAddressID(s.ID)
}

// Mutation returns the ShopOrderMutation object of the builder.
func (souo *ShopOrderUpdateOne) Mutation() *ShopOrderMutation {
	return souo.mutation
}

// ClearShippingAddress clears the "shipping_address" edge to the ShippingAddress entity.
func (souo *ShopOrderUpdateOne) ClearShippingAddress() *ShopOrderUpdateOne {
	souo.mutation.ClearShippingAddress()
	return souo
}

// Where appends a list predicates to the ShopOrderUpdate builder.
func (souo *ShopOrderUpdateOne) Where(ps ...predicate.ShopOrder) *ShopOrderUpdateOne {
	souo.mutation.Where(ps...)
	return souo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (souo *ShopOrderUpdateOne) Select(field string, fields ...string) *ShopOrderUpdateOne {
	souo.fields = append([]string{field}, fields...)
	return souo
}

// Save executes the query and returns the updated ShopOrder entity.
func (souo *ShopOrderUpdateOne) Save(ctx context.Context) (*ShopOrder, error) {
	return withHooks(ctx, souo.sqlSave, souo.mutation, souo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (souo *ShopOrderUpdateOne) SaveX(ctx context.Context) *ShopOrder {
	node, err := souo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (souo *ShopOrderUpdateOne) Exec(ctx context.Context) error {
	_, err := souo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (souo *ShopOrderUpdateOne) ExecX(ctx context.Context) {
	if err := souo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (souo *ShopOrderUpdateOne) check() error {
	if v, ok := souo.mutation.PaymentMethod(); ok {
		if err := shoporder.PaymentMethodValidator(v); err != nil {
			return &ValidationError{Name: "payment_method", err: fmt.Errorf(`ent: validator failed for field "ShopOrder.payment_method": %w`, err)}
		}
	}
	if v, ok := souo.mutation.TotalPrice(); ok {
		if err := shoporder.TotalPriceValidator(v); err != nil {
			return &ValidationError{Name: "total_price", err: fmt.Errorf(`ent: validator failed for field "ShopOrder.total_price": %w`, err)}
		}
	}
	if _, ok := souo.mutation.ShippingAddressID(); souo.mutation.ShippingAddressCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShopOrder.shipping_address"`)
	}
	return nil
}

func (souo *ShopOrderUpdateOne) sqlSave(ctx context.Context) (_node *ShopOrder, err error) {
	if err := souo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(shoporder.Table, shoporder.Columns, sqlgraph.NewFieldSpec(shoporder.FieldID, field.TypeInt))
	id, ok := souo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ShopOrder.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := souo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, shoporder.FieldID)
		for _, f := range fields {
			if !shoporder.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != shoporder.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := souo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := souo.mutation.OrderDateAndTime(); ok {
		_spec.SetField(shoporder.FieldOrderDateAndTime, field.TypeTime, value)
	}
	if value, ok := souo.mutation.PaymentMethod(); ok {
		_spec.SetField(shoporder.FieldPaymentMethod, field.TypeString, value)
	}
	if value, ok := souo.mutation.TotalPrice(); ok {
		_spec.SetField(shoporder.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := souo.mutation.AddedTotalPrice(); ok {
		_spec.AddField(shoporder.FieldTotalPrice, field.TypeFloat64, value)
	}
	if value, ok := souo.mutation.UserID(); ok {
		_spec.SetField(shoporder.FieldUserID, field.TypeInt, value)
	}
	if value, ok := souo.mutation.AddedUserID(); ok {
		_spec.AddField(shoporder.FieldUserID, field.TypeInt, value)
	}
	if value, ok := souo.mutation.ShippingMethodID(); ok {
		_spec.SetField(shoporder.FieldShippingMethodID, field.TypeInt, value)
	}
	if value, ok := souo.mutation.AddedShippingMethodID(); ok {
		_spec.AddField(shoporder.FieldShippingMethodID, field.TypeInt, value)
	}
	if value, ok := souo.mutation.OrderStatusID(); ok {
		_spec.SetField(shoporder.FieldOrderStatusID, field.TypeInt, value)
	}
	if value, ok := souo.mutation.AddedOrderStatusID(); ok {
		_spec.AddField(shoporder.FieldOrderStatusID, field.TypeInt, value)
	}
	if souo.mutation.ShippingAddressCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shoporder.ShippingAddressTable,
			Columns: []string{shoporder.ShippingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shippingaddress.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := souo.mutation.ShippingAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   shoporder.ShippingAddressTable,
			Columns: []string{shoporder.ShippingAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(shippingaddress.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ShopOrder{config: souo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, souo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shoporder.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	souo.mutation.done = true
	return _node, nil
}
