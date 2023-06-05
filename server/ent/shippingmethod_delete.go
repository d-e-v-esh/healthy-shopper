// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/shippingmethod"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShippingMethodDelete is the builder for deleting a ShippingMethod entity.
type ShippingMethodDelete struct {
	config
	hooks    []Hook
	mutation *ShippingMethodMutation
}

// Where appends a list predicates to the ShippingMethodDelete builder.
func (smd *ShippingMethodDelete) Where(ps ...predicate.ShippingMethod) *ShippingMethodDelete {
	smd.mutation.Where(ps...)
	return smd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (smd *ShippingMethodDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, smd.sqlExec, smd.mutation, smd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (smd *ShippingMethodDelete) ExecX(ctx context.Context) int {
	n, err := smd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (smd *ShippingMethodDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(shippingmethod.Table, sqlgraph.NewFieldSpec(shippingmethod.FieldID, field.TypeInt))
	if ps := smd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, smd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	smd.mutation.done = true
	return affected, err
}

// ShippingMethodDeleteOne is the builder for deleting a single ShippingMethod entity.
type ShippingMethodDeleteOne struct {
	smd *ShippingMethodDelete
}

// Where appends a list predicates to the ShippingMethodDelete builder.
func (smdo *ShippingMethodDeleteOne) Where(ps ...predicate.ShippingMethod) *ShippingMethodDeleteOne {
	smdo.smd.mutation.Where(ps...)
	return smdo
}

// Exec executes the deletion query.
func (smdo *ShippingMethodDeleteOne) Exec(ctx context.Context) error {
	n, err := smdo.smd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shippingmethod.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (smdo *ShippingMethodDeleteOne) ExecX(ctx context.Context) {
	if err := smdo.Exec(ctx); err != nil {
		panic(err)
	}
}
