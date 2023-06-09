// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/orderline"
	"healthyshopper/ent/productitem"
	"healthyshopper/ent/userreview"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderLineCreate is the builder for creating a OrderLine entity.
type OrderLineCreate struct {
	config
	mutation *OrderLineMutation
	hooks    []Hook
}

// SetProductItemID sets the "product_item_id" field.
func (olc *OrderLineCreate) SetProductItemID(i int) *OrderLineCreate {
	olc.mutation.SetProductItemID(i)
	return olc
}

// SetShopOrderID sets the "shop_order_id" field.
func (olc *OrderLineCreate) SetShopOrderID(i int) *OrderLineCreate {
	olc.mutation.SetShopOrderID(i)
	return olc
}

// SetQuantity sets the "quantity" field.
func (olc *OrderLineCreate) SetQuantity(i int) *OrderLineCreate {
	olc.mutation.SetQuantity(i)
	return olc
}

// SetPrice sets the "price" field.
func (olc *OrderLineCreate) SetPrice(f float64) *OrderLineCreate {
	olc.mutation.SetPrice(f)
	return olc
}

// SetProductItem sets the "product_item" edge to the ProductItem entity.
func (olc *OrderLineCreate) SetProductItem(p *ProductItem) *OrderLineCreate {
	return olc.SetProductItemID(p.ID)
}

// AddUserReviewIDs adds the "user_review" edge to the UserReview entity by IDs.
func (olc *OrderLineCreate) AddUserReviewIDs(ids ...int) *OrderLineCreate {
	olc.mutation.AddUserReviewIDs(ids...)
	return olc
}

// AddUserReview adds the "user_review" edges to the UserReview entity.
func (olc *OrderLineCreate) AddUserReview(u ...*UserReview) *OrderLineCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return olc.AddUserReviewIDs(ids...)
}

// Mutation returns the OrderLineMutation object of the builder.
func (olc *OrderLineCreate) Mutation() *OrderLineMutation {
	return olc.mutation
}

// Save creates the OrderLine in the database.
func (olc *OrderLineCreate) Save(ctx context.Context) (*OrderLine, error) {
	return withHooks(ctx, olc.sqlSave, olc.mutation, olc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (olc *OrderLineCreate) SaveX(ctx context.Context) *OrderLine {
	v, err := olc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olc *OrderLineCreate) Exec(ctx context.Context) error {
	_, err := olc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olc *OrderLineCreate) ExecX(ctx context.Context) {
	if err := olc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (olc *OrderLineCreate) check() error {
	if _, ok := olc.mutation.ProductItemID(); !ok {
		return &ValidationError{Name: "product_item_id", err: errors.New(`ent: missing required field "OrderLine.product_item_id"`)}
	}
	if _, ok := olc.mutation.ShopOrderID(); !ok {
		return &ValidationError{Name: "shop_order_id", err: errors.New(`ent: missing required field "OrderLine.shop_order_id"`)}
	}
	if _, ok := olc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "OrderLine.quantity"`)}
	}
	if v, ok := olc.mutation.Quantity(); ok {
		if err := orderline.QuantityValidator(v); err != nil {
			return &ValidationError{Name: "quantity", err: fmt.Errorf(`ent: validator failed for field "OrderLine.quantity": %w`, err)}
		}
	}
	if _, ok := olc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "OrderLine.price"`)}
	}
	if v, ok := olc.mutation.Price(); ok {
		if err := orderline.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "OrderLine.price": %w`, err)}
		}
	}
	if _, ok := olc.mutation.ProductItemID(); !ok {
		return &ValidationError{Name: "product_item", err: errors.New(`ent: missing required edge "OrderLine.product_item"`)}
	}
	return nil
}

func (olc *OrderLineCreate) sqlSave(ctx context.Context) (*OrderLine, error) {
	if err := olc.check(); err != nil {
		return nil, err
	}
	_node, _spec := olc.createSpec()
	if err := sqlgraph.CreateNode(ctx, olc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	olc.mutation.id = &_node.ID
	olc.mutation.done = true
	return _node, nil
}

func (olc *OrderLineCreate) createSpec() (*OrderLine, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderLine{config: olc.config}
		_spec = sqlgraph.NewCreateSpec(orderline.Table, sqlgraph.NewFieldSpec(orderline.FieldID, field.TypeInt))
	)
	if value, ok := olc.mutation.ShopOrderID(); ok {
		_spec.SetField(orderline.FieldShopOrderID, field.TypeInt, value)
		_node.ShopOrderID = value
	}
	if value, ok := olc.mutation.Quantity(); ok {
		_spec.SetField(orderline.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := olc.mutation.Price(); ok {
		_spec.SetField(orderline.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if nodes := olc.mutation.ProductItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderline.ProductItemTable,
			Columns: []string{orderline.ProductItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productitem.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductItemID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := olc.mutation.UserReviewIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   orderline.UserReviewTable,
			Columns: []string{orderline.UserReviewColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(userreview.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderLineCreateBulk is the builder for creating many OrderLine entities in bulk.
type OrderLineCreateBulk struct {
	config
	builders []*OrderLineCreate
}

// Save creates the OrderLine entities in the database.
func (olcb *OrderLineCreateBulk) Save(ctx context.Context) ([]*OrderLine, error) {
	specs := make([]*sqlgraph.CreateSpec, len(olcb.builders))
	nodes := make([]*OrderLine, len(olcb.builders))
	mutators := make([]Mutator, len(olcb.builders))
	for i := range olcb.builders {
		func(i int, root context.Context) {
			builder := olcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderLineMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, olcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, olcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, olcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (olcb *OrderLineCreateBulk) SaveX(ctx context.Context) []*OrderLine {
	v, err := olcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (olcb *OrderLineCreateBulk) Exec(ctx context.Context) error {
	_, err := olcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (olcb *OrderLineCreateBulk) ExecX(ctx context.Context) {
	if err := olcb.Exec(ctx); err != nil {
		panic(err)
	}
}
