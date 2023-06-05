// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"healthyshopper/ent/shoppingcart"
	"healthyshopper/ent/shoppingcartitem"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ShoppingCartItem is the model entity for the ShoppingCartItem schema.
type ShoppingCartItem struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ShoppingCartID holds the value of the "shopping_cart_id" field.
	ShoppingCartID int `json:"shopping_cart_id,omitempty"`
	// ProductItemID holds the value of the "product_item_id" field.
	ProductItemID int `json:"product_item_id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShoppingCartItemQuery when eager-loading is set.
	Edges        ShoppingCartItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ShoppingCartItemEdges holds the relations/edges for other nodes in the graph.
type ShoppingCartItemEdges struct {
	// ShoppingCart holds the value of the shopping_cart edge.
	ShoppingCart *ShoppingCart `json:"shopping_cart,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// ShoppingCartOrErr returns the ShoppingCart value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ShoppingCartItemEdges) ShoppingCartOrErr() (*ShoppingCart, error) {
	if e.loadedTypes[0] {
		if e.ShoppingCart == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: shoppingcart.Label}
		}
		return e.ShoppingCart, nil
	}
	return nil, &NotLoadedError{edge: "shopping_cart"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ShoppingCartItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case shoppingcartitem.FieldID, shoppingcartitem.FieldShoppingCartID, shoppingcartitem.FieldProductItemID, shoppingcartitem.FieldQuantity:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ShoppingCartItem fields.
func (sci *ShoppingCartItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case shoppingcartitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sci.ID = int(value.Int64)
		case shoppingcartitem.FieldShoppingCartID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shopping_cart_id", values[i])
			} else if value.Valid {
				sci.ShoppingCartID = int(value.Int64)
			}
		case shoppingcartitem.FieldProductItemID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_item_id", values[i])
			} else if value.Valid {
				sci.ProductItemID = int(value.Int64)
			}
		case shoppingcartitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				sci.Quantity = int(value.Int64)
			}
		default:
			sci.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ShoppingCartItem.
// This includes values selected through modifiers, order, etc.
func (sci *ShoppingCartItem) Value(name string) (ent.Value, error) {
	return sci.selectValues.Get(name)
}

// QueryShoppingCart queries the "shopping_cart" edge of the ShoppingCartItem entity.
func (sci *ShoppingCartItem) QueryShoppingCart() *ShoppingCartQuery {
	return NewShoppingCartItemClient(sci.config).QueryShoppingCart(sci)
}

// Update returns a builder for updating this ShoppingCartItem.
// Note that you need to call ShoppingCartItem.Unwrap() before calling this method if this ShoppingCartItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (sci *ShoppingCartItem) Update() *ShoppingCartItemUpdateOne {
	return NewShoppingCartItemClient(sci.config).UpdateOne(sci)
}

// Unwrap unwraps the ShoppingCartItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sci *ShoppingCartItem) Unwrap() *ShoppingCartItem {
	_tx, ok := sci.config.driver.(*txDriver)
	if !ok {
		panic("ent: ShoppingCartItem is not a transactional entity")
	}
	sci.config.driver = _tx.drv
	return sci
}

// String implements the fmt.Stringer.
func (sci *ShoppingCartItem) String() string {
	var builder strings.Builder
	builder.WriteString("ShoppingCartItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sci.ID))
	builder.WriteString("shopping_cart_id=")
	builder.WriteString(fmt.Sprintf("%v", sci.ShoppingCartID))
	builder.WriteString(", ")
	builder.WriteString("product_item_id=")
	builder.WriteString(fmt.Sprintf("%v", sci.ProductItemID))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", sci.Quantity))
	builder.WriteByte(')')
	return builder.String()
}

// ShoppingCartItems is a parsable slice of ShoppingCartItem.
type ShoppingCartItems []*ShoppingCartItem