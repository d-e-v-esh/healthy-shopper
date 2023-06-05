// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/user"
	"healthyshopper/ent/userreview"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserReviewUpdate is the builder for updating UserReview entities.
type UserReviewUpdate struct {
	config
	hooks    []Hook
	mutation *UserReviewMutation
}

// Where appends a list predicates to the UserReviewUpdate builder.
func (uru *UserReviewUpdate) Where(ps ...predicate.UserReview) *UserReviewUpdate {
	uru.mutation.Where(ps...)
	return uru
}

// SetUserReviewID sets the "user_review_id" field.
func (uru *UserReviewUpdate) SetUserReviewID(i int) *UserReviewUpdate {
	uru.mutation.ResetUserReviewID()
	uru.mutation.SetUserReviewID(i)
	return uru
}

// AddUserReviewID adds i to the "user_review_id" field.
func (uru *UserReviewUpdate) AddUserReviewID(i int) *UserReviewUpdate {
	uru.mutation.AddUserReviewID(i)
	return uru
}

// SetUserID sets the "user_id" field.
func (uru *UserReviewUpdate) SetUserID(i int) *UserReviewUpdate {
	uru.mutation.SetUserID(i)
	return uru
}

// SetOrderedProductID sets the "ordered_product_id" field.
func (uru *UserReviewUpdate) SetOrderedProductID(i int) *UserReviewUpdate {
	uru.mutation.ResetOrderedProductID()
	uru.mutation.SetOrderedProductID(i)
	return uru
}

// AddOrderedProductID adds i to the "ordered_product_id" field.
func (uru *UserReviewUpdate) AddOrderedProductID(i int) *UserReviewUpdate {
	uru.mutation.AddOrderedProductID(i)
	return uru
}

// SetRating sets the "rating" field.
func (uru *UserReviewUpdate) SetRating(i int) *UserReviewUpdate {
	uru.mutation.ResetRating()
	uru.mutation.SetRating(i)
	return uru
}

// AddRating adds i to the "rating" field.
func (uru *UserReviewUpdate) AddRating(i int) *UserReviewUpdate {
	uru.mutation.AddRating(i)
	return uru
}

// SetReview sets the "review" field.
func (uru *UserReviewUpdate) SetReview(s string) *UserReviewUpdate {
	uru.mutation.SetReview(s)
	return uru
}

// SetCreatedAt sets the "created_at" field.
func (uru *UserReviewUpdate) SetCreatedAt(t time.Time) *UserReviewUpdate {
	uru.mutation.SetCreatedAt(t)
	return uru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uru *UserReviewUpdate) SetNillableCreatedAt(t *time.Time) *UserReviewUpdate {
	if t != nil {
		uru.SetCreatedAt(*t)
	}
	return uru
}

// SetUpdatedAt sets the "updated_at" field.
func (uru *UserReviewUpdate) SetUpdatedAt(t time.Time) *UserReviewUpdate {
	uru.mutation.SetUpdatedAt(t)
	return uru
}

// SetUser sets the "user" edge to the User entity.
func (uru *UserReviewUpdate) SetUser(u *User) *UserReviewUpdate {
	return uru.SetUserID(u.ID)
}

// Mutation returns the UserReviewMutation object of the builder.
func (uru *UserReviewUpdate) Mutation() *UserReviewMutation {
	return uru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uru *UserReviewUpdate) ClearUser() *UserReviewUpdate {
	uru.mutation.ClearUser()
	return uru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserReviewUpdate) Save(ctx context.Context) (int, error) {
	uru.defaults()
	return withHooks(ctx, uru.sqlSave, uru.mutation, uru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserReviewUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserReviewUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uru *UserReviewUpdate) defaults() {
	if _, ok := uru.mutation.UpdatedAt(); !ok {
		v := userreview.UpdateDefaultUpdatedAt()
		uru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uru *UserReviewUpdate) check() error {
	if v, ok := uru.mutation.Rating(); ok {
		if err := userreview.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "UserReview.rating": %w`, err)}
		}
	}
	if v, ok := uru.mutation.Review(); ok {
		if err := userreview.ReviewValidator(v); err != nil {
			return &ValidationError{Name: "review", err: fmt.Errorf(`ent: validator failed for field "UserReview.review": %w`, err)}
		}
	}
	if _, ok := uru.mutation.UserID(); uru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserReview.user"`)
	}
	return nil
}

func (uru *UserReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(userreview.Table, userreview.Columns, sqlgraph.NewFieldSpec(userreview.FieldID, field.TypeInt))
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uru.mutation.UserReviewID(); ok {
		_spec.SetField(userreview.FieldUserReviewID, field.TypeInt, value)
	}
	if value, ok := uru.mutation.AddedUserReviewID(); ok {
		_spec.AddField(userreview.FieldUserReviewID, field.TypeInt, value)
	}
	if value, ok := uru.mutation.OrderedProductID(); ok {
		_spec.SetField(userreview.FieldOrderedProductID, field.TypeInt, value)
	}
	if value, ok := uru.mutation.AddedOrderedProductID(); ok {
		_spec.AddField(userreview.FieldOrderedProductID, field.TypeInt, value)
	}
	if value, ok := uru.mutation.Rating(); ok {
		_spec.SetField(userreview.FieldRating, field.TypeInt, value)
	}
	if value, ok := uru.mutation.AddedRating(); ok {
		_spec.AddField(userreview.FieldRating, field.TypeInt, value)
	}
	if value, ok := uru.mutation.Review(); ok {
		_spec.SetField(userreview.FieldReview, field.TypeString, value)
	}
	if value, ok := uru.mutation.CreatedAt(); ok {
		_spec.SetField(userreview.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uru.mutation.UpdatedAt(); ok {
		_spec.SetField(userreview.FieldUpdatedAt, field.TypeTime, value)
	}
	if uru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userreview.UserTable,
			Columns: []string{userreview.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userreview.UserTable,
			Columns: []string{userreview.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userreview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uru.mutation.done = true
	return n, nil
}

// UserReviewUpdateOne is the builder for updating a single UserReview entity.
type UserReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserReviewMutation
}

// SetUserReviewID sets the "user_review_id" field.
func (uruo *UserReviewUpdateOne) SetUserReviewID(i int) *UserReviewUpdateOne {
	uruo.mutation.ResetUserReviewID()
	uruo.mutation.SetUserReviewID(i)
	return uruo
}

// AddUserReviewID adds i to the "user_review_id" field.
func (uruo *UserReviewUpdateOne) AddUserReviewID(i int) *UserReviewUpdateOne {
	uruo.mutation.AddUserReviewID(i)
	return uruo
}

// SetUserID sets the "user_id" field.
func (uruo *UserReviewUpdateOne) SetUserID(i int) *UserReviewUpdateOne {
	uruo.mutation.SetUserID(i)
	return uruo
}

// SetOrderedProductID sets the "ordered_product_id" field.
func (uruo *UserReviewUpdateOne) SetOrderedProductID(i int) *UserReviewUpdateOne {
	uruo.mutation.ResetOrderedProductID()
	uruo.mutation.SetOrderedProductID(i)
	return uruo
}

// AddOrderedProductID adds i to the "ordered_product_id" field.
func (uruo *UserReviewUpdateOne) AddOrderedProductID(i int) *UserReviewUpdateOne {
	uruo.mutation.AddOrderedProductID(i)
	return uruo
}

// SetRating sets the "rating" field.
func (uruo *UserReviewUpdateOne) SetRating(i int) *UserReviewUpdateOne {
	uruo.mutation.ResetRating()
	uruo.mutation.SetRating(i)
	return uruo
}

// AddRating adds i to the "rating" field.
func (uruo *UserReviewUpdateOne) AddRating(i int) *UserReviewUpdateOne {
	uruo.mutation.AddRating(i)
	return uruo
}

// SetReview sets the "review" field.
func (uruo *UserReviewUpdateOne) SetReview(s string) *UserReviewUpdateOne {
	uruo.mutation.SetReview(s)
	return uruo
}

// SetCreatedAt sets the "created_at" field.
func (uruo *UserReviewUpdateOne) SetCreatedAt(t time.Time) *UserReviewUpdateOne {
	uruo.mutation.SetCreatedAt(t)
	return uruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uruo *UserReviewUpdateOne) SetNillableCreatedAt(t *time.Time) *UserReviewUpdateOne {
	if t != nil {
		uruo.SetCreatedAt(*t)
	}
	return uruo
}

// SetUpdatedAt sets the "updated_at" field.
func (uruo *UserReviewUpdateOne) SetUpdatedAt(t time.Time) *UserReviewUpdateOne {
	uruo.mutation.SetUpdatedAt(t)
	return uruo
}

// SetUser sets the "user" edge to the User entity.
func (uruo *UserReviewUpdateOne) SetUser(u *User) *UserReviewUpdateOne {
	return uruo.SetUserID(u.ID)
}

// Mutation returns the UserReviewMutation object of the builder.
func (uruo *UserReviewUpdateOne) Mutation() *UserReviewMutation {
	return uruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (uruo *UserReviewUpdateOne) ClearUser() *UserReviewUpdateOne {
	uruo.mutation.ClearUser()
	return uruo
}

// Where appends a list predicates to the UserReviewUpdate builder.
func (uruo *UserReviewUpdateOne) Where(ps ...predicate.UserReview) *UserReviewUpdateOne {
	uruo.mutation.Where(ps...)
	return uruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uruo *UserReviewUpdateOne) Select(field string, fields ...string) *UserReviewUpdateOne {
	uruo.fields = append([]string{field}, fields...)
	return uruo
}

// Save executes the query and returns the updated UserReview entity.
func (uruo *UserReviewUpdateOne) Save(ctx context.Context) (*UserReview, error) {
	uruo.defaults()
	return withHooks(ctx, uruo.sqlSave, uruo.mutation, uruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserReviewUpdateOne) SaveX(ctx context.Context) *UserReview {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserReviewUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uruo *UserReviewUpdateOne) defaults() {
	if _, ok := uruo.mutation.UpdatedAt(); !ok {
		v := userreview.UpdateDefaultUpdatedAt()
		uruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uruo *UserReviewUpdateOne) check() error {
	if v, ok := uruo.mutation.Rating(); ok {
		if err := userreview.RatingValidator(v); err != nil {
			return &ValidationError{Name: "rating", err: fmt.Errorf(`ent: validator failed for field "UserReview.rating": %w`, err)}
		}
	}
	if v, ok := uruo.mutation.Review(); ok {
		if err := userreview.ReviewValidator(v); err != nil {
			return &ValidationError{Name: "review", err: fmt.Errorf(`ent: validator failed for field "UserReview.review": %w`, err)}
		}
	}
	if _, ok := uruo.mutation.UserID(); uruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserReview.user"`)
	}
	return nil
}

func (uruo *UserReviewUpdateOne) sqlSave(ctx context.Context) (_node *UserReview, err error) {
	if err := uruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(userreview.Table, userreview.Columns, sqlgraph.NewFieldSpec(userreview.FieldID, field.TypeInt))
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserReview.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userreview.FieldID)
		for _, f := range fields {
			if !userreview.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userreview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uruo.mutation.UserReviewID(); ok {
		_spec.SetField(userreview.FieldUserReviewID, field.TypeInt, value)
	}
	if value, ok := uruo.mutation.AddedUserReviewID(); ok {
		_spec.AddField(userreview.FieldUserReviewID, field.TypeInt, value)
	}
	if value, ok := uruo.mutation.OrderedProductID(); ok {
		_spec.SetField(userreview.FieldOrderedProductID, field.TypeInt, value)
	}
	if value, ok := uruo.mutation.AddedOrderedProductID(); ok {
		_spec.AddField(userreview.FieldOrderedProductID, field.TypeInt, value)
	}
	if value, ok := uruo.mutation.Rating(); ok {
		_spec.SetField(userreview.FieldRating, field.TypeInt, value)
	}
	if value, ok := uruo.mutation.AddedRating(); ok {
		_spec.AddField(userreview.FieldRating, field.TypeInt, value)
	}
	if value, ok := uruo.mutation.Review(); ok {
		_spec.SetField(userreview.FieldReview, field.TypeString, value)
	}
	if value, ok := uruo.mutation.CreatedAt(); ok {
		_spec.SetField(userreview.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := uruo.mutation.UpdatedAt(); ok {
		_spec.SetField(userreview.FieldUpdatedAt, field.TypeTime, value)
	}
	if uruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userreview.UserTable,
			Columns: []string{userreview.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userreview.UserTable,
			Columns: []string{userreview.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserReview{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userreview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uruo.mutation.done = true
	return _node, nil
}
