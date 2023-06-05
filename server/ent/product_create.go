// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/product"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetProductImage sets the "product_image" field.
func (pc *ProductCreate) SetProductImage(s string) *ProductCreate {
	pc.mutation.SetProductImage(s)
	return pc
}

// SetProductCategoryID sets the "product_category_id" field.
func (pc *ProductCreate) SetProductCategoryID(i int) *ProductCreate {
	pc.mutation.SetProductCategoryID(i)
	return pc
}

// SetIngredientsListID sets the "ingredients_list_id" field.
func (pc *ProductCreate) SetIngredientsListID(i int) *ProductCreate {
	pc.mutation.SetIngredientsListID(i)
	return pc
}

// SetNutritionalInformationID sets the "nutritional_information_id" field.
func (pc *ProductCreate) SetNutritionalInformationID(i int) *ProductCreate {
	pc.mutation.SetNutritionalInformationID(i)
	return pc
}

// SetPromotionID sets the "promotion_id" field.
func (pc *ProductCreate) SetPromotionID(i int) *ProductCreate {
	pc.mutation.SetPromotionID(i)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProductCreate) SetCreatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProductCreate) SetUpdatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableUpdatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := product.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Product.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := product.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Product.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Product.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := product.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Product.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ProductImage(); !ok {
		return &ValidationError{Name: "product_image", err: errors.New(`ent: missing required field "Product.product_image"`)}
	}
	if v, ok := pc.mutation.ProductImage(); ok {
		if err := product.ProductImageValidator(v); err != nil {
			return &ValidationError{Name: "product_image", err: fmt.Errorf(`ent: validator failed for field "Product.product_image": %w`, err)}
		}
	}
	if _, ok := pc.mutation.ProductCategoryID(); !ok {
		return &ValidationError{Name: "product_category_id", err: errors.New(`ent: missing required field "Product.product_category_id"`)}
	}
	if v, ok := pc.mutation.ProductCategoryID(); ok {
		if err := product.ProductCategoryIDValidator(v); err != nil {
			return &ValidationError{Name: "product_category_id", err: fmt.Errorf(`ent: validator failed for field "Product.product_category_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.IngredientsListID(); !ok {
		return &ValidationError{Name: "ingredients_list_id", err: errors.New(`ent: missing required field "Product.ingredients_list_id"`)}
	}
	if v, ok := pc.mutation.IngredientsListID(); ok {
		if err := product.IngredientsListIDValidator(v); err != nil {
			return &ValidationError{Name: "ingredients_list_id", err: fmt.Errorf(`ent: validator failed for field "Product.ingredients_list_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.NutritionalInformationID(); !ok {
		return &ValidationError{Name: "nutritional_information_id", err: errors.New(`ent: missing required field "Product.nutritional_information_id"`)}
	}
	if v, ok := pc.mutation.NutritionalInformationID(); ok {
		if err := product.NutritionalInformationIDValidator(v); err != nil {
			return &ValidationError{Name: "nutritional_information_id", err: fmt.Errorf(`ent: validator failed for field "Product.nutritional_information_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.PromotionID(); !ok {
		return &ValidationError{Name: "promotion_id", err: errors.New(`ent: missing required field "Product.promotion_id"`)}
	}
	if v, ok := pc.mutation.PromotionID(); ok {
		if err := product.PromotionIDValidator(v); err != nil {
			return &ValidationError{Name: "promotion_id", err: fmt.Errorf(`ent: validator failed for field "Product.promotion_id": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Product.created_at"`)}
	}
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(product.Table, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.ProductImage(); ok {
		_spec.SetField(product.FieldProductImage, field.TypeString, value)
		_node.ProductImage = value
	}
	if value, ok := pc.mutation.ProductCategoryID(); ok {
		_spec.SetField(product.FieldProductCategoryID, field.TypeInt, value)
		_node.ProductCategoryID = value
	}
	if value, ok := pc.mutation.IngredientsListID(); ok {
		_spec.SetField(product.FieldIngredientsListID, field.TypeInt, value)
		_node.IngredientsListID = value
	}
	if value, ok := pc.mutation.NutritionalInformationID(); ok {
		_spec.SetField(product.FieldNutritionalInformationID, field.TypeInt, value)
		_node.NutritionalInformationID = value
	}
	if value, ok := pc.mutation.PromotionID(); ok {
		_spec.SetField(product.FieldPromotionID, field.TypeInt, value)
		_node.PromotionID = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
