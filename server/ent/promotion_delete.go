// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/promotion"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PromotionDelete is the builder for deleting a Promotion entity.
type PromotionDelete struct {
	config
	hooks    []Hook
	mutation *PromotionMutation
}

// Where appends a list predicates to the PromotionDelete builder.
func (pd *PromotionDelete) Where(ps ...predicate.Promotion) *PromotionDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PromotionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PromotionDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PromotionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(promotion.Table, sqlgraph.NewFieldSpec(promotion.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PromotionDeleteOne is the builder for deleting a single Promotion entity.
type PromotionDeleteOne struct {
	pd *PromotionDelete
}

// Where appends a list predicates to the PromotionDelete builder.
func (pdo *PromotionDeleteOne) Where(ps ...predicate.Promotion) *PromotionDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *PromotionDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{promotion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PromotionDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}
