// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/shippingaddress"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ShippingAddress is the model entity for the ShippingAddress schema.
type ShippingAddress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PhoneNumber holds the value of the "phone_number" field.
	PhoneNumber string `json:"phone_number,omitempty"`
	// AddressLine1 holds the value of the "address_line1" field.
	AddressLine1 string `json:"address_line1,omitempty"`
	// AddressLine2 holds the value of the "address_line2" field.
	AddressLine2 string `json:"address_line2,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
	// Country holds the value of the "country" field.
	Country string `json:"country,omitempty"`
	// PostalCode holds the value of the "postal_code" field.
	PostalCode string `json:"postal_code,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShippingAddressQuery when eager-loading is set.
	Edges        ShippingAddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ShippingAddressEdges holds the relations/edges for other nodes in the graph.
type ShippingAddressEdges struct {
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
func (e ShippingAddressEdges) ShopOrderOrErr() ([]*ShopOrder, error) {
	if e.loadedTypes[0] {
		return e.ShopOrder, nil
	}
	return nil, &NotLoadedError{edge: "shop_order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShippingAddress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shippingaddress.FieldID:
			values[i] = new(sql.NullInt64)
		case shippingaddress.FieldPhoneNumber, shippingaddress.FieldAddressLine1, shippingaddress.FieldAddressLine2, shippingaddress.FieldCity, shippingaddress.FieldState, shippingaddress.FieldCountry, shippingaddress.FieldPostalCode:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShippingAddress fields.
func (sa *ShippingAddress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shippingaddress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sa.ID = int(value.Int64)
		case shippingaddress.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				sa.PhoneNumber = value.String
			}
		case shippingaddress.FieldAddressLine1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line1", values[i])
			} else if value.Valid {
				sa.AddressLine1 = value.String
			}
		case shippingaddress.FieldAddressLine2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line2", values[i])
			} else if value.Valid {
				sa.AddressLine2 = value.String
			}
		case shippingaddress.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				sa.City = value.String
			}
		case shippingaddress.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				sa.State = value.String
			}
		case shippingaddress.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				sa.Country = value.String
			}
		case shippingaddress.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				sa.PostalCode = value.String
			}
		default:
			sa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ShippingAddress.
// This includes values selected through modifiers, order, etc.
func (sa *ShippingAddress) Value(name string) (ent.Value, error) {
	return sa.selectValues.Get(name)
}

// QueryShopOrder queries the "shop_order" edge of the ShippingAddress entity.
func (sa *ShippingAddress) QueryShopOrder() *ShopOrderQuery {
	return NewShippingAddressClient(sa.config).QueryShopOrder(sa)
}

// Update returns a builder for updating this ShippingAddress.
// Note that you need to call ShippingAddress.Unwrap() before calling this method if this ShippingAddress
// was returned from a transaction, and the transaction was committed or rolled back.
func (sa *ShippingAddress) Update() *ShippingAddressUpdateOne {
	return NewShippingAddressClient(sa.config).UpdateOne(sa)
}

// Unwrap unwraps the ShippingAddress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sa *ShippingAddress) Unwrap() *ShippingAddress {
	_tx, ok := sa.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShippingAddress is not a transactional entity")
	}
	sa.config.driver = _tx.drv
	return sa
}

// String implements the fmt.Stringer.
func (sa *ShippingAddress) String() string {
	var builder strings.Builder
	builder.WriteString("ShippingAddress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sa.ID))
	builder.WriteString("phone_number=")
	builder.WriteString(sa.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("address_line1=")
	builder.WriteString(sa.AddressLine1)
	builder.WriteString(", ")
	builder.WriteString("address_line2=")
	builder.WriteString(sa.AddressLine2)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(sa.City)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(sa.State)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(sa.Country)
	builder.WriteString(", ")
	builder.WriteString("postal_code=")
	builder.WriteString(sa.PostalCode)
	builder.WriteByte(')')
	return builder.String()
}

// NamedShopOrder returns the ShopOrder named value or an error if the edge was not
// loaded in eager-loading with this name.
func (sa *ShippingAddress) NamedShopOrder(name string) ([]*ShopOrder, error) {
	if sa.Edges.namedShopOrder == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := sa.Edges.namedShopOrder[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (sa *ShippingAddress) appendNamedShopOrder(name string, edges ...*ShopOrder) {
	if sa.Edges.namedShopOrder == nil {
		sa.Edges.namedShopOrder = make(map[string][]*ShopOrder)
	}
	if len(edges) == 0 {
		sa.Edges.namedShopOrder[name] = []*ShopOrder{}
	} else {
		sa.Edges.namedShopOrder[name] = append(sa.Edges.namedShopOrder[name], edges...)
	}
}

// ShippingAddresses is a parsable slice of ShippingAddress.
type ShippingAddresses []*ShippingAddress
