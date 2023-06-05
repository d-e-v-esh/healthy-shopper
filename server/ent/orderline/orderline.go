// Code generated by ent, DO NOT EDIT.

package orderline

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the orderline type in the database.
	Label = "order_line"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProductItemID holds the string denoting the product_item_id field in the database.
	FieldProductItemID = "product_item_id"
	// FieldShopOrderID holds the string denoting the shop_order_id field in the database.
	FieldShopOrderID = "shop_order_id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// Table holds the table name of the orderline in the database.
	Table = "order_lines"
)

// Columns holds all SQL columns for orderline fields.
var Columns = []string{
	FieldID,
	FieldProductItemID,
	FieldShopOrderID,
	FieldQuantity,
	FieldPrice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_lines"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_item_order_line",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	QuantityValidator func(int) error
	// PriceValidator is a validator for the "price" field. It is called by the builders before save.
	PriceValidator func(float64) error
)

// OrderOption defines the ordering options for the OrderLine queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProductItemID orders the results by the product_item_id field.
func ByProductItemID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductItemID, opts...).ToFunc()
}

// ByShopOrderID orders the results by the shop_order_id field.
func ByShopOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShopOrderID, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}
