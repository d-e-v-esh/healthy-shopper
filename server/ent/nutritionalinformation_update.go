// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"healthyshopper/ent/nutritionalinformation"
	"healthyshopper/ent/nutritionalinformationtable"
	"healthyshopper/ent/predicate"
	"healthyshopper/ent/product"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// NutritionalInformationUpdate is the builder for updating NutritionalInformation entities.
type NutritionalInformationUpdate struct {
	config
	hooks    []Hook
	mutation *NutritionalInformationMutation
}

// Where appends a list predicates to the NutritionalInformationUpdate builder.
func (niu *NutritionalInformationUpdate) Where(ps ...predicate.NutritionalInformation) *NutritionalInformationUpdate {
	niu.mutation.Where(ps...)
	return niu
}

// SetNutritionalInformationTableID sets the "nutritional_information_table_id" field.
func (niu *NutritionalInformationUpdate) SetNutritionalInformationTableID(i int) *NutritionalInformationUpdate {
	niu.mutation.SetNutritionalInformationTableID(i)
	return niu
}

// SetNValue sets the "n_value" field.
func (niu *NutritionalInformationUpdate) SetNValue(f float64) *NutritionalInformationUpdate {
	niu.mutation.ResetNValue()
	niu.mutation.SetNValue(f)
	return niu
}

// AddNValue adds f to the "n_value" field.
func (niu *NutritionalInformationUpdate) AddNValue(f float64) *NutritionalInformationUpdate {
	niu.mutation.AddNValue(f)
	return niu
}

// SetNMeasurementUnit sets the "n_measurement_unit" field.
func (niu *NutritionalInformationUpdate) SetNMeasurementUnit(s string) *NutritionalInformationUpdate {
	niu.mutation.SetNMeasurementUnit(s)
	return niu
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (niu *NutritionalInformationUpdate) AddProductIDs(ids ...int) *NutritionalInformationUpdate {
	niu.mutation.AddProductIDs(ids...)
	return niu
}

// AddProduct adds the "product" edges to the Product entity.
func (niu *NutritionalInformationUpdate) AddProduct(p ...*Product) *NutritionalInformationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return niu.AddProductIDs(ids...)
}

// SetNutritionalInformationTable sets the "nutritional_information_table" edge to the NutritionalInformationTable entity.
func (niu *NutritionalInformationUpdate) SetNutritionalInformationTable(n *NutritionalInformationTable) *NutritionalInformationUpdate {
	return niu.SetNutritionalInformationTableID(n.ID)
}

// Mutation returns the NutritionalInformationMutation object of the builder.
func (niu *NutritionalInformationUpdate) Mutation() *NutritionalInformationMutation {
	return niu.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (niu *NutritionalInformationUpdate) ClearProduct() *NutritionalInformationUpdate {
	niu.mutation.ClearProduct()
	return niu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (niu *NutritionalInformationUpdate) RemoveProductIDs(ids ...int) *NutritionalInformationUpdate {
	niu.mutation.RemoveProductIDs(ids...)
	return niu
}

// RemoveProduct removes "product" edges to Product entities.
func (niu *NutritionalInformationUpdate) RemoveProduct(p ...*Product) *NutritionalInformationUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return niu.RemoveProductIDs(ids...)
}

// ClearNutritionalInformationTable clears the "nutritional_information_table" edge to the NutritionalInformationTable entity.
func (niu *NutritionalInformationUpdate) ClearNutritionalInformationTable() *NutritionalInformationUpdate {
	niu.mutation.ClearNutritionalInformationTable()
	return niu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (niu *NutritionalInformationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, niu.sqlSave, niu.mutation, niu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (niu *NutritionalInformationUpdate) SaveX(ctx context.Context) int {
	affected, err := niu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (niu *NutritionalInformationUpdate) Exec(ctx context.Context) error {
	_, err := niu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (niu *NutritionalInformationUpdate) ExecX(ctx context.Context) {
	if err := niu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (niu *NutritionalInformationUpdate) check() error {
	if v, ok := niu.mutation.NValue(); ok {
		if err := nutritionalinformation.NValueValidator(v); err != nil {
			return &ValidationError{Name: "n_value", err: fmt.Errorf(`ent: validator failed for field "NutritionalInformation.n_value": %w`, err)}
		}
	}
	if v, ok := niu.mutation.NMeasurementUnit(); ok {
		if err := nutritionalinformation.NMeasurementUnitValidator(v); err != nil {
			return &ValidationError{Name: "n_measurement_unit", err: fmt.Errorf(`ent: validator failed for field "NutritionalInformation.n_measurement_unit": %w`, err)}
		}
	}
	if _, ok := niu.mutation.NutritionalInformationTableID(); niu.mutation.NutritionalInformationTableCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NutritionalInformation.nutritional_information_table"`)
	}
	return nil
}

func (niu *NutritionalInformationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := niu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(nutritionalinformation.Table, nutritionalinformation.Columns, sqlgraph.NewFieldSpec(nutritionalinformation.FieldID, field.TypeInt))
	if ps := niu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := niu.mutation.NValue(); ok {
		_spec.SetField(nutritionalinformation.FieldNValue, field.TypeFloat64, value)
	}
	if value, ok := niu.mutation.AddedNValue(); ok {
		_spec.AddField(nutritionalinformation.FieldNValue, field.TypeFloat64, value)
	}
	if value, ok := niu.mutation.NMeasurementUnit(); ok {
		_spec.SetField(nutritionalinformation.FieldNMeasurementUnit, field.TypeString, value)
	}
	if niu.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := niu.mutation.RemovedProductIDs(); len(nodes) > 0 && !niu.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := niu.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if niu.mutation.NutritionalInformationTableCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := niu.mutation.NutritionalInformationTableIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, niu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nutritionalinformation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	niu.mutation.done = true
	return n, nil
}

// NutritionalInformationUpdateOne is the builder for updating a single NutritionalInformation entity.
type NutritionalInformationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NutritionalInformationMutation
}

// SetNutritionalInformationTableID sets the "nutritional_information_table_id" field.
func (niuo *NutritionalInformationUpdateOne) SetNutritionalInformationTableID(i int) *NutritionalInformationUpdateOne {
	niuo.mutation.SetNutritionalInformationTableID(i)
	return niuo
}

// SetNValue sets the "n_value" field.
func (niuo *NutritionalInformationUpdateOne) SetNValue(f float64) *NutritionalInformationUpdateOne {
	niuo.mutation.ResetNValue()
	niuo.mutation.SetNValue(f)
	return niuo
}

// AddNValue adds f to the "n_value" field.
func (niuo *NutritionalInformationUpdateOne) AddNValue(f float64) *NutritionalInformationUpdateOne {
	niuo.mutation.AddNValue(f)
	return niuo
}

// SetNMeasurementUnit sets the "n_measurement_unit" field.
func (niuo *NutritionalInformationUpdateOne) SetNMeasurementUnit(s string) *NutritionalInformationUpdateOne {
	niuo.mutation.SetNMeasurementUnit(s)
	return niuo
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (niuo *NutritionalInformationUpdateOne) AddProductIDs(ids ...int) *NutritionalInformationUpdateOne {
	niuo.mutation.AddProductIDs(ids...)
	return niuo
}

// AddProduct adds the "product" edges to the Product entity.
func (niuo *NutritionalInformationUpdateOne) AddProduct(p ...*Product) *NutritionalInformationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return niuo.AddProductIDs(ids...)
}

// SetNutritionalInformationTable sets the "nutritional_information_table" edge to the NutritionalInformationTable entity.
func (niuo *NutritionalInformationUpdateOne) SetNutritionalInformationTable(n *NutritionalInformationTable) *NutritionalInformationUpdateOne {
	return niuo.SetNutritionalInformationTableID(n.ID)
}

// Mutation returns the NutritionalInformationMutation object of the builder.
func (niuo *NutritionalInformationUpdateOne) Mutation() *NutritionalInformationMutation {
	return niuo.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (niuo *NutritionalInformationUpdateOne) ClearProduct() *NutritionalInformationUpdateOne {
	niuo.mutation.ClearProduct()
	return niuo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (niuo *NutritionalInformationUpdateOne) RemoveProductIDs(ids ...int) *NutritionalInformationUpdateOne {
	niuo.mutation.RemoveProductIDs(ids...)
	return niuo
}

// RemoveProduct removes "product" edges to Product entities.
func (niuo *NutritionalInformationUpdateOne) RemoveProduct(p ...*Product) *NutritionalInformationUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return niuo.RemoveProductIDs(ids...)
}

// ClearNutritionalInformationTable clears the "nutritional_information_table" edge to the NutritionalInformationTable entity.
func (niuo *NutritionalInformationUpdateOne) ClearNutritionalInformationTable() *NutritionalInformationUpdateOne {
	niuo.mutation.ClearNutritionalInformationTable()
	return niuo
}

// Where appends a list predicates to the NutritionalInformationUpdate builder.
func (niuo *NutritionalInformationUpdateOne) Where(ps ...predicate.NutritionalInformation) *NutritionalInformationUpdateOne {
	niuo.mutation.Where(ps...)
	return niuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (niuo *NutritionalInformationUpdateOne) Select(field string, fields ...string) *NutritionalInformationUpdateOne {
	niuo.fields = append([]string{field}, fields...)
	return niuo
}

// Save executes the query and returns the updated NutritionalInformation entity.
func (niuo *NutritionalInformationUpdateOne) Save(ctx context.Context) (*NutritionalInformation, error) {
	return withHooks(ctx, niuo.sqlSave, niuo.mutation, niuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (niuo *NutritionalInformationUpdateOne) SaveX(ctx context.Context) *NutritionalInformation {
	node, err := niuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (niuo *NutritionalInformationUpdateOne) Exec(ctx context.Context) error {
	_, err := niuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (niuo *NutritionalInformationUpdateOne) ExecX(ctx context.Context) {
	if err := niuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (niuo *NutritionalInformationUpdateOne) check() error {
	if v, ok := niuo.mutation.NValue(); ok {
		if err := nutritionalinformation.NValueValidator(v); err != nil {
			return &ValidationError{Name: "n_value", err: fmt.Errorf(`ent: validator failed for field "NutritionalInformation.n_value": %w`, err)}
		}
	}
	if v, ok := niuo.mutation.NMeasurementUnit(); ok {
		if err := nutritionalinformation.NMeasurementUnitValidator(v); err != nil {
			return &ValidationError{Name: "n_measurement_unit", err: fmt.Errorf(`ent: validator failed for field "NutritionalInformation.n_measurement_unit": %w`, err)}
		}
	}
	if _, ok := niuo.mutation.NutritionalInformationTableID(); niuo.mutation.NutritionalInformationTableCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "NutritionalInformation.nutritional_information_table"`)
	}
	return nil
}

func (niuo *NutritionalInformationUpdateOne) sqlSave(ctx context.Context) (_node *NutritionalInformation, err error) {
	if err := niuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(nutritionalinformation.Table, nutritionalinformation.Columns, sqlgraph.NewFieldSpec(nutritionalinformation.FieldID, field.TypeInt))
	id, ok := niuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NutritionalInformation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := niuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nutritionalinformation.FieldID)
		for _, f := range fields {
			if !nutritionalinformation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nutritionalinformation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := niuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := niuo.mutation.NValue(); ok {
		_spec.SetField(nutritionalinformation.FieldNValue, field.TypeFloat64, value)
	}
	if value, ok := niuo.mutation.AddedNValue(); ok {
		_spec.AddField(nutritionalinformation.FieldNValue, field.TypeFloat64, value)
	}
	if value, ok := niuo.mutation.NMeasurementUnit(); ok {
		_spec.SetField(nutritionalinformation.FieldNMeasurementUnit, field.TypeString, value)
	}
	if niuo.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := niuo.mutation.RemovedProductIDs(); len(nodes) > 0 && !niuo.mutation.ProductCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := niuo.mutation.ProductIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if niuo.mutation.NutritionalInformationTableCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := niuo.mutation.NutritionalInformationTableIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &NutritionalInformation{config: niuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, niuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nutritionalinformation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	niuo.mutation.done = true
	return _node, nil
}