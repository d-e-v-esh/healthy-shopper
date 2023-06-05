// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/nutritionalinformation"
	"healthyshopper/ent/nutritionalinformationtable"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// NutritionalInformation is the model entity for the NutritionalInformation schema.
type NutritionalInformation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// NutritionalInformationTableID holds the value of the "nutritional_information_table_id" field.
	NutritionalInformationTableID int `json:"nutritional_information_table_id,omitempty"`
	// NValue holds the value of the "n_value" field.
	NValue float64 `json:"n_value,omitempty"`
	// NMeasurementUnit holds the value of the "n_measurement_unit" field.
	NMeasurementUnit string `json:"n_measurement_unit,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NutritionalInformationQuery when eager-loading is set.
	Edges        NutritionalInformationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// NutritionalInformationEdges holds the relations/edges for other nodes in the graph.
type NutritionalInformationEdges struct {
	// Product holds the value of the product edge.
	Product []*Product `json:"product,omitempty"`
	// NutritionalInformationTable holds the value of the nutritional_information_table edge.
	NutritionalInformationTable *NutritionalInformationTable `json:"nutritional_information_table,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedProduct map[string][]*Product
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading.
func (e NutritionalInformationEdges) ProductOrErr() ([]*Product, error) {
	if e.loadedTypes[0] {
		return e.Product, nil
	}
	return nil, &NotLoadedError{edge: "product"}
}

// NutritionalInformationTableOrErr returns the NutritionalInformationTable value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NutritionalInformationEdges) NutritionalInformationTableOrErr() (*NutritionalInformationTable, error) {
	if e.loadedTypes[1] {
		if e.NutritionalInformationTable == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: nutritionalinformationtable.Label}
		}
		return e.NutritionalInformationTable, nil
	}
	return nil, &NotLoadedError{edge: "nutritional_information_table"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NutritionalInformation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case nutritionalinformation.FieldNValue:
			values[i] = new(sql.NullFloat64)
		case nutritionalinformation.FieldID, nutritionalinformation.FieldNutritionalInformationTableID:
			values[i] = new(sql.NullInt64)
		case nutritionalinformation.FieldNMeasurementUnit:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NutritionalInformation fields.
func (ni *NutritionalInformation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case nutritionalinformation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ni.ID = int(value.Int64)
		case nutritionalinformation.FieldNutritionalInformationTableID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nutritional_information_table_id", values[i])
			} else if value.Valid {
				ni.NutritionalInformationTableID = int(value.Int64)
			}
		case nutritionalinformation.FieldNValue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field n_value", values[i])
			} else if value.Valid {
				ni.NValue = value.Float64
			}
		case nutritionalinformation.FieldNMeasurementUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field n_measurement_unit", values[i])
			} else if value.Valid {
				ni.NMeasurementUnit = value.String
			}
		default:
			ni.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NutritionalInformation.
// This includes values selected through modifiers, order, etc.
func (ni *NutritionalInformation) Value(name string) (ent.Value, error) {
	return ni.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the NutritionalInformation entity.
func (ni *NutritionalInformation) QueryProduct() *ProductQuery {
	return NewNutritionalInformationClient(ni.config).QueryProduct(ni)
}

// QueryNutritionalInformationTable queries the "nutritional_information_table" edge of the NutritionalInformation entity.
func (ni *NutritionalInformation) QueryNutritionalInformationTable() *NutritionalInformationTableQuery {
	return NewNutritionalInformationClient(ni.config).QueryNutritionalInformationTable(ni)
}

// Update returns a builder for updating this NutritionalInformation.
// Note that you need to call NutritionalInformation.Unwrap() before calling this method if this NutritionalInformation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ni *NutritionalInformation) Update() *NutritionalInformationUpdateOne {
	return NewNutritionalInformationClient(ni.config).UpdateOne(ni)
}

// Unwrap unwraps the NutritionalInformation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ni *NutritionalInformation) Unwrap() *NutritionalInformation {
	_tx, ok := ni.config.driver.(*txDriver)
	if !ok {
		panic("ent: NutritionalInformation is not a transactional entity")
	}
	ni.config.driver = _tx.drv
	return ni
}

// String implements the fmt.Stringer.
func (ni *NutritionalInformation) String() string {
	var builder strings.Builder
	builder.WriteString("NutritionalInformation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ni.ID))
	builder.WriteString("nutritional_information_table_id=")
	builder.WriteString(fmt.Sprintf("%v", ni.NutritionalInformationTableID))
	builder.WriteString(", ")
	builder.WriteString("n_value=")
	builder.WriteString(fmt.Sprintf("%v", ni.NValue))
	builder.WriteString(", ")
	builder.WriteString("n_measurement_unit=")
	builder.WriteString(ni.NMeasurementUnit)
	builder.WriteByte(')')
	return builder.String()
}

// NamedProduct returns the Product named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ni *NutritionalInformation) NamedProduct(name string) ([]*Product, error) {
	if ni.Edges.namedProduct == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ni.Edges.namedProduct[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ni *NutritionalInformation) appendNamedProduct(name string, edges ...*Product) {
	if ni.Edges.namedProduct == nil {
		ni.Edges.namedProduct = make(map[string][]*Product)
	}
	if len(edges) == 0 {
		ni.Edges.namedProduct[name] = []*Product{}
	} else {
		ni.Edges.namedProduct[name] = append(ni.Edges.namedProduct[name], edges...)
	}
}

// NutritionalInformations is a parsable slice of NutritionalInformation.
type NutritionalInformations []*NutritionalInformation