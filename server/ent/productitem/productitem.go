// Code generated by ent, DO NOT EDIT.

package productitem

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the productitem type in the database.
	Label = "product_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldStockKeepingUnit holds the string denoting the stock_keeping_unit field in the database.
	FieldStockKeepingUnit = "stock_keeping_unit"
	// FieldQuantityInStock holds the string denoting the quantity_in_stock field in the database.
	FieldQuantityInStock = "quantity_in_stock"
	// FieldProductImage holds the string denoting the product_image field in the database.
	FieldProductImage = "product_image"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeProduct holds the string denoting the product edge name in mutations.
	EdgeProduct = "product"
	// Table holds the table name of the productitem in the database.
	Table = "product_items"
	// ProductTable is the table that holds the product relation/edge.
	ProductTable = "product_items"
	// ProductInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductInverseTable = "products"
	// ProductColumn is the table column denoting the product relation/edge.
	ProductColumn = "product_id"
)

// Columns holds all SQL columns for productitem fields.
var Columns = []string{
	FieldID,
	FieldProductID,
	FieldStockKeepingUnit,
	FieldQuantityInStock,
	FieldProductImage,
	FieldPrice,
	FieldCreatedAt,
	FieldUpdatedAt,
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
	// StockKeepingUnitValidator is a validator for the "stock_keeping_unit" field. It is called by the builders before save.
	StockKeepingUnitValidator func(string) error
	// QuantityInStockValidator is a validator for the "quantity_in_stock" field. It is called by the builders before save.
	QuantityInStockValidator func(int) error
	// ProductImageValidator is a validator for the "product_image" field. It is called by the builders before save.
	ProductImageValidator func(string) error
	// PriceValidator is a validator for the "price" field. It is called by the builders before save.
	PriceValidator func(float32) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)

// OrderOption defines the ordering options for the ProductItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByStockKeepingUnit orders the results by the stock_keeping_unit field.
func ByStockKeepingUnit(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStockKeepingUnit, opts...).ToFunc()
}

// ByQuantityInStock orders the results by the quantity_in_stock field.
func ByQuantityInStock(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantityInStock, opts...).ToFunc()
}

// ByProductImage orders the results by the product_image field.
func ByProductImage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductImage, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByProductField orders the results by product field.
func ByProductField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductStep(), sql.OrderByField(field, opts...))
	}
}
func newProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ProductTable, ProductColumn),
	)
}