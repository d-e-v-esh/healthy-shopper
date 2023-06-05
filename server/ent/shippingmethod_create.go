// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/shippingmethod"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShippingMethodCreate is the builder for creating a ShippingMethod entity.
type ShippingMethodCreate struct {
	config
	mutation *ShippingMethodMutation
	hooks    []Hook
}

// SetShippingMethod sets the "shipping_method" field.
func (smc *ShippingMethodCreate) SetShippingMethod(s string) *ShippingMethodCreate {
	smc.mutation.SetShippingMethod(s)
	return smc
}

// SetShippingCost sets the "shipping_cost" field.
func (smc *ShippingMethodCreate) SetShippingCost(f float64) *ShippingMethodCreate {
	smc.mutation.SetShippingCost(f)
	return smc
}

// Mutation returns the ShippingMethodMutation object of the builder.
func (smc *ShippingMethodCreate) Mutation() *ShippingMethodMutation {
	return smc.mutation
}

// Save creates the ShippingMethod in the database.
func (smc *ShippingMethodCreate) Save(ctx context.Context) (*ShippingMethod, error) {
	return withHooks(ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *ShippingMethodCreate) SaveX(ctx context.Context) *ShippingMethod {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *ShippingMethodCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *ShippingMethodCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *ShippingMethodCreate) check() error {
	if _, ok := smc.mutation.ShippingMethod(); !ok {
		return &ValidationError{Name: "shipping_method", err: errors.New(`ent: missing required field "ShippingMethod.shipping_method"`)}
	}
	if v, ok := smc.mutation.ShippingMethod(); ok {
		if err := shippingmethod.ShippingMethodValidator(v); err != nil {
			return &ValidationError{Name: "shipping_method", err: fmt.Errorf(`ent: validator failed for field "ShippingMethod.shipping_method": %w`, err)}
		}
	}
	if _, ok := smc.mutation.ShippingCost(); !ok {
		return &ValidationError{Name: "shipping_cost", err: errors.New(`ent: missing required field "ShippingMethod.shipping_cost"`)}
	}
	if v, ok := smc.mutation.ShippingCost(); ok {
		if err := shippingmethod.ShippingCostValidator(v); err != nil {
			return &ValidationError{Name: "shipping_cost", err: fmt.Errorf(`ent: validator failed for field "ShippingMethod.shipping_cost": %w`, err)}
		}
	}
	return nil
}

func (smc *ShippingMethodCreate) sqlSave(ctx context.Context) (*ShippingMethod, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *ShippingMethodCreate) createSpec() (*ShippingMethod, *sqlgraph.CreateSpec) {
	var (
		_node = &ShippingMethod{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(shippingmethod.Table, sqlgraph.NewFieldSpec(shippingmethod.FieldID, field.TypeInt))
	)
	if value, ok := smc.mutation.ShippingMethod(); ok {
		_spec.SetField(shippingmethod.FieldShippingMethod, field.TypeString, value)
		_node.ShippingMethod = value
	}
	if value, ok := smc.mutation.ShippingCost(); ok {
		_spec.SetField(shippingmethod.FieldShippingCost, field.TypeFloat64, value)
		_node.ShippingCost = value
	}
	return _node, _spec
}

// ShippingMethodCreateBulk is the builder for creating many ShippingMethod entities in bulk.
type ShippingMethodCreateBulk struct {
	config
	builders []*ShippingMethodCreate
}

// Save creates the ShippingMethod entities in the database.
func (smcb *ShippingMethodCreateBulk) Save(ctx context.Context) ([]*ShippingMethod, error) {
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*ShippingMethod, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShippingMethodMutation)
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
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *ShippingMethodCreateBulk) SaveX(ctx context.Context) []*ShippingMethod {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *ShippingMethodCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *ShippingMethodCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
