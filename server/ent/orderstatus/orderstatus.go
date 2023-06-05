// Code generated by ent, DO NOT EDIT.

package orderstatus

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the orderstatus type in the database.
	Label = "order_status"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the orderstatus in the database.
	Table = "order_status"
)

// Columns holds all SQL columns for orderstatus fields.
var Columns = []string{
	FieldID,
	FieldStatus,
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
	// StatusValidator is a validator for the "status" field. It is called by the builders before save.
	StatusValidator func(string) error
)

// OrderOption defines the ordering options for the OrderStatus queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}
