// Code generated by ent, DO NOT EDIT.

package shippingmethod

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the shippingmethod type in the database.
	Label = "shipping_method"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldShippingMethod holds the string denoting the shipping_method field in the database.
	FieldShippingMethod = "shipping_method"
	// FieldShippingCost holds the string denoting the shipping_cost field in the database.
	FieldShippingCost = "shipping_cost"
	// Table holds the table name of the shippingmethod in the database.
	Table = "shipping_methods"
)

// Columns holds all SQL columns for shippingmethod fields.
var Columns = []string{
	FieldID,
	FieldShippingMethod,
	FieldShippingCost,
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
	// ShippingMethodValidator is a validator for the "shipping_method" field. It is called by the builders before save.
	ShippingMethodValidator func(string) error
	// ShippingCostValidator is a validator for the "shipping_cost" field. It is called by the builders before save.
	ShippingCostValidator func(float64) error
)

// OrderOption defines the ordering options for the ShippingMethod queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByShippingMethod orders the results by the shipping_method field.
func ByShippingMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippingMethod, opts...).ToFunc()
}

// ByShippingCost orders the results by the shipping_cost field.
func ByShippingCost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippingCost, opts...).ToFunc()
}
