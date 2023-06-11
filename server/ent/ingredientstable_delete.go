// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthyshopper/ent/ingredientstable"
	"healthyshopper/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IngredientsTableDelete is the builder for deleting a IngredientsTable entity.
type IngredientsTableDelete struct {
	config
	hooks    []Hook
	mutation *IngredientsTableMutation
}

// Where appends a list predicates to the IngredientsTableDelete builder.
func (itd *IngredientsTableDelete) Where(ps ...predicate.IngredientsTable) *IngredientsTableDelete {
	itd.mutation.Where(ps...)
	return itd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (itd *IngredientsTableDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, itd.sqlExec, itd.mutation, itd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (itd *IngredientsTableDelete) ExecX(ctx context.Context) int {
	n, err := itd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (itd *IngredientsTableDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(ingredientstable.Table, sqlgraph.NewFieldSpec(ingredientstable.FieldID, field.TypeInt))
	if ps := itd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, itd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	itd.mutation.done = true
	return affected, err
}

// IngredientsTableDeleteOne is the builder for deleting a single IngredientsTable entity.
type IngredientsTableDeleteOne struct {
	itd *IngredientsTableDelete
}

// Where appends a list predicates to the IngredientsTableDelete builder.
func (itdo *IngredientsTableDeleteOne) Where(ps ...predicate.IngredientsTable) *IngredientsTableDeleteOne {
	itdo.itd.mutation.Where(ps...)
	return itdo
}

// Exec executes the deletion query.
func (itdo *IngredientsTableDeleteOne) Exec(ctx context.Context) error {
	n, err := itdo.itd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{ingredientstable.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (itdo *IngredientsTableDeleteOne) ExecX(ctx context.Context) {
	if err := itdo.Exec(ctx); err != nil {
		panic(err)
	}
}