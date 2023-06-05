// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/nutritionalinformation"
	"healthyshopper/ent/nutritionalinformationtable"
	"healthyshopper/ent/product"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NutritionalInformationCreate is the builder for creating a NutritionalInformation entity.
type NutritionalInformationCreate struct {
	config
	mutation *NutritionalInformationMutation
	hooks    []Hook
}

// SetNutritionalInformationTableID sets the "nutritional_information_table_id" field.
func (nic *NutritionalInformationCreate) SetNutritionalInformationTableID(i int) *NutritionalInformationCreate {
	nic.mutation.SetNutritionalInformationTableID(i)
	return nic
}

// SetNValue sets the "n_value" field.
func (nic *NutritionalInformationCreate) SetNValue(f float64) *NutritionalInformationCreate {
	nic.mutation.SetNValue(f)
	return nic
}

// SetNMeasurementUnit sets the "n_measurement_unit" field.
func (nic *NutritionalInformationCreate) SetNMeasurementUnit(s string) *NutritionalInformationCreate {
	nic.mutation.SetNMeasurementUnit(s)
	return nic
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (nic *NutritionalInformationCreate) AddProductIDs(ids ...int) *NutritionalInformationCreate {
	nic.mutation.AddProductIDs(ids...)
	return nic
}

// AddProduct adds the "product" edges to the Product entity.
func (nic *NutritionalInformationCreate) AddProduct(p ...*Product) *NutritionalInformationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return nic.AddProductIDs(ids...)
}

// SetNutritionalInformationTable sets the "nutritional_information_table" edge to the NutritionalInformationTable entity.
func (nic *NutritionalInformationCreate) SetNutritionalInformationTable(n *NutritionalInformationTable) *NutritionalInformationCreate {
	return nic.SetNutritionalInformationTableID(n.ID)
}

// Mutation returns the NutritionalInformationMutation object of the builder.
func (nic *NutritionalInformationCreate) Mutation() *NutritionalInformationMutation {
	return nic.mutation
}

// Save creates the NutritionalInformation in the database.
func (nic *NutritionalInformationCreate) Save(ctx context.Context) (*NutritionalInformation, error) {
	return withHooks(ctx, nic.sqlSave, nic.mutation, nic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (nic *NutritionalInformationCreate) SaveX(ctx context.Context) *NutritionalInformation {
	v, err := nic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nic *NutritionalInformationCreate) Exec(ctx context.Context) error {
	_, err := nic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nic *NutritionalInformationCreate) ExecX(ctx context.Context) {
	if err := nic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (nic *NutritionalInformationCreate) check() error {
	if _, ok := nic.mutation.NutritionalInformationTableID(); !ok {
		return &ValidationError{Name: "nutritional_information_table_id", err: errors.New(`ent: missing required field "NutritionalInformation.nutritional_information_table_id"`)}
	}
	if _, ok := nic.mutation.NValue(); !ok {
		return &ValidationError{Name: "n_value", err: errors.New(`ent: missing required field "NutritionalInformation.n_value"`)}
	}
	if v, ok := nic.mutation.NValue(); ok {
		if err := nutritionalinformation.NValueValidator(v); err != nil {
			return &ValidationError{Name: "n_value", err: fmt.Errorf(`ent: validator failed for field "NutritionalInformation.n_value": %w`, err)}
		}
	}
	if _, ok := nic.mutation.NMeasurementUnit(); !ok {
		return &ValidationError{Name: "n_measurement_unit", err: errors.New(`ent: missing required field "NutritionalInformation.n_measurement_unit"`)}
	}
	if v, ok := nic.mutation.NMeasurementUnit(); ok {
		if err := nutritionalinformation.NMeasurementUnitValidator(v); err != nil {
			return &ValidationError{Name: "n_measurement_unit", err: fmt.Errorf(`ent: validator failed for field "NutritionalInformation.n_measurement_unit": %w`, err)}
		}
	}
	if len(nic.mutation.ProductIDs()) == 0 {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "NutritionalInformation.product"`)}
	}
	if _, ok := nic.mutation.NutritionalInformationTableID(); !ok {
		return &ValidationError{Name: "nutritional_information_table", err: errors.New(`ent: missing required edge "NutritionalInformation.nutritional_information_table"`)}
	}
	return nil
}

func (nic *NutritionalInformationCreate) sqlSave(ctx context.Context) (*NutritionalInformation, error) {
	if err := nic.check(); err != nil {
		return nil, err
	}
	_node, _spec := nic.createSpec()
	if err := sqlgraph.CreateNode(ctx, nic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	nic.mutation.id = &_node.ID
	nic.mutation.done = true
	return _node, nil
}

func (nic *NutritionalInformationCreate) createSpec() (*NutritionalInformation, *sqlgraph.CreateSpec) {
	var (
		_node = &NutritionalInformation{config: nic.config}
		_spec = sqlgraph.NewCreateSpec(nutritionalinformation.Table, sqlgraph.NewFieldSpec(nutritionalinformation.FieldID, field.TypeInt))
	)
	if value, ok := nic.mutation.NValue(); ok {
		_spec.SetField(nutritionalinformation.FieldNValue, field.TypeFloat64, value)
		_node.NValue = value
	}
	if value, ok := nic.mutation.NMeasurementUnit(); ok {
		_spec.SetField(nutritionalinformation.FieldNMeasurementUnit, field.TypeString, value)
		_node.NMeasurementUnit = value
	}
	if nodes := nic.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   nutritionalinformation.ProductTable,
			Columns: []string{nutritionalinformation.ProductColumn},
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
	if nodes := nic.mutation.NutritionalInformationTableIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   nutritionalinformation.NutritionalInformationTableTable,
			Columns: []string{nutritionalinformation.NutritionalInformationTableColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(nutritionalinformationtable.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.NutritionalInformationTableID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NutritionalInformationCreateBulk is the builder for creating many NutritionalInformation entities in bulk.
type NutritionalInformationCreateBulk struct {
	config
	builders []*NutritionalInformationCreate
}

// Save creates the NutritionalInformation entities in the database.
func (nicb *NutritionalInformationCreateBulk) Save(ctx context.Context) ([]*NutritionalInformation, error) {
	specs := make([]*sqlgraph.CreateSpec, len(nicb.builders))
	nodes := make([]*NutritionalInformation, len(nicb.builders))
	mutators := make([]Mutator, len(nicb.builders))
	for i := range nicb.builders {
		func(i int, root context.Context) {
			builder := nicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NutritionalInformationMutation)
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
					_, err = mutators[i+1].Mutate(root, nicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, nicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, nicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (nicb *NutritionalInformationCreateBulk) SaveX(ctx context.Context) []*NutritionalInformation {
	v, err := nicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (nicb *NutritionalInformationCreateBulk) Exec(ctx context.Context) error {
	_, err := nicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nicb *NutritionalInformationCreateBulk) ExecX(ctx context.Context) {
	if err := nicb.Exec(ctx); err != nil {
		panic(err)
	}
}