// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/orderstatus"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderStatusCreate is the builder for creating a OrderStatus entity.
type OrderStatusCreate struct {
	config
	mutation *OrderStatusMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (osc *OrderStatusCreate) SetStatus(s string) *OrderStatusCreate {
	osc.mutation.SetStatus(s)
	return osc
}

// Mutation returns the OrderStatusMutation object of the builder.
func (osc *OrderStatusCreate) Mutation() *OrderStatusMutation {
	return osc.mutation
}

// Save creates the OrderStatus in the database.
func (osc *OrderStatusCreate) Save(ctx context.Context) (*OrderStatus, error) {
	return withHooks(ctx, osc.sqlSave, osc.mutation, osc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (osc *OrderStatusCreate) SaveX(ctx context.Context) *OrderStatus {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (osc *OrderStatusCreate) Exec(ctx context.Context) error {
	_, err := osc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osc *OrderStatusCreate) ExecX(ctx context.Context) {
	if err := osc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osc *OrderStatusCreate) check() error {
	if _, ok := osc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "OrderStatus.status"`)}
	}
	if v, ok := osc.mutation.Status(); ok {
		if err := orderstatus.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "OrderStatus.status": %w`, err)}
		}
	}
	return nil
}

func (osc *OrderStatusCreate) sqlSave(ctx context.Context) (*OrderStatus, error) {
	if err := osc.check(); err != nil {
		return nil, err
	}
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	osc.mutation.id = &_node.ID
	osc.mutation.done = true
	return _node, nil
}

func (osc *OrderStatusCreate) createSpec() (*OrderStatus, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderStatus{config: osc.config}
		_spec = sqlgraph.NewCreateSpec(orderstatus.Table, sqlgraph.NewFieldSpec(orderstatus.FieldID, field.TypeInt))
	)
	if value, ok := osc.mutation.Status(); ok {
		_spec.SetField(orderstatus.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	return _node, _spec
}

// OrderStatusCreateBulk is the builder for creating many OrderStatus entities in bulk.
type OrderStatusCreateBulk struct {
	config
	builders []*OrderStatusCreate
}

// Save creates the OrderStatus entities in the database.
func (oscb *OrderStatusCreateBulk) Save(ctx context.Context) ([]*OrderStatus, error) {
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OrderStatus, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderStatusMutation)
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
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OrderStatusCreateBulk) SaveX(ctx context.Context) []*OrderStatus {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oscb *OrderStatusCreateBulk) Exec(ctx context.Context) error {
	_, err := oscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscb *OrderStatusCreateBulk) ExecX(ctx context.Context) {
	if err := oscb.Exec(ctx); err != nil {
		panic(err)
	}
}
