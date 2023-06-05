// Code generated by ent, DO NOT EDIT.

package nutritionalinformationtable

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the nutritionalinformationtable type in the database.
	Label = "nutritional_information_table"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParameter holds the string denoting the parameter field in the database.
	FieldParameter = "parameter"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldMeasurementUnit holds the string denoting the measurement_unit field in the database.
	FieldMeasurementUnit = "measurement_unit"
	// EdgeNutritionalInformation holds the string denoting the nutritional_information edge name in mutations.
	EdgeNutritionalInformation = "nutritional_information"
	// Table holds the table name of the nutritionalinformationtable in the database.
	Table = "nutritional_information_tables"
	// NutritionalInformationTable is the table that holds the nutritional_information relation/edge.
	NutritionalInformationTable = "nutritional_informations"
	// NutritionalInformationInverseTable is the table name for the NutritionalInformation entity.
	// It exists in this package in order to avoid circular dependency with the "nutritionalinformation" package.
	NutritionalInformationInverseTable = "nutritional_informations"
	// NutritionalInformationColumn is the table column denoting the nutritional_information relation/edge.
	NutritionalInformationColumn = "nutritional_information_table_id"
)

// Columns holds all SQL columns for nutritionalinformationtable fields.
var Columns = []string{
	FieldID,
	FieldParameter,
	FieldValue,
	FieldMeasurementUnit,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// ParameterValidator is a validator for the "parameter" field. It is called by the builders before save.
	ParameterValidator func(string) error
	// ValueValidator is a validator for the "value" field. It is called by the builders before save.
	ValueValidator func(float64) error
	// MeasurementUnitValidator is a validator for the "measurement_unit" field. It is called by the builders before save.
	MeasurementUnitValidator func(string) error
)

// OrderOption defines the ordering options for the NutritionalInformationTable queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByParameter orders the results by the parameter field.
func ByParameter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParameter, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByMeasurementUnit orders the results by the measurement_unit field.
func ByMeasurementUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMeasurementUnit, opts...).ToFunc()
}

// ByNutritionalInformationCount orders the results by nutritional_information count.
func ByNutritionalInformationCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newNutritionalInformationStep(), opts...)
	}
}

// ByNutritionalInformation orders the results by nutritional_information terms.
func ByNutritionalInformation(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newNutritionalInformationStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newNutritionalInformationStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(NutritionalInformationInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, NutritionalInformationTable, NutritionalInformationColumn),
	)
}