// Code generated by ent, DO NOT EDIT.

package shippingaddress

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the shippingaddress type in the database.
	Label = "shipping_address"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldAddressLine1 holds the string denoting the address_line1 field in the database.
	FieldAddressLine1 = "address_line1"
	// FieldAddressLine2 holds the string denoting the address_line2 field in the database.
	FieldAddressLine2 = "address_line2"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldCountry holds the string denoting the country field in the database.
	FieldCountry = "country"
	// FieldPostalCode holds the string denoting the postal_code field in the database.
	FieldPostalCode = "postal_code"
	// EdgeShopOrder holds the string denoting the shop_order edge name in mutations.
	EdgeShopOrder = "shop_order"
	// Table holds the table name of the shippingaddress in the database.
	Table = "shipping_addresses"
	// ShopOrderTable is the table that holds the shop_order relation/edge.
	ShopOrderTable = "shop_orders"
	// ShopOrderInverseTable is the table name for the ShopOrder entity.
	// It exists in this package in order to avoid circular dependency with the "shoporder" package.
	ShopOrderInverseTable = "shop_orders"
	// ShopOrderColumn is the table column denoting the shop_order relation/edge.
	ShopOrderColumn = "shipping_address_id"
)

// Columns holds all SQL columns for shippingaddress fields.
var Columns = []string{
	FieldID,
	FieldPhoneNumber,
	FieldAddressLine1,
	FieldAddressLine2,
	FieldCity,
	FieldState,
	FieldCountry,
	FieldPostalCode,
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
	// PhoneNumberValidator is a validator for the "phone_number" field. It is called by the builders before save.
	PhoneNumberValidator func(string) error
	// AddressLine1Validator is a validator for the "address_line1" field. It is called by the builders before save.
	AddressLine1Validator func(string) error
	// AddressLine2Validator is a validator for the "address_line2" field. It is called by the builders before save.
	AddressLine2Validator func(string) error
	// CityValidator is a validator for the "city" field. It is called by the builders before save.
	CityValidator func(string) error
	// StateValidator is a validator for the "state" field. It is called by the builders before save.
	StateValidator func(string) error
	// CountryValidator is a validator for the "country" field. It is called by the builders before save.
	CountryValidator func(string) error
	// PostalCodeValidator is a validator for the "postal_code" field. It is called by the builders before save.
	PostalCodeValidator func(string) error
)

// OrderOption defines the ordering options for the ShippingAddress queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPhoneNumber orders the results by the phone_number field.
func ByPhoneNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPhoneNumber, opts...).ToFunc()
}

// ByAddressLine1 orders the results by the address_line1 field.
func ByAddressLine1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressLine1, opts...).ToFunc()
}

// ByAddressLine2 orders the results by the address_line2 field.
func ByAddressLine2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressLine2, opts...).ToFunc()
}

// ByCity orders the results by the city field.
func ByCity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCity, opts...).ToFunc()
}

// ByState orders the results by the state field.
func ByState(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldState, opts...).ToFunc()
}

// ByCountry orders the results by the country field.
func ByCountry(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountry, opts...).ToFunc()
}

// ByPostalCode orders the results by the postal_code field.
func ByPostalCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostalCode, opts...).ToFunc()
}

// ByShopOrderCount orders the results by shop_order count.
func ByShopOrderCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newShopOrderStep(), opts...)
	}
}

// ByShopOrder orders the results by shop_order terms.
func ByShopOrder(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShopOrderStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newShopOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShopOrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ShopOrderTable, ShopOrderColumn),
	)
}
