// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/orderstatus"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderStatus is the model entity for the OrderStatus schema.
type OrderStatus struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderStatusQuery when eager-loading is set.
	Edges        OrderStatusEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderStatusEdges holds the relations/edges for other nodes in the graph.
type OrderStatusEdges struct {
	// ShopOrder holds the value of the shop_order edge.
	ShopOrder []*ShopOrder `json:"shop_order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedShopOrder map[string][]*ShopOrder
}

// ShopOrderOrErr returns the ShopOrder value or an error if the edge
// was not loaded in eager-loading.
func (e OrderStatusEdges) ShopOrderOrErr() ([]*ShopOrder, error) {
	if e.loadedTypes[0] {
		return e.ShopOrder, nil
	}
	return nil, &NotLoadedError{edge: "shop_order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderStatus) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderstatus.FieldID:
			values[i] = new(sql.NullInt64)
		case orderstatus.FieldStatus:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderStatus fields.
func (os *OrderStatus) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderstatus.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			os.ID = int(value.Int64)
		case orderstatus.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				os.Status = value.String
			}
		default:
			os.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderStatus.
// This includes values selected through modifiers, order, etc.
func (os *OrderStatus) Value(name string) (ent.Value, error) {
	return os.selectValues.Get(name)
}

// QueryShopOrder queries the "shop_order" edge of the OrderStatus entity.
func (os *OrderStatus) QueryShopOrder() *ShopOrderQuery {
	return NewOrderStatusClient(os.config).QueryShopOrder(os)
}

// Update returns a builder for updating this OrderStatus.
// Note that you need to call OrderStatus.Unwrap() before calling this method if this OrderStatus
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OrderStatus) Update() *OrderStatusUpdateOne {
	return NewOrderStatusClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the OrderStatus entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OrderStatus) Unwrap() *OrderStatus {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderStatus is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OrderStatus) String() string {
	var builder strings.Builder
	builder.WriteString("OrderStatus(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	builder.WriteString("status=")
	builder.WriteString(os.Status)
	builder.WriteByte(')')
	return builder.String()
}

// NamedShopOrder returns the ShopOrder named value or an error if the edge was not
// loaded in eager-loading with this name.
func (os *OrderStatus) NamedShopOrder(name string) ([]*ShopOrder, error) {
	if os.Edges.namedShopOrder == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := os.Edges.namedShopOrder[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (os *OrderStatus) appendNamedShopOrder(name string, edges ...*ShopOrder) {
	if os.Edges.namedShopOrder == nil {
		os.Edges.namedShopOrder = make(map[string][]*ShopOrder)
	}
	if len(edges) == 0 {
		os.Edges.namedShopOrder[name] = []*ShopOrder{}
	} else {
		os.Edges.namedShopOrder[name] = append(os.Edges.namedShopOrder[name], edges...)
	}
}

// OrderStatusSlice is a parsable slice of OrderStatus.
type OrderStatusSlice []*OrderStatus
