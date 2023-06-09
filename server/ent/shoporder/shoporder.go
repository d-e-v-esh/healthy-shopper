// Code generated by ent, DO NOT EDIT.

package shoporder

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the shoporder type in the database.
	Label = "shop_order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOrderDateAndTime holds the string denoting the order_date_and_time field in the database.
	FieldOrderDateAndTime = "order_date_and_time"
	// FieldPaymentMethod holds the string denoting the payment_method field in the database.
	FieldPaymentMethod = "payment_method"
	// FieldTotalPrice holds the string denoting the total_price field in the database.
	FieldTotalPrice = "total_price"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldShippingAddressID holds the string denoting the shipping_address_id field in the database.
	FieldShippingAddressID = "shipping_address_id"
	// FieldShippingMethodID holds the string denoting the shipping_method_id field in the database.
	FieldShippingMethodID = "shipping_method_id"
	// FieldOrderStatusID holds the string denoting the order_status_id field in the database.
	FieldOrderStatusID = "order_status_id"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeShippingMethod holds the string denoting the shipping_method edge name in mutations.
	EdgeShippingMethod = "shipping_method"
	// EdgeOrderStatus holds the string denoting the order_status edge name in mutations.
	EdgeOrderStatus = "order_status"
	// EdgeShippingAddress holds the string denoting the shipping_address edge name in mutations.
	EdgeShippingAddress = "shipping_address"
	// Table holds the table name of the shoporder in the database.
	Table = "shop_orders"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "shop_orders"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// ShippingMethodTable is the table that holds the shipping_method relation/edge.
	ShippingMethodTable = "shop_orders"
	// ShippingMethodInverseTable is the table name for the ShippingMethod entity.
	// It exists in this package in order to avoid circular dependency with the "shippingmethod" package.
	ShippingMethodInverseTable = "shipping_methods"
	// ShippingMethodColumn is the table column denoting the shipping_method relation/edge.
	ShippingMethodColumn = "shipping_method_id"
	// OrderStatusTable is the table that holds the order_status relation/edge.
	OrderStatusTable = "shop_orders"
	// OrderStatusInverseTable is the table name for the OrderStatus entity.
	// It exists in this package in order to avoid circular dependency with the "orderstatus" package.
	OrderStatusInverseTable = "order_status"
	// OrderStatusColumn is the table column denoting the order_status relation/edge.
	OrderStatusColumn = "order_status_id"
	// ShippingAddressTable is the table that holds the shipping_address relation/edge.
	ShippingAddressTable = "shop_orders"
	// ShippingAddressInverseTable is the table name for the ShippingAddress entity.
	// It exists in this package in order to avoid circular dependency with the "shippingaddress" package.
	ShippingAddressInverseTable = "shipping_addresses"
	// ShippingAddressColumn is the table column denoting the shipping_address relation/edge.
	ShippingAddressColumn = "shipping_address_id"
)

// Columns holds all SQL columns for shoporder fields.
var Columns = []string{
	FieldID,
	FieldOrderDateAndTime,
	FieldPaymentMethod,
	FieldTotalPrice,
	FieldUserID,
	FieldShippingAddressID,
	FieldShippingMethodID,
	FieldOrderStatusID,
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
	// DefaultOrderDateAndTime holds the default value on creation for the "order_date_and_time" field.
	DefaultOrderDateAndTime func() time.Time
	// PaymentMethodValidator is a validator for the "payment_method" field. It is called by the builders before save.
	PaymentMethodValidator func(string) error
	// TotalPriceValidator is a validator for the "total_price" field. It is called by the builders before save.
	TotalPriceValidator func(float64) error
)

// OrderOption defines the ordering options for the ShopOrder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOrderDateAndTime orders the results by the order_date_and_time field.
func ByOrderDateAndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderDateAndTime, opts...).ToFunc()
}

// ByPaymentMethod orders the results by the payment_method field.
func ByPaymentMethod(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPaymentMethod, opts...).ToFunc()
}

// ByTotalPrice orders the results by the total_price field.
func ByTotalPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTotalPrice, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByShippingAddressID orders the results by the shipping_address_id field.
func ByShippingAddressID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippingAddressID, opts...).ToFunc()
}

// ByShippingMethodID orders the results by the shipping_method_id field.
func ByShippingMethodID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShippingMethodID, opts...).ToFunc()
}

// ByOrderStatusID orders the results by the order_status_id field.
func ByOrderStatusID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderStatusID, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByShippingMethodField orders the results by shipping_method field.
func ByShippingMethodField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShippingMethodStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrderStatusField orders the results by order_status field.
func ByOrderStatusField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStatusStep(), sql.OrderByField(field, opts...))
	}
}

// ByShippingAddressField orders the results by shipping_address field.
func ByShippingAddressField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newShippingAddressStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newShippingMethodStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShippingMethodInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ShippingMethodTable, ShippingMethodColumn),
	)
}
func newOrderStatusStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderStatusInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderStatusTable, OrderStatusColumn),
	)
}
func newShippingAddressStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ShippingAddressInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ShippingAddressTable, ShippingAddressColumn),
	)
}
