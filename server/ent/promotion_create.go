// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/product"
	"healthyshopper/ent/promotion"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PromotionCreate is the builder for creating a Promotion entity.
type PromotionCreate struct {
	config
	mutation *PromotionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PromotionCreate) SetName(s string) *PromotionCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PromotionCreate) SetDescription(s string) *PromotionCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetDiscountPercentage sets the "discount_percentage" field.
func (pc *PromotionCreate) SetDiscountPercentage(i int) *PromotionCreate {
	pc.mutation.SetDiscountPercentage(i)
	return pc
}

// SetStartDate sets the "start_date" field.
func (pc *PromotionCreate) SetStartDate(t time.Time) *PromotionCreate {
	pc.mutation.SetStartDate(t)
	return pc
}

// SetEndDate sets the "end_date" field.
func (pc *PromotionCreate) SetEndDate(t time.Time) *PromotionCreate {
	pc.mutation.SetEndDate(t)
	return pc
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (pc *PromotionCreate) AddProductIDs(ids ...int) *PromotionCreate {
	pc.mutation.AddProductIDs(ids...)
	return pc
}

// AddProduct adds the "product" edges to the Product entity.
func (pc *PromotionCreate) AddProduct(p ...*Product) *PromotionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddProductIDs(ids...)
}

// Mutation returns the PromotionMutation object of the builder.
func (pc *PromotionCreate) Mutation() *PromotionMutation {
	return pc.mutation
}

// Save creates the Promotion in the database.
func (pc *PromotionCreate) Save(ctx context.Context) (*Promotion, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PromotionCreate) SaveX(ctx context.Context) *Promotion {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PromotionCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PromotionCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PromotionCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Promotion.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := promotion.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Promotion.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Promotion.description"`)}
	}
	if v, ok := pc.mutation.Description(); ok {
		if err := promotion.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Promotion.description": %w`, err)}
		}
	}
	if _, ok := pc.mutation.DiscountPercentage(); !ok {
		return &ValidationError{Name: "discount_percentage", err: errors.New(`ent: missing required field "Promotion.discount_percentage"`)}
	}
	if v, ok := pc.mutation.DiscountPercentage(); ok {
		if err := promotion.DiscountPercentageValidator(v); err != nil {
			return &ValidationError{Name: "discount_percentage", err: fmt.Errorf(`ent: validator failed for field "Promotion.discount_percentage": %w`, err)}
		}
	}
	if _, ok := pc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "Promotion.start_date"`)}
	}
	if _, ok := pc.mutation.EndDate(); !ok {
		return &ValidationError{Name: "end_date", err: errors.New(`ent: missing required field "Promotion.end_date"`)}
	}
	if len(pc.mutation.ProductIDs()) == 0 {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "Promotion.product"`)}
	}
	return nil
}

func (pc *PromotionCreate) sqlSave(ctx context.Context) (*Promotion, error) {
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

func (pc *PromotionCreate) createSpec() (*Promotion, *sqlgraph.CreateSpec) {
	var (
		_node = &Promotion{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(promotion.Table, sqlgraph.NewFieldSpec(promotion.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(promotion.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(promotion.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.DiscountPercentage(); ok {
		_spec.SetField(promotion.FieldDiscountPercentage, field.TypeInt, value)
		_node.DiscountPercentage = value
	}
	if value, ok := pc.mutation.StartDate(); ok {
		_spec.SetField(promotion.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := pc.mutation.EndDate(); ok {
		_spec.SetField(promotion.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	if nodes := pc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   promotion.ProductTable,
			Columns: []string{promotion.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PromotionCreateBulk is the builder for creating many Promotion entities in bulk.
type PromotionCreateBulk struct {
	config
	builders []*PromotionCreate
}

// Save creates the Promotion entities in the database.
func (pcb *PromotionCreateBulk) Save(ctx context.Context) ([]*Promotion, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Promotion, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PromotionMutation)
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
func (pcb *PromotionCreateBulk) SaveX(ctx context.Context) []*Promotion {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PromotionCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PromotionCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}