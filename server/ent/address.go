// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/address"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Address is the model entity for the Address schema.
type Address struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AddressID holds the value of the "address_id" field.
	AddressID int `json:"address_id,omitempty"`
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
	// The values are being populated by the AddressQuery when eager-loading is set.
	Edges        AddressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AddressEdges holds the relations/edges for other nodes in the graph.
type AddressEdges struct {
	// UserAddress holds the value of the user_address edge.
	UserAddress []*UserAddress `json:"user_address,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedUserAddress map[string][]*UserAddress
}

// UserAddressOrErr returns the UserAddress value or an error if the edge
// was not loaded in eager-loading.
func (e AddressEdges) UserAddressOrErr() ([]*UserAddress, error) {
	if e.loadedTypes[0] {
		return e.UserAddress, nil
	}
	return nil, &NotLoadedError{edge: "user_address"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Address) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case address.FieldID, address.FieldAddressID:
			values[i] = new(sql.NullInt64)
		case address.FieldPhoneNumber, address.FieldAddressLine1, address.FieldAddressLine2, address.FieldCity, address.FieldState, address.FieldCountry, address.FieldPostalCode:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Address fields.
func (a *Address) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case address.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case address.FieldAddressID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field address_id", values[i])
			} else if value.Valid {
				a.AddressID = int(value.Int64)
			}
		case address.FieldPhoneNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone_number", values[i])
			} else if value.Valid {
				a.PhoneNumber = value.String
			}
		case address.FieldAddressLine1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line1", values[i])
			} else if value.Valid {
				a.AddressLine1 = value.String
			}
		case address.FieldAddressLine2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_line2", values[i])
			} else if value.Valid {
				a.AddressLine2 = value.String
			}
		case address.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				a.City = value.String
			}
		case address.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				a.State = value.String
			}
		case address.FieldCountry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field country", values[i])
			} else if value.Valid {
				a.Country = value.String
			}
		case address.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field postal_code", values[i])
			} else if value.Valid {
				a.PostalCode = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Address.
// This includes values selected through modifiers, order, etc.
func (a *Address) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryUserAddress queries the "user_address" edge of the Address entity.
func (a *Address) QueryUserAddress() *UserAddressQuery {
	return NewAddressClient(a.config).QueryUserAddress(a)
}

// Update returns a builder for updating this Address.
// Note that you need to call Address.Unwrap() before calling this method if this Address
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Address) Update() *AddressUpdateOne {
	return NewAddressClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Address entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Address) Unwrap() *Address {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Address is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Address) String() string {
	var builder strings.Builder
	builder.WriteString("Address(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("address_id=")
	builder.WriteString(fmt.Sprintf("%v", a.AddressID))
	builder.WriteString(", ")
	builder.WriteString("phone_number=")
	builder.WriteString(a.PhoneNumber)
	builder.WriteString(", ")
	builder.WriteString("address_line1=")
	builder.WriteString(a.AddressLine1)
	builder.WriteString(", ")
	builder.WriteString("address_line2=")
	builder.WriteString(a.AddressLine2)
	builder.WriteString(", ")
	builder.WriteString("city=")
	builder.WriteString(a.City)
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(a.State)
	builder.WriteString(", ")
	builder.WriteString("country=")
	builder.WriteString(a.Country)
	builder.WriteString(", ")
	builder.WriteString("postal_code=")
	builder.WriteString(a.PostalCode)
	builder.WriteByte(')')
	return builder.String()
}

// NamedUserAddress returns the UserAddress named value or an error if the edge was not
// loaded in eager-loading with this name.
func (a *Address) NamedUserAddress(name string) ([]*UserAddress, error) {
	if a.Edges.namedUserAddress == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := a.Edges.namedUserAddress[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (a *Address) appendNamedUserAddress(name string, edges ...*UserAddress) {
	if a.Edges.namedUserAddress == nil {
		a.Edges.namedUserAddress = make(map[string][]*UserAddress)
	}
	if len(edges) == 0 {
		a.Edges.namedUserAddress[name] = []*UserAddress{}
	} else {
		a.Edges.namedUserAddress[name] = append(a.Edges.namedUserAddress[name], edges...)
	}
}

// Addresses is a parsable slice of Address.
type Addresses []*Address
