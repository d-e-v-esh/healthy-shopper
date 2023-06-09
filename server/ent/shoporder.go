// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/orderstatus"
	"healthyshopper/ent/shippingaddress"
	"healthyshopper/ent/shippingmethod"
	"healthyshopper/ent/shoporder"
	"healthyshopper/ent/user"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ShopOrder is the model entity for the ShopOrder schema.
type ShopOrder struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OrderDateAndTime holds the value of the "order_date_and_time" field.
	OrderDateAndTime time.Time `json:"order_date_and_time,omitempty"`
	// PaymentMethod holds the value of the "payment_method" field.
	PaymentMethod string `json:"payment_method,omitempty"`
	// TotalPrice holds the value of the "total_price" field.
	TotalPrice float64 `json:"total_price,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// ShippingAddressID holds the value of the "shipping_address_id" field.
	ShippingAddressID int `json:"shipping_address_id,omitempty"`
	// ShippingMethodID holds the value of the "shipping_method_id" field.
	ShippingMethodID int `json:"shipping_method_id,omitempty"`
	// OrderStatusID holds the value of the "order_status_id" field.
	OrderStatusID int `json:"order_status_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShopOrderQuery when eager-loading is set.
	Edges        ShopOrderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ShopOrderEdges holds the relations/edges for other nodes in the graph.
type ShopOrderEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// ShippingMethod holds the value of the shipping_method edge.
	ShippingMethod *ShippingMethod `json:"shipping_method,omitempty"`
	// OrderStatus holds the value of the order_status edge.
	OrderStatus *OrderStatus `json:"order_status,omitempty"`
	// ShippingAddress holds the value of the shipping_address edge.
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShopOrderEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ShippingMethodOrErr returns the ShippingMethod value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShopOrderEdges) ShippingMethodOrErr() (*ShippingMethod, error) {
	if e.loadedTypes[1] {
		if e.ShippingMethod == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: shippingmethod.Label}
		}
		return e.ShippingMethod, nil
	}
	return nil, &NotLoadedError{edge: "shipping_method"}
}

// OrderStatusOrErr returns the OrderStatus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShopOrderEdges) OrderStatusOrErr() (*OrderStatus, error) {
	if e.loadedTypes[2] {
		if e.OrderStatus == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: orderstatus.Label}
		}
		return e.OrderStatus, nil
	}
	return nil, &NotLoadedError{edge: "order_status"}
}

// ShippingAddressOrErr returns the ShippingAddress value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShopOrderEdges) ShippingAddressOrErr() (*ShippingAddress, error) {
	if e.loadedTypes[3] {
		if e.ShippingAddress == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: shippingaddress.Label}
		}
		return e.ShippingAddress, nil
	}
	return nil, &NotLoadedError{edge: "shipping_address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShopOrder) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shoporder.FieldTotalPrice:
			values[i] = new(sql.NullFloat64)
		case shoporder.FieldID, shoporder.FieldUserID, shoporder.FieldShippingAddressID, shoporder.FieldShippingMethodID, shoporder.FieldOrderStatusID:
			values[i] = new(sql.NullInt64)
		case shoporder.FieldPaymentMethod:
			values[i] = new(sql.NullString)
		case shoporder.FieldOrderDateAndTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShopOrder fields.
func (so *ShopOrder) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shoporder.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			so.ID = int(value.Int64)
		case shoporder.FieldOrderDateAndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field order_date_and_time", values[i])
			} else if value.Valid {
				so.OrderDateAndTime = value.Time
			}
		case shoporder.FieldPaymentMethod:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_method", values[i])
			} else if value.Valid {
				so.PaymentMethod = value.String
			}
		case shoporder.FieldTotalPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total_price", values[i])
			} else if value.Valid {
				so.TotalPrice = value.Float64
			}
		case shoporder.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				so.UserID = int(value.Int64)
			}
		case shoporder.FieldShippingAddressID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shipping_address_id", values[i])
			} else if value.Valid {
				so.ShippingAddressID = int(value.Int64)
			}
		case shoporder.FieldShippingMethodID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shipping_method_id", values[i])
			} else if value.Valid {
				so.ShippingMethodID = int(value.Int64)
			}
		case shoporder.FieldOrderStatusID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_status_id", values[i])
			} else if value.Valid {
				so.OrderStatusID = int(value.Int64)
			}
		default:
			so.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ShopOrder.
// This includes values selected through modifiers, order, etc.
func (so *ShopOrder) Value(name string) (ent.Value, error) {
	return so.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the ShopOrder entity.
func (so *ShopOrder) QueryUser() *UserQuery {
	return NewShopOrderClient(so.config).QueryUser(so)
}

// QueryShippingMethod queries the "shipping_method" edge of the ShopOrder entity.
func (so *ShopOrder) QueryShippingMethod() *ShippingMethodQuery {
	return NewShopOrderClient(so.config).QueryShippingMethod(so)
}

// QueryOrderStatus queries the "order_status" edge of the ShopOrder entity.
func (so *ShopOrder) QueryOrderStatus() *OrderStatusQuery {
	return NewShopOrderClient(so.config).QueryOrderStatus(so)
}

// QueryShippingAddress queries the "shipping_address" edge of the ShopOrder entity.
func (so *ShopOrder) QueryShippingAddress() *ShippingAddressQuery {
	return NewShopOrderClient(so.config).QueryShippingAddress(so)
}

// Update returns a builder for updating this ShopOrder.
// Note that you need to call ShopOrder.Unwrap() before calling this method if this ShopOrder
// was returned from a transaction, and the transaction was committed or rolled back.
func (so *ShopOrder) Update() *ShopOrderUpdateOne {
	return NewShopOrderClient(so.config).UpdateOne(so)
}

// Unwrap unwraps the ShopOrder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (so *ShopOrder) Unwrap() *ShopOrder {
	_tx, ok := so.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShopOrder is not a transactional entity")
	}
	so.config.driver = _tx.drv
	return so
}

// String implements the fmt.Stringer.
func (so *ShopOrder) String() string {
	var builder strings.Builder
	builder.WriteString("ShopOrder(")
	builder.WriteString(fmt.Sprintf("id=%v, ", so.ID))
	builder.WriteString("order_date_and_time=")
	builder.WriteString(so.OrderDateAndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("payment_method=")
	builder.WriteString(so.PaymentMethod)
	builder.WriteString(", ")
	builder.WriteString("total_price=")
	builder.WriteString(fmt.Sprintf("%v", so.TotalPrice))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", so.UserID))
	builder.WriteString(", ")
	builder.WriteString("shipping_address_id=")
	builder.WriteString(fmt.Sprintf("%v", so.ShippingAddressID))
	builder.WriteString(", ")
	builder.WriteString("shipping_method_id=")
	builder.WriteString(fmt.Sprintf("%v", so.ShippingMethodID))
	builder.WriteString(", ")
	builder.WriteString("order_status_id=")
	builder.WriteString(fmt.Sprintf("%v", so.OrderStatusID))
	builder.WriteByte(')')
	return builder.String()
}

// ShopOrders is a parsable slice of ShopOrder.
type ShopOrders []*ShopOrder
