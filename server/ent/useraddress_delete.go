// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/useraddress"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserAddressDelete is the builder for deleting a UserAddress entity.
type UserAddressDelete struct {
	config
	hooks    []Hook
	mutation *UserAddressMutation
}

// Where appends a list predicates to the UserAddressDelete builder.
func (uad *UserAddressDelete) Where(ps ...predicate.UserAddress) *UserAddressDelete {
	uad.mutation.Where(ps...)
	return uad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uad *UserAddressDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uad.sqlExec, uad.mutation, uad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uad *UserAddressDelete) ExecX(ctx context.Context) int {
	n, err := uad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uad *UserAddressDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(useraddress.Table, sqlgraph.NewFieldSpec(useraddress.FieldID, field.TypeInt))
	if ps := uad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uad.mutation.done = true
	return affected, err
}

// UserAddressDeleteOne is the builder for deleting a single UserAddress entity.
type UserAddressDeleteOne struct {
	uad *UserAddressDelete
}

// Where appends a list predicates to the UserAddressDelete builder.
func (uado *UserAddressDeleteOne) Where(ps ...predicate.UserAddress) *UserAddressDeleteOne {
	uado.uad.mutation.Where(ps...)
	return uado
}

// Exec executes the deletion query.
func (uado *UserAddressDeleteOne) Exec(ctx context.Context) error {
	n, err := uado.uad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{useraddress.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uado *UserAddressDeleteOne) ExecX(ctx context.Context) {
	if err := uado.Exec(ctx); err != nil {
		panic(err)
	}
}